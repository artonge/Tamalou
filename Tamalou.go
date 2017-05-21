package main

import (
	"github.com/artonge/Tamalou/HPO"
	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Omim"
	"github.com/artonge/Tamalou/Orpha"
	"github.com/artonge/Tamalou/Queries"
)

func main() {
	// startCLI()
	startServer()
}

func fetchDiseases(query Queries.ITamalouQuery, diseaseChanel chan []*Models.Disease, errorChanel chan error) {
	orphaChanel := make(chan []*Models.Disease, 1)
	omimChanel := make(chan []*Models.Disease, 1)
	hpoChanel := make(chan []*Models.Disease, 1)
	omimFromHPOChanel := make(chan []*Models.Disease, 1)
	orphaFromHPOChanel := make(chan []*Models.Disease, 1)

	// Fetch diseases
	// HPO
	go HPO.QueryHPOAsync(query, hpoChanel, errorChanel)
	// ORPHA
	go orpha.QueryAsync(query, orphaChanel, errorChanel)
	// OMIM
	go Omim.QueryOmimIndexAsync(query, omimChanel, errorChanel)
	// Orpha and OMIM from HPO
	var orphaIDs []float64
	var omimIDs []string
	resultHPO := <-hpoChanel
	for _, d := range resultHPO {
		if d.OMIMID != "" {
			omimIDs = append(omimIDs, d.OMIMID)
		}
		if d.OrphaID != 0 {
			orphaIDs = append(orphaIDs, d.OrphaID)
		}
	}
	go orpha.GetDiseasesFromIDsAsync(orphaIDs, orphaFromHPOChanel, errorChanel)
	go Omim.DiseasesFromIDsAsync(omimIDs, omimFromHPOChanel, errorChanel)

	results := make([]*Models.Disease, 0)
	// Merge results
	results = Models.Merge(results, <-orphaChanel, "or")
	results = Models.Merge(results, resultHPO, "or")
	results = Models.Merge(results, <-omimChanel, "or")
	results = Models.Merge(results, <-orphaFromHPOChanel, "or")
	results = Models.Merge(results, <-omimFromHPOChanel, "or")

	// Return filtered results (remove double apparition)
	diseaseChanel <- filterDiseases(results)
}

func fetchDrugs(query Queries.ITamalouQuery) ([]*Models.Drug, error) {
	// Fetch drugs
	results := make([]*Models.Drug, 0)

	return results, nil
}

func filterDiseases(diseaseArray []*Models.Disease) []*Models.Disease {
	var filteredDiseases []*Models.Disease

	// Loop throught the diseaseArray
	for _, d := range diseaseArray {
		// Check that filteredDiseases doesn't contains the current disease
		contains := false
		for _, df := range filteredDiseases {
			if (d.OMIMID != "" && d.OMIMID == df.OMIMID) || (d.OrphaID != 0 && d.OrphaID == df.OrphaID) {
				// Increment Score of the disease
				// ==> better score when the results comes from multiple sources
				df.Score++
				df.Sources = append(df.Sources, d.Sources...)
				contains = true
				break
			}
		}

		// If filteredDiseases doesn't contains the disease, add it
		if !contains {
			d.Score = 1
			filteredDiseases = append(filteredDiseases, d)
		}
	}

	return filteredDiseases
}
