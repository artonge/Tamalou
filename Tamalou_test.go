package main

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
)

func TestFetchDiseases(t *testing.T) {
	diseasesChanel := make(chan []*Models.Disease)
	errorChanel := make(chan error)

	query := Queries.NewTamalouQuery("and", "", []Queries.ITamalouQuery{
		Queries.NewTamalouQuery("symptom", "Multicystic kidney dysplasia", nil),
	})
	go fetchDiseases(query, diseasesChanel, errorChanel)

	diseases, err := diseasesChanel, errorChanel

	if err != nil {
		fmt.Println(diseases)
		fmt.Println("Unit test error: TestFetchDiseases:\n ==> ", err)
		t.Fail()
	}
}
