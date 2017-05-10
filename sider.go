package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnection() {

	db, err := sql.Open("mysql", "gmd-read:esial@tcp(neptune.telecomnancy.univ-lorraine.fr:3306)/gmd")

	if err != nil {
		log.Fatal("Error1", err)
	}

	tables, err := getTables(db)

	if err != nil {
		log.Fatal("Error2", err)
		return
	}

	for _, table := range tables {
		columns, err1 := getColumns(db, table)
		if err != nil {
			log.Fatal("Error3", err1)
			continue
		}

		fmt.Println(table, " ==> ", columns)
	}
}

func getColumns(db *sql.DB, table string) ([]string, error) {

	rows, err := db.Query("select * from " + table + ";")

	if err != nil {
		return nil, err
	}

	return rows.Columns()
}

func getTables(db *sql.DB) ([]string, error) {

	var tables = make([]string, 10)
	var i int

	rows, err := db.Query("show tables;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&tables[i])
		if err != nil {
			return nil, err
		}
		i++
	}

	return tables, nil
}
