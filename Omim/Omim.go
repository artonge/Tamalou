package Omim

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
)

var index bleve.Index

func init() {
	fmt.Println("Indexing omim file...")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory:\n Error ==> ", err, pwd)
	}
	err = os.RemoveAll(pwd + "/omim-search.bleve")
	if err != nil {
		fmt.Println("Error while removing old omim index:\n Error ==> ", err)
	}
	index, err = indexOmim()
	if err != nil {
		fmt.Println("Error while indexing Omim file:\n Error ==> ", err)
	}
	fmt.Println("Omim file indexed.")
}
