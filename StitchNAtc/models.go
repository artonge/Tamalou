package stitchnatc

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
