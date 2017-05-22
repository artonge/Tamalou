package sider

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// Import MySQL driver

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// var dbPath = "gmd-read:esial@tcp(neptune.telecomnancy.univ-lorraine.fr:3306)/gmd"

var dbPath = "root:root@tcp(localhost:3306)/sider"
var NbThread = 7
var clinicalSignsSql = ""

// init DB Connection
func init() {
	var err error
	db, err = sql.Open("mysql", dbPath)
	if err != nil {
		log.Fatal("Error init Sider MySQL connector init: ", err)
	}
}

// Take an array of meddra data and split it into 'lim' 2D array
func SplitQueries(buf []*Models.Drug, lim int) [][]*Models.Drug {
	var chunkSize = int(len(buf) / lim)
	chunks := make([][]*Models.Drug, 0, lim)
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
func JoinQueries(buckets [][]*Models.Drug) []*Models.Drug {
	var newDrugs = make([]*Models.Drug, 0, 100)
	for _, bucket := range buckets {
		for _, item := range bucket {
			newDrugs = append(newDrugs, item)
		}
	}
	return newDrugs
}

// Retrieve information on side effects of matched drug
func QueryMeddraFreq(drugs []*Models.Drug, clinicalSigns []string) ([]*Models.Drug, error) {

	// Build SQL string conditions based on clinical signs
	for _, value := range clinicalSigns {
		clinicalSignsSql += "WHEN meddra_freq.side_effect_name LIKE '" + value + "' THEN 1\n"
	}

	// Split queries into NbThread slices
	buckets := SplitQueries(drugs, NbThread)
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
func ClinicalSignsFromCompoundId(drugs []*Models.Drug, clinicalSignsSQL string, wg *sync.WaitGroup, treated *int) ([]*Models.Drug, error) {
	defer wg.Done()

	var currentQuery = ""
	for _, drug := range drugs {
		currentQuery = "SELECT DISTINCT(meddra_freq.side_effect_name), meddra_freq.placebo, meddra_freq.frequency_description, meddra_freq.freq_lower_bound, meddra_freq.freq_upper_bound, (CASE " + clinicalSignsSQL
		currentQuery += "ELSE 0 END) as matched FROM meddra_freq WHERE meddra_freq.stitch_compound_id1 = '" + drug.STITCH_ID_SIDER + "'"

		// Start query
		rows, err := db.Query(currentQuery)
		if err != nil {
			return nil, fmt.Errorf("Error while querying sider (meddra): %v", err)
		}
		defer rows.Close()

		// Retrieve results
		for rows.Next() {
			tmpMeddra := new(Models.SideEffect)
			err := rows.Scan(&tmpMeddra.SideEffectName, &tmpMeddra.Placebo, &tmpMeddra.Frequency, &tmpMeddra.FrequencyLowerBound, &tmpMeddra.FrequencyUpperBound, &tmpMeddra.Matched)
			if err != nil {
				return drugs, err
			}
			drug.SideEffects = append(drug.SideEffects, tmpMeddra)
			*treated += 1
		}
	}

	return drugs, nil
}

// QueryMeddra with parsed input string : a binary tree structure
func QueryMeddraTree(query Queries.ITamalouQuery) ([]*Models.Drug, error) {
	fmt.Println("Querying meddra")
	fullQuery := Queries.BuildSiderQuery(" stitch_compound_id1 IN (SELECT stitch_compound_id1 FROM meddra_all_se WHERE side_effect_name LIKE", query)
	return Getdrugs(fullQuery)
}

// QueryMeddra with input string ("ClinicalSign OR symptome")
func QueryMeddraStr(inputStr string) ([]*Models.Drug, error) {
	fullQuery := Queries.BuildSiderQuery(" stitch_compound_id1 IN (SELECT stitch_compound_id1 FROM meddra_all_se WHERE side_effect_name LIKE ", Queries.ParseQuery(inputStr))
	return Getdrugs(fullQuery)
}

// Retrieve all compoundId's drugs returned by the query
func Getdrugs(query string) ([]*Models.Drug, error) {
	fmt.Println("Getting drugs")
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra): %v", err)
	}
	defer rows.Close()
	var results = make([]*Models.Drug, 0, 100)

	for rows.Next() {
		tmpMeddra := new(Models.Drug)
		err := rows.Scan(&tmpMeddra.STITCH_ID_SIDER)
		if err != nil {
			return results, err
		}
		results = append(results, tmpMeddra)
	}
	return results, nil
}

func GetSideEffects(drugId string) ([]*Models.SideEffect, error) {
	currentQuery := "SELECT DISTINCT(meddra_freq.side_effect_name), meddra_freq.placebo, meddra_freq.frequency_description, meddra_freq.freq_lower_bound, meddra_freq.freq_upper_bound, (CASE " + clinicalSignsSql
	currentQuery += "ELSE 0 END) as matched FROM meddra_freq WHERE meddra_freq.stitch_compound_id1 = '" + drugId + "'"

	// Start query
	rows, err := db.Query(currentQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra): %v", err)
	}
	defer rows.Close()

	var sideEffects []*Models.SideEffect
	// Retrieve results
	for rows.Next() {
		tmpMeddra := new(Models.SideEffect)
		err := rows.Scan(&tmpMeddra.SideEffectName, &tmpMeddra.Placebo, &tmpMeddra.Frequency, &tmpMeddra.FrequencyLowerBound, &tmpMeddra.FrequencyUpperBound, &tmpMeddra.Matched)
		if err != nil {
			return sideEffects, err
		}
		sideEffects = append(sideEffects, tmpMeddra)
	}
	return sideEffects, nil
}
