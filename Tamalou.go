package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/blevesearch/bleve"
)

func main() {

}

// Bleve example
func bleveExample() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}

	// index some data
	index.Index("id", data)

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}

// mgo example
type person struct {
	Name  string
	Phone string
}

func mgoExemple() {

	session, err := mgo.Dial("server1.example.com,server2.example.com")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")

	err = c.Insert(
		&person{"Ale", "+55 53 8116 9639"},
		&person{"Cla", "+55 53 8402 8510"},
	)

	if err != nil {
		log.Fatal(err)
	}

	result := person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

type testDocument struct {
	Title string
	Note  string
}
