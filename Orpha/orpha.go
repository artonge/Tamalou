package orpha

import (
	"fmt"
	"log"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
)

var DB *couchdb.Database

// Init CouchDB connection
func init() {
	conn, err := couchdb.NewConnection(
		"couchdb.telecomnancy.univ-lorraine.fr",
		80,
		time.Duration(35000*time.Millisecond),
	)

	if err != nil {
		log.Fatal("Error in Orpha DB init: ", err)
	}

	DB = conn.SelectDB("orphadatabase", nil)
}

func Query(query map[string]interface{}) ([]ViewResult, error) {

	results := ViewResponse{}

	err := DB.GetView("clinicalsigns", "GetDiseaseByClinicalSign", &results, nil)

	if err != nil {
		return nil, fmt.Errorf("Error while Querying Orpha: ", err)
	}

	return results.Rows, err
}
