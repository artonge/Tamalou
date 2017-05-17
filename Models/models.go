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
		CUI  string
	}

	Symptom struct {
		Name  string
		HPOID string
	}
)
