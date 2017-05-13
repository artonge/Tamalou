package HPO

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
)

var index bleve.Index

func init() {
	fmt.Println("Indexing obo file...")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory:\n Error ==> ", err, pwd)
	}
	err = os.RemoveAll(pwd + "/obo-search.bleve")
	if err != nil {
		fmt.Println("Error while removing old obo index:\n Error ==> ", err)
	}
	index, err = indexOBO()
	if err != nil {
		fmt.Println("Error while indexing obo file:\n Error ==> ", err)
	}
	fmt.Println("Obo file indexed.")
}
