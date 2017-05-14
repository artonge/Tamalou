package sider

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
)

func TestQueryMeddra(t *testing.T) {

	results, err := QueryMeddra(Queries.DBQuery{
		"1": 1,
	})

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}

func TestQueryMeddraAllIndications(t *testing.T) {

	results, err := QueryMeddraAllIndications(Queries.DBQuery{
		"1": 1,
	})

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}

func TestQueryMeddraAllSe(t *testing.T) {

	results, err := QueryMeddraAllSe(Queries.DBQuery{
		"1": 1,
	})

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}

func TestQueryMeddraFreq(t *testing.T) {

	results, err := QueryMeddraFreq(Queries.DBQuery{
		"1": 1,
	})

	if err != nil {
		fmt.Println("Error in Sider Test : \n	==>", err, "\n	==>", results)
		t.Fail()
	}
}
