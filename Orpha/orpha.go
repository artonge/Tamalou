package orpha

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/artonge/Tamalou/Queries"
	couchdb "github.com/rhinoman/couchdb-go"
)

var DB *couchdb.Database

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

	DB = conn.SelectDB("orphadatabase", nil)
}

// Fetch all diceases for the given ITamalouQuery
func Query(query Queries.ITamalouQuery) ([]ViewResult, error) {
	switch query.Type() {
	case "or":
		// Make a Query for all children of the OR node
		// Append results together then remove duplicates
		var results []ViewResult
		for _, child := range query.Children() {
			subResults, err := Query(child)
			if err != nil {
				return results, err
			}
			// Merge if necessary
			if len(results) > 0 {
				results = merge(results, subResults, "or")
			} else {
				results = subResults
			}
		}
		return results, nil
	case "and":
		// Make a Query for all children of the AND node
		// Merge the results in order to only have the diseases shared by all clinicalSigns
		var results []ViewResult
		for _, child := range query.Children() {
			subResults, err := Query(child)
			if err != nil {
				return results, err
			}
			// Merge if necessary
			if len(results) > 0 {
				results = merge(results, subResults, "and")
			} else {
				results = subResults
			}
		}
		return results, nil
	case "symptom":
		// Make a request to couchDB when the query's type is a symptom
		return getDiseaseByClinicalSign(query.Value())
	default:
		return nil, fmt.Errorf("Error while querying Orpha:\n	==> Error in query format: %v", query)
	}
}

// Interface to the'getDiseaseByClinicalSign' view of the DB
// Supports wild cards (*)
func getDiseaseByClinicalSign(clinicalSign string) ([]ViewResult, error) {
	results := ViewResponse{}
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

	// Make the request
	err := DB.GetView("clinicalsigns", "GetDiseaseByClinicalSign", &results, &params)
	if err != nil {
		return results.Rows, fmt.Errorf("Error while Querying Orpha:\n	==>  %v", err)
	}

	return results.Rows, err
}

// Merge to arrays
// Support "or" and "and" logic operator
func merge(list1 []ViewResult, list2 []ViewResult, operator string) []ViewResult {
	var results []ViewResult

	switch operator {
	case "and":
		// Put all item of list1 contained in list2 in results
		for _, item1 := range list1 {
			for _, item2 := range list2 {
				if item1.Value.Disease.Name.Text == item2.Value.Disease.Name.Text {
					results = append(results, item1)
				}
			}
		}

		// Put all item of list2 contained in list1 in results
		// Check that the item is not allready in results
		for _, item2 := range list2 {
			for _, item1 := range list1 {
				if item1.Value.Disease.Name.Text == item2.Value.Disease.Name.Text {
					contains := false
					for _, itemR := range results {
						if itemR.Value.Disease.Name.Text == item2.Value.Disease.Name.Text {
							contains = true
							break
						}
					}
					if !contains {
						results = append(results, item2)
					}
				}
			}
		}
	case "or":
		// Put all item of list1 in results
		for _, item1 := range list1 {
			results = append(results, item1)
		}
		// Put all item if list2 in results
		// Check that the item is not allready in results
		for _, item2 := range list2 {
			contains := false
			for _, itemR := range results {
				if itemR.Value.Disease.Name.Text == item2.Value.Disease.Name.Text {
					contains = true
					break
				}
			}
			if !contains {
				results = append(results, item2)
			}
		}
	default:
		fmt.Println("Operator <" + operator + "> not suported for merging")
	}

	return results
}
