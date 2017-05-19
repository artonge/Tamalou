package Omim

import (
	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve/document"
)

type omimStruct struct {
	FieldNumber          string
	FieldDeseaseName     string
	FieldDescription     string
	FieldSymptome        string
	FieldObsolete        bool
	FieldCUI             string
	FieldSemanticTypes   string
	FieldDiseaseSynonyms string
}

// GetID - To comply with the Indexable interface
func (me omimStruct) GetID() string {
	return me.FieldNumber
}

func buildOmimStructFromDoc(doc *document.Document) indexing.Indexable {
	me := omimStruct{}
	for _, field := range doc.Fields {
		val := field.Value()
		switch field.Name() {
		case "FieldCUI":
			me.FieldCUI = string(val)
		case "FieldDescription":
			me.FieldDescription = string(val)
		case "FieldDeseaseName":
			me.FieldDeseaseName = string(val)
		case "FieldDiseaseSynonyms":
			me.FieldDiseaseSynonyms = string(val)
		case "FieldNumber":
			me.FieldNumber = string(val)
		case "FieldObsolete":
			me.FieldObsolete = string(val) != "F"
		case "FieldSemanticTypes":
			me.FieldSemanticTypes = string(val)
		case "FieldSymptome":
			me.FieldSymptome = string(val)
		}
	}
	return me
}
