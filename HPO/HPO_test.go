package HPO

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
	"github.com/blevesearch/bleve"
)

func TestSQLiteQuery(t *testing.T) {
	// Fetch all
	hpoArray, err := HPOQuery(Queries.NewTamalouQuery("1", "1", nil))
	if err != nil {
		fmt.Println("Unit Test Error (1=1): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}

	// Fetch none - wrong query
	hpoArray, err = HPOQuery(Queries.NewTamalouQuery("fail", "success", nil))
	if err == nil {
		fmt.Println("Unit Test Error (fail=success): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}

	// Fetch some
	hpoArray, err = HPOQuery(Queries.NewTamalouQuery("and", "", []Queries.ITamalouQuery{
		Queries.NewTamalouQuery("disease_id", 1, nil),
		Queries.NewTamalouQuery("1", 1, nil),
	}))
	fmt.Println(hpoArray)
	if err != nil || hpoArray[0].DiseaseID != "1" {
		fmt.Println("Unit Test Error (disease_id=1 AND 1=1): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}
}

func TestIndexOBO(t *testing.T) {
	query := bleve.NewQueryStringQuery("Name:head")
	// query := bleve.NewTermQuery("head")
	// query.SetField("Name")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)

	if err != nil {
		fmt.Println(err)
	}

	// docID := searchResults.Hits[0].ID
	// doc, err := index.Document(docID)
	fmt.Println(searchResults)
}
