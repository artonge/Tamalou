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
	index, err = indexing.InitIndex("omim-search.bleve")
	if err != nil {
		fmt.Println("Error while initing omim index:\n	Error ==> ", err)
	}

	err = indexOmim()
	if err != nil {
		fmt.Println("Error while indexing omim's csv:\n	Error ==> ", err)
	}
}

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
		results, err := indexing.SearchQuery(index, query, BuildOmimStructFromDoc)
		if err != nil {
			return nil, fmt.Errorf("Error while querying omim's index\n	Error ==> %v", err)
		}

		var diseaseArray []*Models.Disease

		for _, r := range results {
			var tmpDisease Models.Disease
			tmpDisease.Name = r.(OmimStruct).FieldDeseaseName
			tmpDisease.UMLSID = r.(OmimStruct).FieldCUI
			tmpDisease.OMIMID = r.(OmimStruct).FieldNumber
			tmpDisease.Sources = append(tmpDisease.Sources, "OMIM")
			diseaseArray = append(diseaseArray, &tmpDisease)
		}

		return diseaseArray, nil
	}
	return nil, nil
}
