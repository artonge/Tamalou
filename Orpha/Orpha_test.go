package orpha

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
)

func TestQuery(t *testing.T) {

	rawQuery := "Nausea/vomiting/regurgitation/merycism/hyperemesis OR Splenomegaly"
	// rawQuery := "Weight loss/loss of appetite/break in weight curve/general health alteration AND Splenomegaly"

	results, err := Query(Queries.ParseQuery(rawQuery))

	// fmt.Println(len(results), "results")
	// if len(results) > 0 {
	// 	fmt.Println("Diseases for ", rawQuery, ":")
	// 	for _, disease := range results {
	// 		fmt.Println("	- ", disease.Value.Disease.Name.Text, "(", disease.Value.Data.SignFreq.Name.Text, ")")
	// 	}
	// }
	if err != nil {
		fmt.Println("Error in Orpha TestQuery:\n	==> ", err, "\n	==> ", results)
		t.Fail()
	}
}

func TestDiseasesFromIDs(t *testing.T) {

	diseaseArray, err := GetDiseasesFromIDs([]float64{5, 46})

	fmt.Println(len(diseaseArray), "results:")
	for _, disease := range diseaseArray {
		fmt.Println("	- ", disease.Name)
	}
	if err != nil {
		fmt.Println(diseaseArray, err)
		t.Fail()
	}
}
