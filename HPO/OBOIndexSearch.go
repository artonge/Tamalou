package HPO

import (
	"github.com/artonge/Tamalou/Queries"
	"github.com/artonge/Tamalou/indexing"
)

// Query obo index
// Get IDs of terms matching the query (name or synonymes)
func QueryIndex(query Queries.ITamalouQuery) ([]string, error) {
	results, err := indexing.SearchQuery(index, query, BuildOboStructFromDoc)
	if err != nil {
		return nil, err
	}

	var diseasesIDs []string

	for _, res := range results {
		diseasesIDs = append(diseasesIDs, res.GetID())
	}

	return diseasesIDs, nil
}
