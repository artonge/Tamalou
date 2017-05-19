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

func fetchDiseases(query Queries.ITamalouQuery) ([]*Models.Disease, error) {
	// Fetch diseases
	// HPO
	resultsHPO, err := HPO.QueryHPO(query)
	if err != nil {
		return nil, err
	}
	// ORPHA
	resultsOrpha, err := orpha.Query(query)
	if err != nil {
		return nil, err
	}
	// OMIM
	resultsOMIM, err := Omim.QueryOmimIndex(query)
	if err != nil {
		return nil, err
	}
	// Orpha and OMIM from HPO
	var orphaIDs []float64
	var omimIDs []string
	for _, d := range resultsHPO {
		if d.OMIMID != "" {
			omimIDs = append(omimIDs, d.OMIMID)
		}
		if d.OrphaID != 0 {
			orphaIDs = append(orphaIDs, d.OrphaID)
		}
	}
	resultsOrphaFromHPO, err := orpha.GetDiseasesFromIDs(orphaIDs)
	if err != nil {
		return nil, err
	}
	resultsOMIMFromHPO, err := Omim.DiseasesFromIDs(omimIDs)
	if err != nil {
		return nil, err
	}

	// Merge results
	results := Models.Merge(resultsOrpha, resultsHPO, "or")
	results = Models.Merge(results, resultsOMIM, "or")
	results = Models.Merge(results, resultsOrphaFromHPO, "or")
	results = Models.Merge(results, resultsOMIMFromHPO, "or")

	// Return filtered results (remove double apparition)
	return filterDiseases(results), nil
}

func fetchDrugs(query Queries.ITamalouQuery) ([]*Models.Drug, error) {
	// Fetch drugs

	return nil, nil
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
