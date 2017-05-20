package sider

import (
	"database/sql"
	"fmt"
	"log"

	// Import MySQL river

	"github.com/artonge/Tamalou/Queries"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// init DB Connection
func init() {
	var err error
	//db, err = sql.Open("mysql", "gmd-read:esial@tcp(neptune.telecomnancy.univ-lorraine.fr:3306)/gmd")
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/sider")
	if err != nil {
		log.Fatal("Error init Sider MySQL connector init: ", err)
	}
}

func QueryMeddraTree(query Queries.ITamalouQuery) ([]*Meddra, error) {
	// Build query from required CS

	fullQuery := Queries.BuildSiderQuery(" stitch_compound_id1 IN (SELECT stitch_compound_id1 FROM meddra_all_se WHERE side_effect_name =", query)
	fmt.Println(fullQuery)

	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra): %v", err)
	}
	defer rows.Close()
	fmt.Println(rows.Columns())
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

// QueryMeddra ...
func QueryMeddraStr(inputStr string) ([]*Meddra, error) {
	// Build query from required CS

	fullQuery := Queries.BuildSiderQuery(" stitch_compound_id1 IN (SELECT stitch_compound_id1 FROM meddra_all_se WHERE side_effect_name =", Queries.ParseQuery(inputStr))
	fmt.Println(fullQuery)

	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while querying sider (meddra): %v", err)
	}
	defer rows.Close()
	fmt.Println(rows.Columns())
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

//
// // QueryMeddraAllIndications...
// func QueryMeddraAllIndications(query Queries.ITamalouQuery) ([]*MeddraAllIndications, error) {
// 	fullQuery := "SELECT * FROM meddra_all_indications WHERE " + Queries.BuildSQLQuery(query)
//
// 	// Make the query
// 	rows, err := db.Query(fullQuery)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error while querying sider: %v", err)
// 	}
// 	defer rows.Close()
//
// 	var results = make([]*MeddraAllIndications, 0, 100)
//
// 	for rows.Next() {
// 		tmpMeddraAllIndications := new(MeddraAllIndications)
// 		err := rows.Scan(&tmpMeddraAllIndications.StitchCompoundID, &tmpMeddraAllIndications.CUI, &tmpMeddraAllIndications.MethodOfDetection, &tmpMeddraAllIndications.ConceptName, &tmpMeddraAllIndications.MeddraConceptType, &tmpMeddraAllIndications.CUIOfMeddraTerm, &tmpMeddraAllIndications.MeddraConceptName)
// 		if err != nil {
// 			return results, err
// 		}
// 		results = append(results, tmpMeddraAllIndications)
// 	}
// 	return results, nil
// }
//
// func Query(query Queries.ITamalouQuery) ([]*MeddraAllSe, error) {
//
// 	return QuerySideEffects(query)
// }
//
// func QuerySideEffects(query Queries.ITamalouQuery) ([]*MeddraAllSe, error) {
// 	// Build SQL query base on the ITamalouQuery
// 	fullQuery := "SELECT stitch_compound_id1, stitch_compound_id2, cui, side_effect_name FROM meddra_all_se WHERE " + Queries.BuildSQLQuery(query) + ";"
// 	// Replace 'symptom' with 'side_effect_name'
// 	fullQuery = strings.Replace(fullQuery, "symptom", "side_effect_name", -1)
//
// 	// Make the query
// 	rows, err := db.Query(fullQuery)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error while querying sider (meddra_all_se): %v", err)
// 	}
// 	defer rows.Close()
//
// 	var results = make([]*MeddraAllSe, 0, 100)
//
// 	// Parse the results
// 	for rows.Next() {
// 		tmpMeddraAllSe := new(MeddraAllSe)
// 		err := rows.Scan(&tmpMeddraAllSe.StitchCompoundID1, &tmpMeddraAllSe.StitchCompoundID2, &tmpMeddraAllSe.CUI, &tmpMeddraAllSe.SideEffectName)
// 		if err != nil {
// 			return results, err
// 		}
// 		results = append(results, tmpMeddraAllSe)
// 	}
// 	return results, nil
// }
//
// func QueryMeddraFreq(query Queries.ITamalouQuery) ([]*MeddraFreq, error) {
// 	fullQuery := "SELECT * FROM meddra_freq WHERE " + Queries.BuildSQLQuery(query)
//
// 	// Make the query
// 	rows, err := db.Query(fullQuery)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error while querying sider: %v", err)
// 	}
// 	defer rows.Close()
//
// 	var results = make([]*MeddraFreq, 0, 100)
//
// 	for rows.Next() {
// 		tmpMeddraFreq := new(MeddraFreq)
// 		err := rows.Scan(&tmpMeddraFreq.StitchCompoundID1, &tmpMeddraFreq.StitchCompoundID2, &tmpMeddraFreq.CUI, &tmpMeddraFreq.Placebo, &tmpMeddraFreq.FrequencyDescription, &tmpMeddraFreq.FreqLowerBound, &tmpMeddraFreq.FreqUpperBound, &tmpMeddraFreq.MeddraConceptType, &tmpMeddraFreq.MeddraConceptID, &tmpMeddraFreq.SideEffectName)
// 		if err != nil {
// 			return results, err
// 		}
// 		results = append(results, tmpMeddraFreq)
// 	}
// 	return results, nil
// }
