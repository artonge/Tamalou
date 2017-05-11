package orpha

// Orpha struct
type (
	nameStruct struct {
		Lang string `json:"lang"`
		Text string `json:"text"`
	}

	disease struct {
		ID          string     `json:"id"`
		OrphaNumber int        `json:"OrphaNumber"`
		Name        nameStruct `json:"Name"`
	}

	clinicalSign struct {
		ID   string     `json:"id"`
		Name nameStruct `json:"Name"`
	}

	freqStruct struct {
		ID   string     `json:"id"`
		Name nameStruct `json:"Name"`
	}

	dataStruct struct {
		SignFreq freqStruct `json:"signFreq"`
	}

	diseaseByClinicalSign struct {
		Disease      disease      `json:"disease"`
		ClinicalSign clinicalSign `json:"clinicalSign"`
		Data         dataStruct   `json:"data"`
	}
)

// CouchDB structs
type (
	ViewResponse struct {
		TotalRows int          `json:"total_rows"`
		Offset    int          `json:"offset"`
		Rows      []ViewResult `json:"rows,omitempty"`
	}

	ViewResult struct {
		Id    string                `json:"id"`
		Key   string                `json:"key"`
		Value diseaseByClinicalSign `json:"value"`
	}
)
