package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/**
un record = un document
un document est composé de plusieurs champs (faire une map je pense)
faire en sorte d'envoyer un document a l'indexation quand il a fini d'être parser

utiliser les champs du schema global
parser en plus le champs CS
**/

// OmimStruct structure du schema global des données d'omim
type OmimStruct struct {
	fieldNumber          string
	fieldDeseaseName     string
	fieldDescription     string
	fieldSymptome        []string
	fieldObsolete        string
	fieldCUI             string
	fieldSemanticTypes   string
	fieldDiseaseSynonyms string
}

var omimTxtFilePath = "/home/carl/Documents/Tamalou/omim/omim.txt"
var omimCsvFilePath = "~/Documents/Tamalou/omim/omim_onto.csv"

// ParseOmim : parse le fichier texte omim.txt
func ParseOmim(ch chan OmimStruct) {
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
		if len(line) > 6 && line[1:6] == "FIELD" {
			if currentField == "CS" {
				parseCS(&document, data)
				data = ""
			}
			currentField = line[8 : len(line)-1]
			//fmt.Println(document)
		} else if len(line) > 7 && line[1:7] == "RECORD" {
			if document.fieldNumber != "" {
				ch <- document
			}
			i++
			if i%10 == 0 {
				fmt.Println(i, "record save")
			}
			document = OmimStruct{}
		} else {
			switch currentField {
			case "NO":
				document.fieldNumber = line
			case "TX":
				document.fieldDescription += line
			case "CS":
				data += line
			}
		}
		//fmt.Printf("line: %q\n", line[:len(line)])
	}
	close(ch)
}

func parseCS(document *OmimStruct, data string) {
	document.fieldSymptome = append(document.fieldSymptome, data)
	//fmt.Println(document)
}
