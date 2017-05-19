package Omim

import (
	"fmt"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve"
)

var index bleve.Index

func init() {
	var err error
	fmt.Println("Indexing omim...")
	index, err = indexing.InitIndex("omim-search.bleve")
	if err != nil {
		fmt.Println("Error while initing omim index:\n	Error ==> ", err)
	}

	err = indexOmim()
	if err != nil {
		fmt.Println("Error while indexing omim's csv:\n	Error ==> ", err)
	}
	fmt.Println("Omim index done")
}

// DiseasesFromIDs -
func DiseasesFromIDs(OMIMIDs []string) ([]*Models.Disease, error) {
	var idQueries []Queries.ITamalouQuery
	for _, id := range OMIMIDs {
		idQueries = append(idQueries, Queries.NewTamalouQuery("id", id, nil))
	}

	query := Queries.NewTamalouQuery("or", "", idQueries)

	results, err := QueryOmimIndex(query)
	if err != nil {
		return nil, fmt.Errorf("Error querying omim with OMIM IDs\n	Error ==> %v", err)
	}

	return results, nil
}

// QueryOmimIndex -
func QueryOmimIndex(query Queries.ITamalouQuery) ([]*Models.Disease, error) {
	switch query.Type() {
	case "or":
	case "and":
		var mergeDiseases []*Models.Disease
		for _, child := range query.Children() {
			diseases, err := QueryOmimIndex(child)
			if err != nil {
				return nil, err
			}
			if len(mergeDiseases) > 0 {
				mergeDiseases = Models.Merge(mergeDiseases, diseases, string(query.Type()))
			} else {
				mergeDiseases = diseases
			}
		}
		return mergeDiseases, nil
	default:
		results, err := indexing.QueryIndex(index, query, buildOmimStructFromDoc)
		if err != nil {
			return nil, fmt.Errorf("Error while querying omim's index\n	Error ==> %v", err)
		}

		var diseaseArray []*Models.Disease

		for _, r := range results {
			var tmpDisease Models.Disease
			tmpDisease.Name = r.(omimStruct).FieldDeseaseName
			tmpDisease.UMLSID = r.(omimStruct).FieldCUI
			tmpDisease.OMIMID = r.(omimStruct).FieldNumber
			tmpDisease.Sources = append(tmpDisease.Sources, "OMIM")
			diseaseArray = append(diseaseArray, &tmpDisease)
		}

		return diseaseArray, nil
	}
	return nil, nil
}
