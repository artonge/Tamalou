package Omim

import (
	"fmt"
	"log"
	"testing"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
)

func TestOmimSearchQuery(t *testing.T) {
	results, err := QueryOmimIndex(Queries.ParseQuery("head"))
	if err != nil {
		fmt.Println(results)
		log.Fatal(err)
	}
}

func TestOmimSearchQueryAsync(t *testing.T) {
	diseasesChanel := make(chan []*Models.Disease)
	errorChanel := make(chan error)
	go QueryOmimIndexAsync(Queries.ParseQuery("head"), diseasesChanel, errorChanel)

	diseases, err := <-diseasesChanel, <-errorChanel

	if err != nil {
		fmt.Println(diseases, err)
		t.Fail()
	}
}
