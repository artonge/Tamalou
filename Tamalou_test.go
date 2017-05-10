package main

import (
	"fmt"
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
