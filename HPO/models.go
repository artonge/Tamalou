package HPO

import (
	"fmt"

	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve/document"
)

type hpoSQLiteStruct struct {
	DiseaseDB    string
	DiseaseID    string
	DiseaseLabel string
	SignID       string
}

type hpoOBOStruct struct {
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

// GetID - To comply with the Indexable interface
func (obo hpoOBOStruct) GetID() string {
	return obo.ID
}

func buildOboStructFromDoc(doc *document.Document) indexing.Indexable {
	var term hpoOBOStruct
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
