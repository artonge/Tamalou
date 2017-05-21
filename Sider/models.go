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
)
