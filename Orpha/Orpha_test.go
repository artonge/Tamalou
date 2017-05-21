package orpha

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
)

func TestQuery(t *testing.T) {
	// rawQuery := "Nausea/vomiting/regurgitation/merycism/hyperemesis OR Splenomegaly"
	rawQuery := "Weight loss/loss of appetite/break in weight curve/general health alteration AND Splenomegaly"

	results, err := QueryOrpha(Queries.ParseQuery(rawQuery))

	if err != nil {
		fmt.Println(len(results), "results")
		if len(results) > 0 {
			fmt.Println("Diseases for ", rawQuery, ":")
			for _, disease := range results {
				fmt.Println("	- ", disease.Name)
			}
		}
		fmt.Println("Error in Orpha TestQuery:\n	==> ", err, "\n	==> ", results)
		t.Fail()
	}
}

func TestDiseasesFromIDs(t *testing.T) {
	diseaseArray, err := GetDiseasesFromIDs([]float64{5, 46})

	if err != nil {
		fmt.Println(len(diseaseArray), "results:")
		for _, disease := range diseaseArray {
			fmt.Println("	- ", disease.Name)
		}
		fmt.Println(diseaseArray, err)
		t.Fail()
	}
}

func TestQueryAsync(t *testing.T) {
	// rawQuery := "Weight loss/loss of appetite/break in weight curve/general health alteration AND Splenomegaly"
	rawQuery := "Nausea/vomiting/regurgitation/merycism/hyperemesis OR Splenomegaly"

	diseasesChanel := make(chan []*Models.Disease)
	errorChanel := make(chan error)

	go QueryAsync(Queries.ParseQuery(rawQuery), diseasesChanel, errorChanel)

	diseases, err := <-diseasesChanel, <-errorChanel

	if err != nil {
		fmt.Println(diseases, err)
		t.Fail()
	}
}

func TestDiseasesFromIDsAsync(t *testing.T) {
	diseasesChanel := make(chan []*Models.Disease)
	errorChanel := make(chan error)

	go GetDiseasesFromIDsAsync([]float64{5, 46}, diseasesChanel, errorChanel)

	diseases, err := <-diseasesChanel, <-errorChanel

	if err != nil {
		fmt.Println(len(diseases), "results:")
		for _, disease := range diseases {
			fmt.Println("	- ", disease.Name)
		}
		fmt.Println(diseases, err)
		t.Fail()
	}
}
