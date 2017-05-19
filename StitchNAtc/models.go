package stitchnatc

import (
	"github.com/artonge/Tamalou/indexing"
	"github.com/blevesearch/bleve/document"
)

type StitchNAtcStruct struct {
	ATCCode  string
	ATCLabel string
	Chemical string
	Alias    string
}

func (me StitchNAtcStruct) GetID() string {
	return me.ATCCode
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

func BuildStitchNAtcStructFromDoc(doc *document.Document) indexing.Indexable {
	var sitchItem StitchNAtcStruct
	for _, field := range doc.Fields {
		switch field.Name() {
		case "ATCCode":
			sitchItem.ATCCode = string(field.Value())
		case "ATCLabel":
			sitchItem.ATCLabel = string(field.Value())
		case "Chemical":
			sitchItem.Chemical = string(field.Value())
		case "Alias":
			sitchItem.Alias = string(field.Value())
		}
	}
	return nil
}
