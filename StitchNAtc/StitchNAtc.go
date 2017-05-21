package stitchnatc

import (
	"fmt"
	"os"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve"
)

var index bleve.Index

func init() {
	fmt.Println("Indexing stitch & atc file...")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory:\n Error ==> ", err, pwd)
	}
	err = os.RemoveAll(pwd + "/stitchnatc-search.bleve")
	if err != nil {
		fmt.Println("Error while removing old stitch & atc index:\n Error ==> ", err)
	}
	index, err = indexStitchNAtc()
	if err != nil {
		fmt.Println("Error while indexing stitch & atc file:\n Error ==> ", err)
	}
	fmt.Println("stitch & atc file indexed.")
}

func GetChemicalFromID(stitchIDs []string) ([]*Models.Drug, error) {
	var queryString string
	for _, id := range stitchIDs {
		newID := id[:3] + "m" + id[4:]
		if queryString == "" {
			queryString = newID
		} else {
			queryString += " " + newID
		}
	}
	indexQuery := bleve.NewQueryStringQuery(queryString)
	search := bleve.NewSearchRequest(indexQuery)
	searchResults, err := index.Search(search)
	if err != nil {
		return nil, fmt.Errorf("Error querying index\n	Error ==> %v\n	Index ==> %v", err, index)
	}
	var results []StitchNAtcStruct
	for _, hit := range searchResults.Hits {
		doc, err := index.Document(hit.ID)
		if err != nil {
			return nil, fmt.Errorf("Error building indexable\n	Error ==> %v\n	Index ==> %v", err, index)
		}
		results = append(results, BuildStitchNAtcStructFromDoc(doc).(StitchNAtcStruct))
	}
	var drugArray []*Models.Drug

	for _, r := range results {
		drug := Models.Drug{
			CUI:  r.Chemical,
			Name: r.ATCLabel,
		}
		drugArray = append(drugArray, &drug)
	}

	return drugArray, nil
}

func QueryStitchIndex(query Queries.ITamalouQuery) ([]*Models.Drug, error) {
	switch query.Type() {
	case "or":
	case "and":
		var mergeDrug []*Models.Drug
		for _, child := range query.Children() {
			drug, err := QueryStitchIndex(child)
			if err != nil {
				return nil, err
			}
			if len(mergeDrug) > 0 {
				mergeDrug = append(mergeDrug, drug...)
			} else {
				mergeDrug = drug
			}
		}
		return mergeDrug, nil
	default:
		results, err := indexing.QueryIndex(index, query, BuildStitchNAtcStructFromDoc)
		if err != nil {
			return nil, fmt.Errorf("Error while querying omim's index\n	Error ==> %v", err)
		}

		var drugArray []*Models.Drug

		for _, r := range results {
			drug := Models.Drug{
				Name: r.(StitchNAtcStruct).ATCLabel,
			}
			drugArray = append(drugArray, &drug)
		}

		return drugArray, nil
	}
	return nil, nil
}
