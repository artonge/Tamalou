package Models

type (
	Disease struct {
		Name    string
		OMIMID  string
		OrphaID float64
		UMLSID  string
	}

	Drug struct {
		Name string
	}

	Symptom struct {
		Name   string
		HPO_ID string
	}
)
