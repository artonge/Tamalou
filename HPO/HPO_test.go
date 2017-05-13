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

func TestIndexOBO(t *testing.T) {
	fmt.Println(index.Fields())
	fmt.Println(index.Document("HP:0000003"))
	query := bleve.NewMatchQuery("+AltIDs:HP:0000003 -AltIDs:mustnot")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(searchResults)
}
