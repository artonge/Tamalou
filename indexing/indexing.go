package indexing

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/artonge/Tamalou/Queries"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
)

func InitIndex(indexFile string) (bleve.Index, error) {
	// Get path of the index file
	pwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Error while getting current working directory:\n	Index file ==> %v\n	Error ==> %v", indexFile, err)
	}

	// Remove the old index
	err = os.RemoveAll(pwd + "/" + indexFile)
	if err != nil {
		return nil, fmt.Errorf("Error while removing old omim index:\n	Index file ==> %v\n	Error ==> %v", indexFile, err)
	}

	// Create a nex index file
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("omim-search.bleve", mapping)
	if err != nil {
		return nil, fmt.Errorf("Error while creating a new index for omim:\n	Index file ==> %v\n	Error ==> %v", indexFile, err)
	}

	return index, nil
}

func IndexDocs(index bleve.Index, nextDoc func() (Indexable, error)) error {
	batch := index.NewBatch()
	batchCount := 100
	// Loop through the Docs with a custom passed function (nextDoc)
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

		// I batch is full, write it to the index
		if batchCount == 0 {
			err = index.Batch(batch)
			if err != nil {
				return fmt.Errorf("Error while indexing batch:\n	Error ==> %v", err)
			}
			batch = index.NewBatch()
			// batch.Reset()
			batchCount = 100

			// REMOVE - This stops the indexing at 100 Docs !!
			fmt.Println("REMOVE ME")
			return nil
		}
	}

	// flush the last batch
	if batchCount > 0 {
		err := index.Batch(batch)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func SearchQuery(index bleve.Index, query Queries.ITamalouQuery, buildIndexable func(*document.Document) Indexable) ([]Indexable, error) {
	strQuery := Queries.BuildIndexQuery(query)
	indexQuery := bleve.NewQueryStringQuery(strQuery)
	search := bleve.NewSearchRequest(indexQuery)
	searchResults, err := index.Search(search)
	if err != nil {
		return nil, err
	}
	var results []Indexable
	for _, hit := range searchResults.Hits {
		doc, err := index.Document(hit.ID)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, buildIndexable(doc))
	}
	return results, nil
}
