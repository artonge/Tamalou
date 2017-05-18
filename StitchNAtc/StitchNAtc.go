package stitchnatc

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
)

var index bleve.Index

func init() {
	fmt.Println("Indexing stitch & atc file...")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory:\n Error ==> ", err, pwd)
	}
	err = os.RemoveAll(pwd + "/stitchnatc-search.bleve")
	if err != nil {
		fmt.Println("Error while removing old stitch & atc index:\n Error ==> ", err)
	}
	index, err = indexStitchNAtc()
	if err != nil {
		fmt.Println("Error while indexing stitch & atc file:\n Error ==> ", err)
	}
	fmt.Println("stitch & atc file indexed.")
}
