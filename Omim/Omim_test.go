package Omim

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/artonge/Tamalou/Queries"
	"github.com/artonge/Tamalou/indexing"
)

func TestIndexOmim(t *testing.T) {
	fmt.Println("Indexing omim file...")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory:\n Error ==> ", err, pwd)
	}
	err = os.RemoveAll(pwd + "/omim-search.bleve")
	if err != nil {
		fmt.Println("Error while removing old obo index:\n Error ==> ", err)
	}
	index, err = indexOmim()
	if err != nil {
		fmt.Println("Error while indexing Omim file:\n Error ==> ", err)
	}
	fmt.Println("Omim file indexed.")
}

func TestOmimSearchQuery(t *testing.T) {
	tquery := Queries.ParseQuery("head")
	results, err := indexing.SearchQuery(index, tquery, BuildOmimStructFromDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results[0])
}
