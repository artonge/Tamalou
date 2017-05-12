package HPO

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/blevesearch/bleve"
)

func IndexOBO() (bleve.Index, error) {
	// Set IndexPath
	IndexPath := flag.String("index", "obo-search.bleve", "index path")
	// Open the index
	index, err := bleve.Open(*IndexPath)
	// Create one it the index doesn't exist
	if err == bleve.ErrorIndexPathDoesNotExist {
		log.Printf("Creating new index...")
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(*IndexPath, mapping)
		if err != nil {
			return index, fmt.Errorf("Error while creating a new index for obo: %v", err)
		}
	}

	// Open the obo file
	file, err := os.Open("../datas/hpo/hp.obo")
	if err != nil {
		// return fmt.Errorf("Error in HPO's obo connector init: ", err)
		return index, fmt.Errorf("Error while querying sider (meddra): %v", err)
	}

	// Create a new Scanner to parse the file
	scanner := bufio.NewScanner(file)

	// Loop through the Terms
	for {
		term, err := nextTerm(scanner)
		if err != nil {
			return index, fmt.Errorf("Error while parsing obo file\n	==> %v", err)
		}
		// End of the file
		if term == nil {
			break
		}
		// Add index the term
		index.Index(term.ID, term)
	}

	return index, nil
}
