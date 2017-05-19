package stitchnatc

import (
	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
	"github.com/artonge/Tamalou/indexing"
)

func QueryKegIndex(query Queries.ITamalouQuery) ([]*Models.Drug, error) {

	results, err := indexing.SearchQuery(index, query, BuildKegStructFromDoc)
	if err != nil {
		return nil, err
	}

	var diseasesIDs []string

	for _, res := range results {
		diseasesIDs = append(diseasesIDs, res.GetID())
	}

	return nil, nil
}
