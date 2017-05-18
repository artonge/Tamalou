package stitchnatc

type StitchNAtcStruct struct {
	CompoundID string
	ATC        string
	label      string
}

func (me StitchNAtcStruct) GetID() string {
	return me.CompoundID
}
