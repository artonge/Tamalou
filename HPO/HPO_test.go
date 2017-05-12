package HPO

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
	"github.com/blevesearch/bleve"
)

func TestSQLiteQuery(t *testing.T) {
	// Fetch all
	hpoArray, err := HPOQuery(Queries.DBQuery{
		"1": "1",
	})
	if err != nil {
		fmt.Println("Unit Test Error (1=1): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}

	// Fetch none - wrong query
	hpoArray, err = HPOQuery(Queries.DBQuery{
		"fail": "success",
	})
	if err == nil {
		fmt.Println("Unit Test Error (fail=success): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}

	// Fetch some
	hpoArray, err = HPOQuery(Queries.DBQuery{
		"and": Queries.DBQuery{
			"disease_id": "1",
			"1":          1,
		},
	})
	if err != nil || hpoArray[0].DiseaseID != "1" {
		fmt.Println("Unit Test Error (disease_id=1 AND 1 = 1): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}
}

func TestOBOQuery(t *testing.T) {
	// Find one
	results, err := HPOOBOQuery(Queries.DBQuery{
		"id": "HP:0000001",
	})

	if err != nil || len(results) != 1 {
		fmt.Println("Unit Test Error (id=HP:0000001): \n	==> ", err, "\n	==> ", results)
		fmt.Println(results, err)
		t.Fail()
	}

	// Find none
	results, err = HPOOBOQuery(Queries.DBQuery{
		"id": "none",
	})

	if err != nil || len(results) != 0 {
		fmt.Println("Unit Test Error (id=HP:0000001): \n	==> ", err, "\n	==> ", results)
		fmt.Println(results, err)
		t.Fail()
	}

	// Fail query
	results, err = HPOOBOQuery(Queries.DBQuery{
		"id": 1,
	})

	if err != nil || len(results) != 0 {
		fmt.Println("Unit Test Error (id=HP:0000001): \n	==> ", err, "\n	==> ", results)
		fmt.Println(results, err)
		t.Fail()
	}
}

func TestIndexOBO(t *testing.T) {

	index, err := IndexOBO()

	if err != nil {
		fmt.Println(err)
	}

	query := bleve.NewMatchQuery("Root of all terms in the Human Phenotype Ontology.")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)

	fmt.Println(searchResults)
}
