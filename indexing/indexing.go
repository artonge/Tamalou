package indexing

import (
	"fmt"
	"io"
	"os"

	"github.com/artonge/Tamalou/Queries"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
)

// InitIndex -
func InitIndex(indexFile string) (bleve.Index, error) {
	err := removeIndex(indexFile)
	if err != nil {
		return nil, fmt.Errorf("Error while deleting old index:\n	Index file ==> %v\n	Error ==> %v", indexFile, err)
	}
	// Create a nex index file
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexFile, mapping)
	if err != nil {
		return nil, fmt.Errorf("Error while creating a new index:\n	Index file ==> %v\n	Error ==> %v", indexFile, err)
	}
	return index, nil
}

func removeIndex(indexFile string) error {
	// Get path of the index file
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error while getting current working directory:\n	Index file ==> %v\n	Error ==> %v", indexFile, err)
	}

	// Remove the old index
	err = os.RemoveAll(pwd + "/" + indexFile)
	if err != nil {
		return fmt.Errorf("Error while removing old index:\n	Index file ==> %v\n	Error ==> %v", indexFile, err)
	}

	return nil
}

// OpenIndex -
func OpenIndex(indexFile string) (bleve.Index, error) {
	index, err := bleve.Open(indexFile)
	if err == bleve.ErrorIndexPathDoesNotExist {
		index, err = InitIndex(indexFile)
		if err != nil {
			return nil, fmt.Errorf("Error while initing index file:\n	Index file ==> %v\n	Error ==> %v", indexFile, err)
		}
	}

	return index, err
}

// IndexDocs - index some Docs
// nextDoc allow you to parse your file progressivly
func IndexDocs(index bleve.Index, nextDoc func() (Indexable, error)) error {
	batch := index.NewBatch()
	batchCount := 100
	// Loop through the Docs with a custom passed function (nextDoc)
	nbDone := 0
	for {
		doc, err := nextDoc()
		if err != nil {
			// End of the file
			if err == io.EOF {
				break
			}
			// Other errors
			return fmt.Errorf("Error while parsing file\n	==> %v\n	==> %v", err, doc)
		}
		// Add doc to the current batch
		err = batch.Index(doc.GetID(), doc)
		if err != nil {
			return fmt.Errorf("Error while adding doc to the batch\n	==> %v\nDoc:	==> %v\nBatch	==>%v", err, doc, batch)
		}
		// Decrement batchCount
		batchCount--

		// If batch is full, write it to the index
		if batchCount == 0 {
			err = index.Batch(batch)
			if err != nil {
				return fmt.Errorf("Error while indexing batch:\n	Error ==> %v", err)
			}
			batch = index.NewBatch()
			// batch.Reset()
			batchCount = 100

			// REMOVE - This stops the indexing at 100 Docs !!
			//fmt.Println("REMOVE ME")
			//return nil
		}

		nbDone++
		if nbDone%100 == 0 {
			fmt.Println("Done :", nbDone)
		}
	}
	fmt.Println("Over")

	// flush the last batch
	if batchCount > 0 {
		err := index.Batch(batch)
		if err != nil {
			return fmt.Errorf("Error while indexing batch:\n	Error ==> %v", err)
		}
	}

	return nil
}

// QueryIndex - Make a query on the given index
// Use the passed function to build the document into the wished struct
func QueryIndex(index bleve.Index, query Queries.ITamalouQuery, buildIndexable func(*document.Document) Indexable) ([]Indexable, error) {
	strQuery := Queries.BuildIndexQuery(query)
	indexQuery := bleve.NewQueryStringQuery(strQuery)
	search := bleve.NewSearchRequest(indexQuery)
	searchResults, err := index.Search(search)
	if err != nil {
		return nil, fmt.Errorf("Error querying index\n	Error ==> %v\n	Index ==> %v", err, index)
	}
	var results []Indexable
	for _, hit := range searchResults.Hits {
		doc, err := index.Document(hit.ID)
		if err != nil {
			return nil, fmt.Errorf("Error building indexable\n	Error ==> %v\n	Index ==> %v", err, index)
		}
		results = append(results, buildIndexable(doc))
	}
	return results, nil
}
