package sider

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/artonge/Tamalou/Queries"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Init DB Connection
func init() {
	var err error
	db, err = sql.Open("mysql", "gmd-read:esial@tcp(neptune.telecomnancy.univ-lorraine.fr:3306)/gmd")
	if err != nil {
		log.Fatal("Error init Sider MySQL connector init: ", err)
	}
}

func QueryMeddra(query Queries.ITamalouQuery) ([]*Meddra, error) {
	fullQuery := "SELECT * FROM meddra WHERE " + Queries.BuildSQLQuery(query)

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra): %v", err)
	}
	defer rows.Close()

	var results = make([]*Meddra, 0, 100)

	for rows.Next() {
		tmpMeddra := new(Meddra)
		err := rows.Scan(&tmpMeddra.CUI, &tmpMeddra.ConceptType, &tmpMeddra.MeddraID, &tmpMeddra.Label)
		if err != nil {
			return results, err
		}
		results = append(results, tmpMeddra)
	}
	return results, nil
}

func QueryMeddraAllIndications(query Queries.ITamalouQuery) ([]*MeddraAllIndications, error) {
	fullQuery := "SELECT * FROM meddra_all_indications WHERE " + Queries.BuildSQLQuery(query)

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider: %v", err)
	}
	defer rows.Close()

	var results = make([]*MeddraAllIndications, 0, 100)

	for rows.Next() {
		tmpMeddraAllIndications := new(MeddraAllIndications)
		err := rows.Scan(&tmpMeddraAllIndications.StitchCompoundID, &tmpMeddraAllIndications.CUI, &tmpMeddraAllIndications.MethodOfDetection, &tmpMeddraAllIndications.ConceptName, &tmpMeddraAllIndications.MeddraConceptType, &tmpMeddraAllIndications.CUIOfMeddraTerm, &tmpMeddraAllIndications.MeddraConceptName)
		if err != nil {
			return results, err
		}
		results = append(results, tmpMeddraAllIndications)
	}
	return results, nil
}

func Query(query Queries.ITamalouQuery) ([]*MeddraAllSe, error) {

	return QuerySideEffects(query)
	return nil, nil
}

func QuerySideEffects(query Queries.ITamalouQuery) ([]*MeddraAllSe, error) {
	// Build SQL query base on the ITamalouQuery
	fullQuery := "SELECT stitch_compound_id1, stitch_compound_id2, cui, side_effect_name FROM meddra_all_se WHERE " + Queries.BuildSQLQuery(query) + ";"
	// Replace 'symptom' with 'side_effect_name'
	fullQuery = strings.Replace(fullQuery, "symptom", "side_effect_name", -1)

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra_all_se): %v", err)
	}
	defer rows.Close()

	var results = make([]*MeddraAllSe, 0, 100)

	// Parse the results
	for rows.Next() {
		tmpMeddraAllSe := new(MeddraAllSe)
		err := rows.Scan(&tmpMeddraAllSe.StitchCompoundID1, &tmpMeddraAllSe.StitchCompoundID2, &tmpMeddraAllSe.CUI, &tmpMeddraAllSe.SideEffectName)
		if err != nil {
			return results, err
		}
		results = append(results, tmpMeddraAllSe)
	}
	return results, nil
}

func QueryMeddraFreq(query Queries.ITamalouQuery) ([]*MeddraFreq, error) {
	fullQuery := "SELECT * FROM meddra_freq WHERE " + Queries.BuildSQLQuery(query)

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider: %v", err)
	}
	defer rows.Close()

	var results = make([]*MeddraFreq, 0, 100)

	for rows.Next() {
		tmpMeddraFreq := new(MeddraFreq)
		err := rows.Scan(&tmpMeddraFreq.StitchCompoundID1, &tmpMeddraFreq.StitchCompoundID2, &tmpMeddraFreq.CUI, &tmpMeddraFreq.Placebo, &tmpMeddraFreq.FrequencyDescription, &tmpMeddraFreq.FreqLowerBound, &tmpMeddraFreq.FreqUpperBound, &tmpMeddraFreq.MeddraConceptType, &tmpMeddraFreq.MeddraConceptID, &tmpMeddraFreq.SideEffectName)
		if err != nil {
			return results, err
		}
		results = append(results, tmpMeddraFreq)
	}
	return results, nil
}
