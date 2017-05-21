package sider

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
)

var rawQuery = "Abdominal pain OR Gastrointestinal pain AND anorexia" //  AND Alopecia OR Decreased appetite OR Anxiety

// func BenchmarkSplitJoinQueries(b *testing.B) {
// 	fmt.Println("Benchmarking sider split queries")
// 	results, err := QueryMeddraStr(rawQuery)
// 	fmt.Println("Got ", len(results), " results")
// 	if err != nil {
// 		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
// 		b.Fail()
// 	}
// 	JoinQueries(SplitQueries(results, 10))
// }

// func TestQueryMeddraStr(t *testing.T) {
//
// 	results, err := QueryMeddraStr(rawQuery)
// 	if err != nil {
// 		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
// 		t.Fail()
// 	}
// }

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

	// for _, value := range results {
	// 	fmt.Println("Stitch compound id =>>", value.StitchCompoundId)
	// for _, se := range value.SideEffects {
	// 	fmt.Println("\tSide Effect :")
	// 	fmt.Println("\t\tName =>>", se.SideEffectName)
	// 	fmt.Println("\t\tPlacebo =>>", se.Placebo)
	// 	fmt.Println("\t\tFrequency =>>", se.Frequency)
	// 	fmt.Println("\t\tFrequency Lower =>>", se.FrequencyLowerBound)
	// 	fmt.Println("\t\tFrequency Upper=>>", se.FrequencyUpperBound)
	// 	fmt.Println("\t\tMatched =>>", se.Matched)
	// }
	// }
}
