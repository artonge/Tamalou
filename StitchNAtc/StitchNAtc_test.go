package stitchnatc

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve"
)

func TestStitchNAtc(t *testing.T) {
	// Create the index if it doesn't exist
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("stitchnatc-search.bleve", mapping)
	if err != nil {
		os.RemoveAll("stitchnatc-search.bleve")
		index, err = bleve.New("stitchnatc-search.bleve", mapping)
		if err != nil {
			log.Fatal(err)
		}
	}

	dicKeg, err := parseKeg()
	if err != nil {
		log.Fatal(err)
	}
	// Open the tsv file
	file, err := os.Open("/media/carl/DATA/Downloads/chemical.sources.v5.0.tsv/chemical.sources.v5.0.tsv")
	if err != nil {
		log.Fatal(err)
	}
	// Create a new Reader to parse the file
	reader := bufio.NewReader(file)
	reader.ReadString('\n')
	err = indexing.IndexDocs(index, func() (indexing.Indexable, error) {
		return nextTerm(reader, dicKeg)
	})
	if err != nil {
		log.Fatal(err)
	}
}
