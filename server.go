package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/artonge/Tamalou/Queries"
	sider "github.com/artonge/Tamalou/Sider"
)

func request(w rest.ResponseWriter, r *rest.Request) {
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

func getSideEffects(w rest.ResponseWriter, r *rest.Request) {
	drugID := r.FormValue("drugID")

	// Get diseases and drugs
	sideeffects, err := sider.GetSideEffects(drugID)
	if err != nil {
		w.WriteJson(err)
	}

	fmt.Println(drugID)
	fmt.Println(sideeffects)
	w.WriteJson(map[string]interface{}{
		"sideeffects": sideeffects,
	})
}

func startServer() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/request", request),
		rest.Get("/sideeffect", getSideEffects),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	http.Handle("/", http.StripPrefix("", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
