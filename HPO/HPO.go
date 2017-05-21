package HPO

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
	fmt.Println("Indexing obo...")
	index, err = indexing.InitIndex("obo-search.bleve")
	if err != nil {
		fmt.Println("Error while initing obo index:\n	Error ==> ", err)
	}
	err = indexOBO()
	if err != nil {
		fmt.Println("Error while indexing obo:\n	Error ==> ", err)
	}
	fmt.Println("Obo index done")
}

// QueryHPOAsync - Return array of diseases from the HPO databases in channels
func QueryHPOAsync(query Queries.ITamalouQuery, diseaseChanel chan []*Models.Disease, errorChanel chan error) {
	diseasesArray, err := QueryHPO(query)
	diseaseChanel <- diseasesArray
	errorChanel <- err
}

// QueryHPO - Return array of diseases from the HPO databases
func QueryHPO(query Queries.ITamalouQuery) ([]*Models.Disease, error) {
	switch query.Type() {
	case "or":
	case "and":
		var mergeDiseases []*Models.Disease
		for _, child := range query.Children() {
			diseases, err := QueryHPO(child)
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
		diseasesHPIDs, err := QueryIndex(query)
		if err != nil {
			return nil, fmt.Errorf("Error in QueryHPO when querying the index: %v", err)
		}
		return SQLiteQuery(diseasesHPIDs)
	}

	return nil, nil
}
