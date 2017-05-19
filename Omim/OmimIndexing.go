package Omim

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/artonge/Tamalou/indexing"
)

func indexOmim() error {
	dicCsv, err := parseOmimCsv()
	if err != nil {
		return fmt.Errorf("Error while parsing the file omim_onto.csv: %v", err)
	}
	// Open the omim file
	file, err := os.Open("datas/omim/omim.txt")
	if err != nil {
		return fmt.Errorf("Error in Omim's connector init\n	Error	==> %v", err)
	}
	// Create a new Reader to parse the file
	reader := bufio.NewReader(file)

	// Index the CSV
	err = indexing.IndexDocs(index, func() (indexing.Indexable, error) {
		return nextTerm(reader, dicCsv)
	})

	return err
}

// ParseOmimCsv parse Omim_onto.csv
// @return a map with the FieldNumber in key
func parseOmimCsv() (map[string]omimStruct, error) {
	omimCsvFile, err := os.Open("datas/omim/omim_onto.csv")
	if err != nil {
		return nil, err
	}
	dicCsv := make(map[string]omimStruct)
	reader := bufio.NewReader(omimCsvFile)
	headLine, err := reader.ReadString('\n') // first line
	if err != nil {
		return nil, err
	}
	header := strings.Split(headLine, ",") // array with the fields names
	for {
		line, err := reader.ReadString('\n') // current line
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		finalValues := make([]string, len(header)) // array with the final field's values
		fields := strings.Split(line[:len(line)-1], ",")
		if len(fields) > len(header) {
			var fieldVal string
			var isQuote bool
			i := 0
			for _, value := range fields {
				if len(value) <= 0 {
					if !isQuote {
						i++
					}
				} else if value[0] == '"' {
					isQuote = true
					fieldVal = value
				} else if isQuote {
					fieldVal += "," + value
					if value[len(value)-1] == '"' {
						isQuote = false
						finalValues[i] = fieldVal
						i++
					}
				} else {
					finalValues[i] = value
					i++
				}
			}
		} else if len(fields) == len(header) {
			finalValues = fields
		} else {
			err = errors.New("Incorrect Document")
			return nil, err
		}
		id := finalValues[0][len(fields[0])-6:]
		doc := omimStruct{
			FieldNumber:          id,
			FieldDeseaseName:     finalValues[1],
			FieldObsolete:        finalValues[4] == "true",
			FieldCUI:             finalValues[5],
			FieldSemanticTypes:   finalValues[6],
			FieldDiseaseSynonyms: finalValues[2],
		}
		dicCsv[id] = doc
	}
	return dicCsv, nil
}

func nextTerm(reader *bufio.Reader, dicCsv map[string]omimStruct) (*omimStruct, error) {
	var document omimStruct
	var currentField, data string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if len(line) > 6 && line[1:6] == "FIELD" { // begin a new field
			if currentField == "CS" {
				document.FieldSymptome = data
				data = ""
			}
			currentField = line[8 : len(line)-1] // get the name of the current field
		} else if len(line) > 7 && line[1:7] == "RECORD" { // begin a new record
			if document.FieldNumber != "" {
				return &document, nil
			}
		} else {
			switch currentField {
			case "NO":
				document = dicCsv[line[:len(line)-1]]
			case "TX":
				document.FieldDescription += line
			case "CS":
				data += line
			}
		}
	}
	return &document, io.EOF
}
