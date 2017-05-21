package sider

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// Import MySQL driver

	"github.com/artonge/Tamalou/Queries"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//var dbPath = "gmd-read:esial@tcp(neptune.telecomnancy.univ-lorraine.fr:3306)/gmd"
var dbPath = "root:root@tcp(localhost:3306)/sider"
var NbThread = 7

// init DB Connection
func init() {
	var err error
	db, err = sql.Open("mysql", dbPath)
	if err != nil {
		log.Fatal("Error init Sider MySQL connector init: ", err)
	}
}

// Take an array of meddra data and split it into 'lim' 2D array
func SplitQueries(buf []*Meddra, lim int) [][]*Meddra {
	var chunkSize = int(len(buf) / lim)
	chunks := make([][]*Meddra, 0, lim)
	for i := 0; i < lim; i++ {
		if i+1 == lim {
			chunks = append(chunks, buf[i*chunkSize:len(buf)])
		} else {
			chunks = append(chunks, buf[i*chunkSize:(i+1)*chunkSize])
		}
	}
	return chunks
}

// Take a 2D array of meddra data and return 1D array
func JoinQueries(buckets [][]*Meddra) []*Meddra {
	var newCompoundIds = make([]*Meddra, 0, 100)
	for _, bucket := range buckets {
		for _, item := range bucket {
			newCompoundIds = append(newCompoundIds, item)
		}
	}
	return newCompoundIds
}

// Retrieve information on side effects of matched drug
func QueryMeddraFreq(compoundIds []*Meddra, clinicalSigns []string) ([]*Meddra, error) {

	// Build SQL string conditions based on clinical signs
	var clinicalSignsSql = ""
	for _, value := range clinicalSigns {
		clinicalSignsSql += "WHEN meddra_freq.side_effect_name LIKE '" + value + "' THEN 1\n"
	}

	// Split queries into NbThread slices
	buckets := SplitQueries(compoundIds, NbThread)
	var wg sync.WaitGroup
	wg.Add(NbThread)

	// Start multi threading
	var treated = 0
	for i := 0; i < NbThread; i++ {
		go ClinicalSignsFromCompoundId(buckets[i], clinicalSignsSql, &wg, &treated)
	}
	wg.Wait()
	fmt.Println("Done")
	fmt.Println(treated, " rows treated")

	return JoinQueries(buckets), nil
}

// Job function to retrieve side effects information on matched compounds
func ClinicalSignsFromCompoundId(compoundIds []*Meddra, clinicalSignsSQL string, wg *sync.WaitGroup, treated *int) ([]*Meddra, error) {
	defer wg.Done()

	var currentQuery = ""
	for _, compoundId := range compoundIds {
		currentQuery = "SELECT DISTINCT(meddra_freq.side_effect_name), meddra_freq.placebo, meddra_freq.frequency_description, meddra_freq.freq_lower_bound, meddra_freq.freq_upper_bound, (CASE " + clinicalSignsSQL
		currentQuery += "ELSE 0 END) as matched FROM meddra_freq WHERE meddra_freq.stitch_compound_id1 = '" + compoundId.StitchCompoundId + "'"

		// Start query
		rows, err := db.Query(currentQuery)
		if err != nil {
			return nil, fmt.Errorf("Error while querying sider (meddra): %v", err)
		}
		defer rows.Close()

		// Retrieve results
		for rows.Next() {
			tmpMeddra := new(SideEffect)
			err := rows.Scan(&tmpMeddra.SideEffectName, &tmpMeddra.Placebo, &tmpMeddra.Frequency, &tmpMeddra.FrequencyLowerBound, &tmpMeddra.FrequencyUpperBound, &tmpMeddra.Matched)
			if err != nil {
				return compoundIds, err
			}
			compoundId.SideEffects = append(compoundId.SideEffects, tmpMeddra)
			*treated += 1
		}
	}

	return compoundIds, nil
}

// QueryMeddra with parsed input string : a binary tree structure
func QueryMeddraTree(query Queries.ITamalouQuery) ([]*Meddra, error) {
	fullQuery := Queries.BuildSiderQuery(" stitch_compound_id1 IN (SELECT stitch_compound_id1 FROM meddra_all_se WHERE side_effect_name LIKE", query)
	return GetCompoundIds(fullQuery)
}

// QueryMeddra with input string ("ClinicalSign OR symptome")
func QueryMeddraStr(inputStr string) ([]*Meddra, error) {
	fullQuery := Queries.BuildSiderQuery(" stitch_compound_id1 IN (SELECT stitch_compound_id1 FROM meddra_all_se WHERE side_effect_name LIKE ", Queries.ParseQuery(inputStr))
	return GetCompoundIds(fullQuery)
}

// Retrieve all compoundId's drugs returned by the query
func GetCompoundIds(query string) ([]*Meddra, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra): %v", err)
	}
	defer rows.Close()
	var results = make([]*Meddra, 0, 100)

	for rows.Next() {
		tmpMeddra := new(Meddra)
		err := rows.Scan(&tmpMeddra.StitchCompoundId)
		if err != nil {
			return results, err
		}
		results = append(results, tmpMeddra)
	}
	return results, nil
}
