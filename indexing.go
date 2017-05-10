package main

import (
	"flag"
	"log"

	"github.com/blevesearch/bleve"
)

var IndexPath = flag.String("index", "omim-search.bleve", "index path")

// TODO: soit trouver le bon analyzer ou créer un analyzer

// voir mapping.go de beer-search
// permet d'analiser le texte
// func buildIndexMapping() (mapping.IndexMapping, error) {
//
// 	textFieldMapping := bleve.NewTextFieldMapping()
// 	textFieldMapping.Analyzer = en.AnalyzerName
//
// 	//	omimMapping := bleve.NewDocumentMapping()
// 	//omimMapping.AddFieldMappingsAt("property", fms)
//
// 	return nil, nil
// }

//Index : ouvre ou creer l'index s'il n'existe pas
func Index() bleve.Index {
	omimIndex, err := bleve.Open(*IndexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		log.Printf("Creating new index...")
		indexMapping := bleve.NewIndexMapping()
		omimIndex, err = bleve.New(*IndexPath, indexMapping)
		if err != nil {
			log.Fatal(err)
		}
		err = indexOmim(omimIndex)
		if err != nil {
			log.Fatal(err)
		}
	}
	return omimIndex
}

func indexOmim(i bleve.Index) error {
	ch := make(chan OmimStruct)
	go ParseOmim(ch)
	for doc := range ch {
		i.Index(doc.fieldNumber, doc)
	}

	return nil
}
