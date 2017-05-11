package HPO

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
}
