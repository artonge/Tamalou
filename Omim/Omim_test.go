package Omim

import (
	"fmt"
	"log"
	"testing"

	"github.com/artonge/Tamalou/Queries"
	"github.com/artonge/Tamalou/indexing"
)

func TestOmimSearchQuery(t *testing.T) {
	tquery := Queries.ParseQuery("head")
	results, err := indexing.QueryIndex(index, tquery, buildOmimStructFromDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results[0].(omimStruct))
}
