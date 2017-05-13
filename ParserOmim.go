package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// OmimStruct structure du schema global des données d'omim
type OmimStruct struct {
	FieldNumber          string
	FieldDeseaseName     string
	FieldDescription     string
	FieldSymptome        string
	FieldObsolete        bool
	FieldCUI             string
	FieldSemanticTypes   string
	FieldDiseaseSynonyms string
}

var omimTxtFilePath = "datas/omim/omim.txt"
var omimCsvFilePath = "datas/omim/omim_onto.csv"

// ParseOmimCsv parse le fichier Omim_onto.csv
// retourne un dictionnaire avec comme clé l'ID de l'élément et en valeur la structure correspondant à l'enregistrement
func ParseOmimCsv() map[string]OmimStruct {
	omimCsvFile, err := os.Open(omimCsvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	dicCsv := make(map[string]OmimStruct)
	reader := bufio.NewReader(omimCsvFile)
	headLine, err := reader.ReadString('\n') // première ligne du csv avec le nom des champs
	if err != nil {
		log.Fatal(err)
	}
	header := strings.Split(headLine, ",") // tableau avec le nom des champs
	for {
		line, err := reader.ReadString('\n') // ligne courante du document
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		finalValues := make([]string, len(header))       // tableau avec les valeurs finales
		fields := strings.Split(line[:len(line)-1], ",") // tableau avec les valeurs de la ligne avec potentiellement des champs dans plusieurs cases du tableau
		if len(fields) > len(header) {
			var fieldVal string
			var isQuote bool // si la valeur est dans un champs entre quotes
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
			log.Fatal("the document is incorrect")
		}
		id := finalValues[0][len(fields[0])-6:]
		doc := OmimStruct{
			FieldNumber:          id,
			FieldDeseaseName:     finalValues[1],
			FieldObsolete:        finalValues[4] == "true",
			FieldCUI:             finalValues[5],
			FieldSemanticTypes:   finalValues[6],
			FieldDiseaseSynonyms: finalValues[2]}
		dicCsv[id] = doc
	}
	return dicCsv
}

// ParseOmim : parse le fichier texte omim.txt
func ParseOmim(ch chan OmimStruct, dicCsv map[string]OmimStruct) {
	fmt.Println("debut du parcours de Omim")
	omimTxtFile, err := os.Open(omimTxtFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer omimTxtFile.Close()

	reader := bufio.NewReader(omimTxtFile)
	var document OmimStruct
	var currentField, data string
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(line) > 6 && line[1:6] == "FIELD" { // recupere le nom du nouveau champs
			if currentField == "CS" {
				document.FieldSymptome = data
				data = ""
			}
			currentField = line[8 : len(line)-1]
		} else if len(line) > 7 && line[1:7] == "RECORD" { // commence un nouvel enregistrement
			if document.FieldNumber != "" {
				ch <- document
				i++
				if i%1000 == 0 {
					fmt.Printf("%d record send to indexation\n", i)
					break
				}
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
	close(ch)
}
