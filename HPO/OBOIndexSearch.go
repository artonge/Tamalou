package HPO

import (
	"fmt"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
	"github.com/artonge/Tamalou/indexing"
)

func QueryIndex(query Queries.ITamalouQuery) ([]Models.Disease, error) {

	results, err := indexing.SearchQuery(index, query, BuildOboStructFromDoc)

	fmt.Println(results, err)
	return nil, nil
}
