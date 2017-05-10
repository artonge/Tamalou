package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type HPOStruct struct {
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "hpo_annotations.sqlite")
	if err != nil {
		fmt.Println("Error in HPO SQLite init: ", err)
	}
}

func SQLiteConnection() {

	rows, err := db.Query("SELECT * FROM phenotype_annotation")
	if err != nil {
		fmt.Println("Error 2: ", err)
	}

	types, err := rows.ColumnTypes()
	if err != nil {
		fmt.Println("Error 3: ", err)
	}

	for _, t := range types {
		fmt.Println(t.Name(), t.DatabaseTypeName())
	}

	// fmt.Println("Err:", rows)

	rows.Next()
	var (
		disease_db        string
		disease_id        string
		disease_label     string
		col_4             interface{}
		sign_id           string
		disease_db_and_id string
		col_7             interface{}
		col_8             interface{}
		col_9             interface{}
		col_10            interface{}
		col_11            interface{}
		col_12            interface{}
		col_13            interface{}
		col_14            interface{}
	)

	err = rows.Scan(
		&disease_db,
		&disease_id,
		&disease_label,
		&col_4,
		&sign_id,
		&disease_db_and_id,
		&col_7,
		&col_8,
		&col_9,
		&col_10,
		&col_11,
		&col_12,
		&col_13,
		&col_14)
	if err != nil {
		fmt.Println("Error 3: ", err)
	}

	fmt.Println("ok",
		"disease_db: ", disease_db, " - \n",
		"disease_id: ", disease_id, " - \n",
		"disease_label: ", disease_label, " - \n",
		"sign_id: ", sign_id, " - \n",
		"disease_db_and_id: ", disease_db_and_id, " - \n")
}

func HPOQuery(query string) ([]HPOStruct, error) {

	rows, err := db.Query("SELECT * FROM phenotype_annotation")
	if err != nil {
		return nil, fmt.Errorf("Error in HPOQuery when querying: ", err)
	}

	fmt.Println(rows)

	return nil, nil
}
