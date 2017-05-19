package sider

import (
	"fmt"
	"testing"
)

func TestQueryMeddra(t *testing.T) {

	results, err := QueryMeddra("Abdominal pain OR Gastrointestinal pain AND anorexia")

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}

	fmt.Println("Results: ", len(results))
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
