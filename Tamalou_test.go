package main

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
)

func TestFetchDiseases(t *testing.T) {
	results, err := fetchDiseases(Queries.NewTamalouQuery("and", "", []Queries.ITamalouQuery{
		Queries.NewTamalouQuery("symptom", "Multicystic kidney dysplasia", nil),
	}))

	for _, r := range results {
		fmt.Println(r.Name)
	}

	if err != nil {
		fmt.Println("Unit test error: TestFetchDiseases:\n ==> ", err)
		t.Fail()
	}
}
