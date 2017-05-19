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
	viewResponse struct {
		TotalRows int          `json:"total_rows"`
		Offset    int          `json:"offset"`
		Rows      []viewResult `json:"rows,omitempty"`
	}

	viewResult struct {
		ID    string                 `json:"id"`
		Key   interface{}            `json:"key"`
		Value map[string]interface{} `json:"value"`
	}
	//
	// GetDiseasesResults struct {
	// 	Id    string  `json:"id"`
	// 	Key   int     `json:"key"`
	// 	Value disease `json:"value"`
	// }
)

func (result *viewResult) String() string {
	return result.Value["disease"].(map[string]interface{})["Name"].(map[string]interface{})["text"].(string)
}

// Value.Disease.Name.Text
