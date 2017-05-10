package main

import (
	"log"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
)

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
	dANDcs struct {
		Disease      disease      `json:"disease"`
		ClinicalSign clinicalSign `json:"clinicalSign"`
		Data         dataStruct   `json:"data"`
	}
	ViewResponse struct {
		TotalRows int          `json:"total_rows"`
		Offset    int          `json:"offset"`
		Rows      []ViewResult `json:"rows,omitempty"`
	}
	ViewResult struct {
		Id    string `json:"id"`
		Key   string `json:"key"`
		Value dANDcs `json:"value"`
	}
)

func CouchDBConnection() {

	conn, err := couchdb.NewConnection(
		"couchdb.telecomnancy.univ-lorraine.fr",
		80,
		time.Duration(35000*time.Millisecond),
	)

	if err != nil {
		log.Fatal("Error1: ", err)
	}

	DB := conn.SelectDB("orphadatabase", nil)

	if err != nil {
		log.Fatal("Error2: ", err)
	}

	results := ViewResponse{}
	// results := make(map[string]interface{})

	err = DB.GetView("clinicalsigns", "GetDiseaseByClinicalSign", &results, nil)

	if err != nil {
		log.Fatal("Error3: ", err)
	}
}
