package HPO

import (
	"fmt"
	"testing"
)

func TestSQLiteQuery(t *testing.T) {
	// Fetch all
	hpoArray, err := HPOQuery(map[string]interface{}{
		"1": "1",
	})
	if err != nil {
		fmt.Println("Unit Test Error (1=1): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}

	// Fetch none - wrong query
	hpoArray, err = HPOQuery(map[string]interface{}{
		"fail": "success",
	})
	if err == nil {
		fmt.Println("Unit Test Error (fail=success): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}

	// Fetch some
	hpoArray, err = HPOQuery(map[string]interface{}{
		"and": map[string]interface{}{
			"disease_id": "1",
			"1":          1,
		},
	})
	if err != nil || hpoArray[0].DiseaseID != "1" {
		fmt.Println("Unit Test Error (disease_id=1 AND 1 = 1): \n	==> ", err, "\n	==> ", hpoArray)
		t.Fail()
	}
}

func TestOBOQuery(t *testing.T) {
	// Find one
	results, err := HPOOBOQuery(map[string]interface{}{
		"id": "HP:0000001",
	})

	if err != nil || len(results) != 1 {
		fmt.Println("Unit Test Error (id=HP:0000001): \n	==> ", err, "\n	==> ", results)
		fmt.Println(results, err)
		t.Fail()
	}

	// Find none
	results, err = HPOOBOQuery(map[string]interface{}{
		"id": "none",
	})

	if err != nil || len(results) != 0 {
		fmt.Println("Unit Test Error (id=HP:0000001): \n	==> ", err, "\n	==> ", results)
		fmt.Println(results, err)
		t.Fail()
	}

	// Fail query
	results, err = HPOOBOQuery(map[string]interface{}{
		"id": 1,
	})

	if err != nil || len(results) != 0 {
		fmt.Println("Unit Test Error (id=HP:0000001): \n	==> ", err, "\n	==> ", results)
		fmt.Println(results, err)
		t.Fail()
	}
}
