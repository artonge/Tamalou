package sider

type (
	Meddra struct {
		CUI         string
		ConceptType string
		MeddraID    int
		Label       string
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
