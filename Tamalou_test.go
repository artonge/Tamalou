package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBleve(t *testing.T) {
	// 	err := removeContents(*IndexPath)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	err = os.Remove(*IndexPath)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	Index()
}

func TestBleveRequest(t *testing.T) {
	// i, err := bleve.Open(*IndexPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// query := bleve.NewMatchQuery("131750")
	// search := bleve.NewSearchRequest(query)
	// searchResults, err := i.Search(search)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(searchResults)
}

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
