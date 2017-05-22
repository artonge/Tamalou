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
		Name            string `json:"name"`
		CUI             string
		STITCH_ID_SIDER string `json:"id"`
		STITCH_ID_ATC   string
		SideEffects     []*SideEffect `json:sideEffects`
	}

	// SideEffect - Side effect of a drug
	SideEffect struct {
		SideEffectName      string `json:name`
		Placebo             string `json:placebo`
		Frequency           string `json:freq`
		FrequencyLowerBound string `json:freqlower`
		FrequencyUpperBound string `json:frequpper`
		Matched             string `json:matched`
	}

	// Symptom - Global symptom structure
	Symptom struct {
		Name  string
		HPOID string
	}
)
