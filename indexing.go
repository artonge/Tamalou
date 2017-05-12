package main

import (
	"flag"
	"log"
	"time"

	"github.com/blevesearch/bleve"
)

var IndexPath = flag.String("index", "omim-search.bleve", "index path")

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
		//err = indexOmimWithoutBatch(omimIndex)
		err = indexOmimBatch(omimIndex)
		if err != nil {
			log.Fatal(err)
		}
	}
	return omimIndex
}

func indexOmimBatch(i bleve.Index) error {
	dicCsv := ParseOmimCsv()    // Load all csv file in memory (dictionary)
	ch := make(chan OmimStruct) // Chanel for multithreading
	startTime := time.Now()
	go ParseOmim(ch, dicCsv) // Parse OMIM to complete csv dictionary
	count := 0
	batch := i.NewBatch()
	batchCount := 0
	batchSize := 100
	for doc := range ch {
		batch.Index(doc.FieldNumber, doc)
		batchCount++
		count++
		if batchCount >= batchSize {
			err := i.Batch(batch)
			if err != nil {
				return err
			}
			indexDuration := time.Since(startTime)
			indexDurationSeconds := float64(indexDuration) / float64(time.Second)
			timePerDoc := float64(indexDuration) / float64(count)
			log.Printf("Indexed %d documents, in %.2fs (average %.2fms/doc)", count, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
			batch = i.NewBatch()
			batchCount = 0
		}
	}
	if batchCount > 0 {
		err := i.Batch(batch)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// func indexOmimWithoutBatch(i bleve.Index) error {
// 	ch := make(chan OmimStruct)
// 	startTime := time.Now()
// 	go ParseOmim(ch)
// 	count := 0
// 	for doc := range ch {
// 		i.Index(doc.FieldNumber, doc)
// 		count++
// 		if count%10 == 0 {
// 			indexDuration := time.Since(startTime)
// 			indexDurationSeconds := float64(indexDuration) / float64(time.Second)
// 			timePerDoc := float64(indexDuration) / float64(count)
// 			log.Printf("Indexed %d documents, in %.2fs (average %.2fms/doc)", count, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
// 		}
// 	}
//
// 	return nil
// }
