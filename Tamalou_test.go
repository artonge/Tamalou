package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCouchExample(t *testing.T) {
	// CouchExample()
}

func TestMysqlConnection(t *testing.T) {
	// Uncomment to show the Sider database
	// MysqlConnection()
}

func TestCouchDBConnection(t *testing.T) {
	//CouchDBConnection()
}

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

func TestHPOQuery(t *testing.T) {
	hpoArray, err := HPOQuery("'1'='1'")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	hpoArray, err = HPOQuery("fail")
	if err == nil {
		fmt.Println(hpoArray)
		t.Fail()
	}

	hpoArray, err = HPOQuery("disease_id='1'")
	if err != nil || hpoArray[0].DiseaseID != "1" {
		fmt.Println(err)
		fmt.Println(hpoArray[0])
		t.Fail()
	}
}
