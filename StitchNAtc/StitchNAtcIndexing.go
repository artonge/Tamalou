package stitchnatc

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve"
)

func indexStitchNAtc() (bleve.Index, error) {
	// Create the index if it doesn't exist
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("stitchnatc-search.bleve", mapping)
	if err != nil {
		return index, fmt.Errorf("Error while creating a new index for omim: %v", err)
	}

	dicKeg := ParseKeg()
	if err != nil {
		return index, fmt.Errorf("Error while parsing the file omim_onto.csv: %v", err)
	}
	// Open the omim file
	file, err := os.Open("../datas/omim/omim.txt")
	if err != nil {
		return index, fmt.Errorf("Error in Omim's connector init\n	Error	==> %v", err)
	}
	// Create a new Reader to parse the file
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.Comment = '#'

	err = indexing.IndexDocs(index, func() (indexing.Indexable, error) {
		return nextTerm(reader, dicKeg)
	})

	return index, err
}

func ParseKeg() map[string]StitchNAtcStruct {
	return nil
}

func nextTerm(reader *csv.Reader, dicStitch map[string]StitchNAtcStruct) (*StitchNAtcStruct, error) {
	record, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, fmt.Errorf("error in Stitch and Atc parsing ==> Error: %v", err)
	}
	ID := record[0]
	term := dicStitch[ID]
	return &term, nil
}
