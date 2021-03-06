package orpha

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
	couchdb "github.com/rhinoman/couchdb-go"
)

var db *couchdb.Database

// Init CouchDB connection
func init() {
	conn, err := couchdb.NewConnection(
		"couchdb.telecomnancy.univ-lorraine.fr",
		80,
		time.Duration(55000*time.Millisecond),
	)

	if err != nil {
		log.Fatal("Error in Orpha DB init:\n	==> ", err)
	}

	db = conn.SelectDB("orphadatabase", nil)
}

// Query - Fetch all diceases for the given ITamalouQuery
func Query(query Queries.ITamalouQuery) ([]*Models.Disease, error) {
	switch query.Type() {
	case "or":
	case "and":
		// Make a Query for all children of the AND node
		// Merge the results in order to only have the diseases shared by all clinicalSigns
		var results []*Models.Disease
		for _, child := range query.Children() {
			subResults, err := Query(child)
			if err != nil {
				return nil, err
			}
			// Merge if necessary
			if len(results) > 0 {
				results = Models.Merge(results, subResults, string(query.Type()))
			} else {
				results = subResults
			}
		}
		return results, nil
	case "symptom":
		// Make a request to couchDB when the query's type is a symptom
		return getDiseaseByClinicalSign(query.Value().(string))
	default:
		return nil, fmt.Errorf("Error while querying Orpha:\n	==> Error in query format: %v", query)
	}

	return nil, nil
}

// Interface to the'getDiseaseByClinicalSign' view of the DB
// Supports wild cards (*)
func getDiseaseByClinicalSign(clinicalSign string) ([]*Models.Disease, error) {
	queryResults := viewResponse{}
	var params url.Values
	// Add quotes around the sign for json format
	formatedClinicalSign := "\"" + clinicalSign + "\""

	// Allow wild ward in the request
	if strings.Contains(clinicalSign, "*") {
		// Replace the '*' char with 'a' and 'Z'
		// Then use the 'startkey' and 'endkey' params of couchDB
		// All string between 'startkey' and 'endkey' will be returned
		params = url.Values{
			"startkey": []string{strings.Replace(formatedClinicalSign, "*", "a", -1)},
			"endkey":   []string{strings.Replace(formatedClinicalSign, "*", "Z", -1)},
		}
	} else {
		// Simple matching request to couchDB
		params = url.Values{"key": []string{formatedClinicalSign}}
	}

	// Make the request to couchDB
	err := db.GetView("clinicalsigns", "GetDiseaseByClinicalSign", &queryResults, &params)
	if err != nil {
		return nil, fmt.Errorf("Error while Querying Orpha:\n	==>  %v", err)
	}

	// Put diseases from queryResults in diseasesArray
	var diseasesArray []*Models.Disease

	// Get all the diseases from queryResults, format them and put them in diseasesArray
	for _, row := range queryResults.Rows {
		tmpDisease := new(Models.Disease)
		tmpDisease.Name = row.Value["disease"].(map[string]interface{})["Name"].(map[string]interface{})["text"].(string)
		tmpDisease.OrphaID = row.Value["disease"].(map[string]interface{})["OrphaNumber"].(float64)
		tmpDisease.Sources = append(tmpDisease.Sources, "Orpha")
		diseasesArray = append(diseasesArray, tmpDisease)
	}

	return diseasesArray, err
}

// GetDiseasesFromIDs - Interface to the'getDisease' view of the DB
// Return the diseases informations from their IDs
func GetDiseasesFromIDs(diseasesIDs []float64) ([]*Models.Disease, error) {
	// Build ID json array
	IDList, err := json.Marshal(diseasesIDs)
	if err != nil {
		return nil, fmt.Errorf("Error while converting IDs array to json :\n	==>  %v", err)
	}
	params := url.Values{"keys": []string{strings.Replace(fmt.Sprint(IDList), " ", ", ", -1)}}

	// Make the request to couchDB
	queryResults := viewResponse{}
	err = db.GetView("diseases", "GetDiseases", &queryResults, &params)
	if err != nil {
		return nil, fmt.Errorf("Error while Querying Orpha:\n	==>  %v", err)
	}

	// Put diseases from queryResults in diseasesArray
	var diseasesArray []*Models.Disease

	// Get all the diseases from queryResults, format them and put them in diseasesArray
	for _, row := range queryResults.Rows {
		tmpDisease := new(Models.Disease)
		tmpDisease.Name = row.Value["Name"].(map[string]interface{})["text"].(string)
		tmpDisease.OrphaID = row.Value["OrphaNumber"].(float64)
		tmpDisease.Sources = append(tmpDisease.Sources, "Orpha")
		diseasesArray = append(diseasesArray, tmpDisease)
	}

	return diseasesArray, err
}
