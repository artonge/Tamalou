package HPO

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/artonge/Tamalou/Models"
	// Import SQLite drivers
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Init DB connection to HPO SQLite DB
func init() {
	var err error
	db, err = sql.Open("sqlite3", "datas/hpo/hpo_annotations.sqlite")
	if err != nil {
		fmt.Println("Error in HPO SQLite init: ", err)
	}
}

// SQLiteQuery - Make a request on HPO SQLite database
// @param query : The query to make
//                Note that the query will be prepend with "SELECT ... FROM ... WHERE "
//                You can use the following fields names :
//                          'disease_db', 'disease_id', 'disease_label', 'sign_id'
// @return results : Array fill with HPOSQLiteStruct
// @return error : Contains an error if one occurs
func SQLiteQuery(diseasesIDs []string) ([]*Models.Disease, error) {
	fullQuery := "SELECT disease_db, disease_id, disease_label, sign_id FROM phenotype_annotation WHERE "

	for _, id := range diseasesIDs {
		fullQuery += "sign_id='" + id + "' OR "
	}

	// Remove the last " OR "
	fullQuery = fullQuery[:len(fullQuery)-len(" OR ")]

	// Make the query
	rows, err := db.Query(fullQuery)
	if err != nil {
		return nil, fmt.Errorf("Error in SQLiteQuery when querying: %v", err)
	}

	// Prepare the resuts array
	diseasesArray := make([]*Models.Disease, 0, 100)

	// Parse rows
	for rows.Next() {
		tmpHPOSQLiteStruct := new(hpoSQLiteStruct)
		tmpDisease := new(Models.Disease)
		err = rows.Scan(
			&tmpHPOSQLiteStruct.DiseaseDB,
			&tmpHPOSQLiteStruct.DiseaseID,
			&tmpHPOSQLiteStruct.DiseaseLabel,
			&tmpHPOSQLiteStruct.SignID,
		)
		if err != nil {
			return nil, fmt.Errorf("Error in SQLiteQuery while scanning a row \n	==> %v\n	==> Error : %v", tmpHPOSQLiteStruct, err)
		}

		tmpDisease.Name = tmpHPOSQLiteStruct.DiseaseLabel
		tmpDisease.Sources = append(tmpDisease.Sources, "HPO")
		switch tmpHPOSQLiteStruct.DiseaseDB {
		case "ORPHA":
			value, err := strconv.ParseFloat(tmpHPOSQLiteStruct.DiseaseID, 64)
			if err != nil {
				return nil, fmt.Errorf("Error in SQLiteQuery while parsing float\n	==> %v\n	==> Error : %v", tmpHPOSQLiteStruct.DiseaseID, err)
			}
			tmpDisease.OrphaID = value
		case "OMIM":
			tmpDisease.OMIMID = tmpHPOSQLiteStruct.DiseaseID
		}
		// Append
		diseasesArray = append(diseasesArray, tmpDisease)
	}

	return diseasesArray, nil
}
