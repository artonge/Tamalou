package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type HPOStruct struct {
	DiseaseDB    string
	DiseaseID    string
	DiseaseLabel string
	SignID       string
}

var db *sql.DB

// Init DB connection to HPO SQLite DB
func init() {
	var err error
	db, err = sql.Open("sqlite3", "hpo_annotations.sqlite")
	if err != nil {
		fmt.Println("Error in HPO SQLite init: ", err)
	}
}

// HPOQuery - Make a request on HPO SQLite database
// @param query : The query to make
//                Note that the query will be prepend with "SELECT ... FROM ... WHERE "
//                You can use the following fields names :
//                          'disease_db', 'disease_id', 'disease_label', 'sign_id'
// @return results : Array fill with HPOStruct
// @return error : Contains an error if one occurs
func HPOQuery(query string) ([]*HPOStruct, error) {
	// Make the query
	rows, err := db.Query("SELECT disease_db, disease_id, disease_label, sign_id FROM phenotype_annotation WHERE " + query)

	if err != nil {
		return nil, fmt.Errorf("Error in HPOQuery when querying: ", err)
	}

	// Prepare the resuts array
	results := make([]*HPOStruct, 0, 100)

	// Parse rows
	for rows.Next() {
		tmpHPOStruct := new(HPOStruct)
		err = rows.Scan(
			&tmpHPOStruct.DiseaseDB,
			&tmpHPOStruct.DiseaseID,
			&tmpHPOStruct.DiseaseLabel,
			&tmpHPOStruct.SignID,
		)

		if err != nil {
			return results, fmt.Errorf("Error in HPOQuery while scanning a row \n	==> %v\n	==> Error : %v", tmpHPOStruct, err)
		}

		// Append
		results = append(results, tmpHPOStruct)
	}

	return results, nil
}
