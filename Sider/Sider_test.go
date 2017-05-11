package sider

import (
	"fmt"
	"testing"
)

func TestQueryMeddra(t *testing.T) {

	results, err := QueryMeddra(map[string]interface{}{
		"1": 1,
	})

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}

func TestQueryMeddraAllIndications(t *testing.T) {

	results, err := QueryMeddraAllIndications(map[string]interface{}{
		"1": 1,
	})

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}

func TestQueryMeddraAllSe(t *testing.T) {

	results, err := QueryMeddraAllSe(map[string]interface{}{
		"1": 1,
	})

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}

func TestQueryMeddraFreq(t *testing.T) {

	results, err := QueryMeddraFreq(map[string]interface{}{
		"1": 1,
	})

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}
