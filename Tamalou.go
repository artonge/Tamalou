package main

import (
	"github.com/artonge/Tamalou/HPO"
	"github.com/artonge/Tamalou/Models"
	orpha "github.com/artonge/Tamalou/Orpha"
	"github.com/artonge/Tamalou/Queries"
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Query string `cli:"q,query" usage:"ventre AND tete OR hand"`
}

func main() {
	// startCLI()
	startServer()
}

func fetchDiseases(query Queries.ITamalouQuery) ([]*Models.Disease, error) {
	// Fetch diseases
	// ORPHA
	resultsOrpha, err := orpha.Query(query)
	if err != nil {
		return nil, err
	}
	// HPO
	resultsHPO, err := HPO.QueryHPO(query)
	if err != nil {
		return nil, err
	}
	results := Models.Merge(resultsOrpha, resultsHPO, "or")
	// ---

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
			if d.Name == df.Name {
				contains = true
				break
			}
		}

		// If filteredDiseases doesn't contains the disease, add it
		if !contains {
			filteredDiseases = append(filteredDiseases, d)
		}
	}

	return filteredDiseases
}
