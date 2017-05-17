package HPO

import (
	"fmt"
	"os"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
	"github.com/blevesearch/bleve"
)

var index bleve.Index

func init() {
	fmt.Println("Indexing obo file...")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory:\n Error ==> ", err, pwd)
	}
	err = os.RemoveAll(pwd + "/obo-search.bleve")
	if err != nil {
		fmt.Println("Error while removing old obo index:\n Error ==> ", err)
	}
	index, err = indexOBO()
	if err != nil {
		fmt.Println("Error while indexing obo file:\n Error ==> ", err)
	}
	fmt.Println("Obo file indexed.")
}

// Return array of diseases from the HPO databases
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
		diseasesHP_IDs, err := QueryIndex(query)
		if err != nil {
			return nil, fmt.Errorf("Error in QueryHPO when querying the index: %v", err)
		}
		return SQLiteQuery(diseasesHP_IDs)
	}

	return nil, nil
}
