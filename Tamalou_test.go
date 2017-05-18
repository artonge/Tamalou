package main

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
)

func TestFetchDiseases(t *testing.T) {
	_, err := fetchDiseases(Queries.NewTamalouQuery("and", "", []Queries.ITamalouQuery{
		Queries.NewTamalouQuery("symptom", "Multicystic kidney dysplasia", nil),
	}))

	if err != nil {
		fmt.Println("Unit test error: TestFetchDiseases:\n ==> ", err)
		t.Fail()
	}
}
