package sider

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/artonge/Tamalou/TamalouQuery"
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

func QueryMeddra(query map[string]interface{}) ([]*Meddra, error) {
	fullQuery := "SELECT * FROM meddra WHERE " + TamalouQuery.BuildSQLQuery(query, "")

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra): ", err)
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

func QueryMeddraAllIndications(query map[string]interface{}) ([]*MeddraAllIndications, error) {
	fullQuery := "SELECT * FROM meddra_all_indications WHERE " + TamalouQuery.BuildSQLQuery(query, "")

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider: ", err)
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

func QueryMeddraAllSe(query map[string]interface{}) ([]*MeddraAllSe, error) {
	fullQuery := "SELECT * FROM meddra_all_se WHERE " + TamalouQuery.BuildSQLQuery(query, "")

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra_all_se): ", err)
	}
	defer rows.Close()

	var results = make([]*MeddraAllSe, 0, 100)

	for rows.Next() {
		tmpMeddraAllSe := new(MeddraAllSe)
		err := rows.Scan(&tmpMeddraAllSe.StitchCompoundID1, &tmpMeddraAllSe.StitchCompoundID2, &tmpMeddraAllSe.CUI, &tmpMeddraAllSe.MeddraConceptType, &tmpMeddraAllSe.CUIOfMeddraTerm, &tmpMeddraAllSe.SideEffectName)
		if err != nil {
			return results, err
		}
		results = append(results, tmpMeddraAllSe)
	}
	return results, nil
}

func QueryMeddraFreq(query map[string]interface{}) ([]*MeddraFreq, error) {
	fullQuery := "SELECT * FROM meddra_freq WHERE " + TamalouQuery.BuildSQLQuery(query, "")

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider: ", err)
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
