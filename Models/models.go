package Models

type (
	Disease struct {
		Name    string `json:"name"`
		OMIMID  string
		OrphaID float64
		UMLSID  string
	}

	Drug struct {
		Name string `json:"name"`
		CUI  string
	}

	Symptom struct {
		Name  string
		HPOID string
	}
)
