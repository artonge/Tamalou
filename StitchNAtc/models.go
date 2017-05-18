package stitchnatc

import (
	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve/document"
)

type StitchNAtcStruct struct {
	CompoundID string
	ATC        string
	label      string
}

func (me StitchNAtcStruct) GetID() string {
	return me.CompoundID
}

type KegDocument struct {
	Name string
	ID   string
}

func (doc KegDocument) GetID() string {
	return doc.ID
}

func BuildKegStructFromDoc(doc *document.Document) indexing.Indexable {
	var kegItem KegDocument
	for _, field := range doc.Fields {
		switch field.Name() {
		case "ID":
			kegItem.ID = string(field.Value())
		case "Name":
			kegItem.Name = string(field.Value())
		}
	}
	return nil
}
