package HPO

import (
	"fmt"

	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve/document"
)

type HPOSQLiteStruct struct {
	DiseaseDB    string
	DiseaseID    string
	DiseaseLabel string
	SignID       string
}

type HPOOBOStruct struct {
	ID         string
	AltIDs     []string
	Name       string
	Comment    string
	Definition string
	Synonymes  []string
	Xrefs      []string
	IsA        string
	Obsolete   bool
	count      int
}

func (obo HPOOBOStruct) GetID() string {
	return obo.ID
}

func BuildOboStructFromDoc(doc *document.Document) indexing.Indexable {
	var term HPOOBOStruct
	for _, field := range doc.Fields {
		val := field.Value()
		switch field.Name() {
		case "ID":
			term.ID = string(val)
		case "AltIDs":
			term.AltIDs = []string{string(val)}
		case "Name":
			term.Name = string(val)
		case "Definition":
			term.Definition = string(val)
		case "Comment":
			term.Comment = string(val)
		case "Synonymes":
			term.Synonymes = []string{string(val)}
		case "Xrefs":
			term.Xrefs = []string{string(val)}
		case "IsA":
			term.IsA = string(val)
		case "Obsolete": // use with consider ?
			term.Obsolete = true
		case "created_by", "property_value", "replaced_by", "creation_date", "subset", "is_anonymous", "consider":
		default:
			fmt.Println("Warning: Unexpected field <", field.Name(), "> during obo index query")
		}
	}
	return term
}
