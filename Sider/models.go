package sider

type (
	SideEffect struct {
		SideEffectName      string
		Placebo             string
		Frequency           string
		FrequencyLowerBound string
		FrequencyUpperBound string
		Matched             string
	}

	Meddra struct {
		StitchCompoundId string
		SideEffects      []*SideEffect
	}

	MeddraAllIndications struct {
		StitchCompoundID  string
		CUI               string
		MethodOfDetection string
		ConceptName       string
		MeddraConceptType string
		CUIOfMeddraTerm   string
		MeddraConceptName string
	}

	MeddraAllSe struct {
		StitchCompoundID1 string
		StitchCompoundID2 string
		CUI               string
		MeddraConceptType string
		CUIOfMeddraTerm   string
		SideEffectName    string
	}

	MeddraFreq struct {
		StitchCompoundID1    string
		StitchCompoundID2    string
		CUI                  string
		Placebo              string
		FrequencyDescription string
		FreqLowerBound       string
		FreqUpperBound       string
		MeddraConceptType    string
		MeddraConceptID      string
		SideEffectName       string
	}
)
