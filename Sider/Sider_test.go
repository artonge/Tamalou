package sider

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
)

var rawQuery = "Abdominal pain OR Gastrointestinal pain AND anorexia"

func TestQueryMeddraStr(t *testing.T) {

	results, err := QueryMeddraStr(rawQuery)
	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}

func TestQueryMeddraTree(t *testing.T) {

	tree := Queries.ParseQuery(rawQuery)
	clinicalSigns := Queries.GetClinicalSigns(rawQuery)
	results, err := QueryMeddraTree(tree)

	fmt.Println("Got ", len(results), " results")

	// Get frequency from all ids
	results, err = QueryMeddraFreq(results, clinicalSigns)

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}

	// for index, result := range results {
	// 	fmt.Println("Result: ", index, " =>> ", result.StitchCompoundId)
	// }
}

// func TestQueryMeddraAllIndications(t *testing.T) {
//
// 	results, err := QueryMeddraAllIndications(Queries.DBQuery{
// 		"1": 1,
// 	})
//
// 	if err != nil {
// 		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
// 		t.Fail()
// 	}
// }

// func TestQueryMeddraAllSe(t *testing.T) {
//
// 	results, err := QuerySideEffects(Queries.ParseQuery("Anaemia"))
//
// 	if err != nil {
// 		fmt.Println("Error in Sider Test:\n	==> ", err, "\n	==> ", results)
// 		t.Fail()
// 	}
//
// 	fmt.Println("Results: ", len(results))
// }

// func TestQueryMeddraFreq(t *testing.T) {
//
// 	results, err := QueryMeddraFreq(Queries.DBQuery{
// 		"1": 1,
// 	})
//
// 	if err != nil {
// 		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
// 		t.Fail()
// 	}
// }
