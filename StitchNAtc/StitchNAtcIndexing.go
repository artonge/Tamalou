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

func indexStitchNAtc() (bleve.Index, error) {
	// Create the index if it doesn't exist
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("stitchnatc-search.bleve", mapping)
	if err != nil {
		return index, fmt.Errorf("Error while creating a new index for stitch & atc: %v", err)
	}

	dicKeg, err := parseKeg()
	if err != nil {
		if err != io.EOF {
			return index, fmt.Errorf("Error while parsing the file br08303.keg: %v", err)
		}
	}
	// Open the tsv file
	file, err := os.Open("datas/chemical.sources.v5.0.tsv")
	if err != nil {
		return index, fmt.Errorf("Error in Stitch&ATC's connector init\n	Error	==> %v", err)
	}
	// Create a new Reader to parse the file
	reader := bufio.NewReader(file)
	reader.ReadString('\n')
	err = indexing.IndexDocs(index, func() (indexing.Indexable, error) {
		for {
			if term, err := nextTerm(reader, dicKeg); term != nil || err != nil {
				return term, err
			}
		}
	})

	return index, err
}

func nextTerm(reader *bufio.Reader, dicStitch map[string]StitchNAtcStruct) (*StitchNAtcStruct, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, fmt.Errorf("error in Stitch and Atc parsing ==> Error: %v", err)
	}
	for line[0] == '#' {
		line, err = reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil, err
			}
			return nil, fmt.Errorf("error in Stitch and Atc parsing ==> Error: %v", err)
		}
	}
	record := strings.Split(line, "\t")
	ID := record[3]
	term := dicStitch[ID[:len(ID)-1]]
	if term.GetID() == "" {
		return nil, nil
	}
	term.Chemical = record[0]
	term.Alias = record[1]
	return &term, nil
}

func parseKeg() (map[string]StitchNAtcStruct, error) {
	file, err := os.Open("datas/br08303.keg")
	if err != nil {
		return nil, err
	}
	dicKeg := make(map[string]StitchNAtcStruct)
	// Create a new Scanner to parse the file
	scanner := bufio.NewScanner(file)

	// Skip file meta datas
	for scanner.Scan() {
		if scanner.Text() == "!" {
			break
		}
	}
	// Continue the file parsing from the last position
	for {
		if !scanner.Scan() {
			if scanner.Err() == nil {
				return dicKeg, io.EOF
			}
			return nil, scanner.Err()
		}
		// Get the line without the Letter and trim white space
		line := strings.TrimSpace(scanner.Text()[1:])
		splitedLine := strings.SplitN(line, " ", 2)
		if len(splitedLine) != 2 {
			continue
		}
		// Init the new term
		doc := StitchNAtcStruct{
			ATCCode:  splitedLine[0],
			ATCLabel: splitedLine[1],
		}
		dicKeg[doc.GetID()] = doc
	}
}
