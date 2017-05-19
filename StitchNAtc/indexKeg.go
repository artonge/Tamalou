package stitchnatc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve"
)

// Index the keg file
func indexKeg() error {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("obo-search.bleve", mapping)
	if err != nil {
		return fmt.Errorf("Error while creating a new index for obo: %v", err)
	}

	// Open the obo file
	file, err := os.Open("datas/hpo/hp.obo")
	if err != nil {
		return fmt.Errorf("Error in HPO's obo connector init\n	Error	==> %v", err)

	}
	// Create a new Scanner to parse the file
	scanner := bufio.NewScanner(file)

	// Skip file meta datas
	for scanner.Scan() {
		if scanner.Text() == "!" {
			break
		}
	}

	err = indexing.IndexDocs(index, func() (indexing.Indexable, error) {
		// Continue the file parsing from the last position
		if !scanner.Scan() {
			if scanner.Err() == nil {
				return nil, io.EOF
			} else {
				return nil, scanner.Err()
			}
		}
		// Get the line without the Letter and trim white space
		line := strings.TrimSpace(scanner.Text()[1:])
		splitedLine := strings.SplitN(line, "  ", 2)

		// Init the new term
		doc := new(KegDocument)
		doc.ID = splitedLine[0]
		doc.Name = splitedLine[1]

		return doc, nil
	})

	return err
}
