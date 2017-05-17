package HPO

import (
	"database/sql"
	"fmt"

	"github.com/artonge/Tamalou/Queries"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Init DB connection to HPO SQLite DB
func init() {
	var err error
	db, err = sql.Open("sqlite3", "../datas/hpo/hpo_annotations.sqlite")
	if err != nil {
		fmt.Println("Error in HPO SQLite init: ", err)
	}
}

// HPOQuery - Make a request on HPO SQLite database
// @param query : The query to make
//                Note that the query will be prepend with "SELECT ... FROM ... WHERE "
//                You can use the following fields names :
//                          'disease_db', 'disease_id', 'disease_label', 'sign_id'
// @return results : Array fill with HPOSQLiteStruct
// @return error : Contains an error if one occurs
func HPOQuery(query Queries.ITamalouQuery) ([]*HPOSQLiteStruct, error) {
	fullQuery := "SELECT disease_db, disease_id, disease_label, sign_id FROM phenotype_annotation WHERE " + Queries.BuildSQLQuery(query)

	// Make the query
	rows, err := db.Query(fullQuery)

	if err != nil {
		return nil, fmt.Errorf("Error in HPOQuery when querying: %v", err)
	}

	// Prepare the resuts array
	results := make([]*HPOSQLiteStruct, 0, 100)

	// Parse rows
	for rows.Next() {
		tmpHPOSQLiteStruct := new(HPOSQLiteStruct)
		err = rows.Scan(
			&tmpHPOSQLiteStruct.DiseaseDB,
			&tmpHPOSQLiteStruct.DiseaseID,
			&tmpHPOSQLiteStruct.DiseaseLabel,
			&tmpHPOSQLiteStruct.SignID,
		)

		if err != nil {
			return results, fmt.Errorf("Error in HPOQuery while scanning a row \n	==> %v\n	==> Error : %v", tmpHPOSQLiteStruct, err)
		}

		// Append
		results = append(results, tmpHPOSQLiteStruct)
	}

	return results, nil
}
