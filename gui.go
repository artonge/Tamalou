package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/artonge/Tamalou/Queries"
)

func index(w rest.ResponseWriter, r *rest.Request) {
	rawQuery := r.FormValue("request")
	query := Queries.ParseQuery(rawQuery)

	// Get diseases and drugs
	diseases, err := fetchDiseases(query)
	if err != nil {
		w.WriteJson(err)
	}
	drugs, err := fetchDrugs(query)
	if err != nil {
		w.WriteJson(err)
	}

	w.WriteJson(map[string]interface{}{
		"diseases": diseases,
		"drugs":    drugs,
	})
}

func startServer() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(index))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
