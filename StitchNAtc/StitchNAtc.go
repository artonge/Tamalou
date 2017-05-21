package stitchnatc

import (
	"fmt"

	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve"
)

var index bleve.Index

// Open the index on startup
func init() {
	var err error
	index, err = indexing.OpenIndex("stitchnatc-search.bleve")
	if err != nil {
		fmt.Println("Error while initing stitchnatc index:\n	Error ==> ", err)
	}
}

// IndexStitchNAtc -
func IndexStitchNAtc() error {
	var err error
	fmt.Println("Indexing stitch & atc file...")
	index, err = indexing.InitIndex("stitchnatc-search.bleve")
	if err != nil {
		return fmt.Errorf("Error while initing stitchnatc index:\n	Error ==> %v", err)
	}
	index, err = indexStitchNAtc()
	if err != nil {
		return fmt.Errorf("Error while indexing stitch & atc file:\n Error ==> %v", err)
	}
	fmt.Println("stitch & atc file indexed.")
	return nil
}

func StitchIdSider2ATC(str string) string {
	return str[:3] + "m" + str[4:]
}

func GetChemicalFromID(drugs []*Models.Drug) error {
	var queryString string
	for _, drug := range drugs {
		drug.STITCH_ID_ATC = StitchIdSider2ATC(drug.STITCH_ID_SIDER)
		if queryString == "" {
			queryString = drug.STITCH_ID_ATC
		} else {
			queryString += " " + drug.STITCH_ID_ATC
		}
	}
	indexQuery := bleve.NewQueryStringQuery(queryString)
	search := bleve.NewSearchRequest(indexQuery)
	searchResults, err := index.Search(search)
	if err != nil {
		// return nil, fmt.Errorf("Error querying index\n	Error ==> %v\n	Index ==> %v", err, index)
		return nil
	}
	var results []StitchNAtcStruct
	for _, hit := range searchResults.Hits {
		doc, err := index.Document(hit.ID)
		if err != nil {
			// return nil, fmt.Errorf("Error building indexable\n	Error ==> %v\n	Index ==> %v", err, index)
			return nil
		}
		results = append(results, BuildStitchNAtcStructFromDoc(doc).(StitchNAtcStruct))
	}
	//var drugArray []*Models.Drug

	// Brute force results
	nbHit := 0
	fmt.Println("Got ", len(results), " results")
	for _, r := range results {
		fmt.Println("Looking for a drug with id ", r.Chemical)
		for _, drug := range drugs {
			fmt.Println("\tIs this one good ? => ", drug.STITCH_ID_ATC)
			if drug.STITCH_ID_ATC == r.Chemical {
				fmt.Println("\t\tYes it is !")
				drug.Name = r.ATCLabel
				nbHit++
				break
			}
		}
	}

	fmt.Println("Nombre de hit ", nbHit, " / ", len(drugs))

	// for _, r := range results {
	// 	drug := Models.Drug{
	// 		CUI:  r.Chemical,
	// 		Name: r.ATCLabel,
	// 	}
	// 	drugArray = append(drugArray, &drug)
	// }

	// return drugArray, nil
	return nil
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
