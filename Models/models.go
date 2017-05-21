package Models

type (
	// Disease - Global disease structure
	Disease struct {
		Name    string `json:"name"`
		OMIMID  string
		OrphaID float64
		UMLSID  string
		Score   int      `json:"score"`
		Sources []string `json:"sources"`
	}

	// Drug - Global drug structure
	Drug struct {
		Name        string `json:"name"`
		CUI         string
		SideEffects []SideEffect
	}

	// SideEffect - Side effect of a drug
	SideEffect struct {
		SideEffectName      string
		Placebo             string
		Frequency           string
		FrequencyLowerBound string
		FrequencyUpperBound string
		Matched             string
	}

	// Symptom - Global symptom structure
	Symptom struct {
		Name  string
		HPOID string
	}
)
