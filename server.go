package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/artonge/Tamalou/Models"
	"github.com/artonge/Tamalou/Queries"
)

func request(w rest.ResponseWriter, r *rest.Request) {
	// Build the query to an ITamalouQuery
	rawQuery := r.FormValue("request")
	query := Queries.ParseQuery(rawQuery)

	// Prepare chanels
	errorChanel := make(chan error, 1)
	diseasesChanel := make(chan []*Models.Disease, 1)
	// drugsChanel := make(chan []*Models.Drug, 1)

	// Get diseases and drugs
	go fetchDiseases(query, diseasesChanel, errorChanel)
	drugs, err := fetchDrugs(query)
	if err != nil {
		w.WriteJson(err)
	}

	w.WriteJson(map[string]interface{}{
		"diseases": <-diseasesChanel,
		"drugs":    drugs,
	})
}

func startServer() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/request", request),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
