package HPO

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
)

func TestSQLiteQuery(t *testing.T) {
	results, err := SQLiteQuery([]string{"HP:0000003", "HP:0000110", "HP:0000113", "HP:0000077", "HP:0000086", "HP:0000085", "HP:0000105", "HP:0000075", "HP:0000104", "HP:0000107"})
	if err != nil {
		fmt.Println("Unit Test Failed when testing SQLiteQuery:\n	==> ", err, "\n	==> ", results)
		t.Fail()
	}
}

func TestOBOIndexQuery(t *testing.T) {
	_, err := QueryIndex(Queries.NewTamalouQuery("symptom", "Name:Multicystic kidney dysplasia", nil))

	if err != nil {
		fmt.Println("Error while querying obo index:\n	==> ", err)
		fmt.Println(err)
	}
}

func TestHPOQuery(t *testing.T) {
	// Fetch all
	diseasesArray, err := QueryHPO(Queries.NewTamalouQuery("symptom", "Multicystic kidney dysplasia", nil))
	if err != nil {
		fmt.Println("Unit Test Error (1=1): \n	==> ", err, "\n	==> ", diseasesArray)
		t.Fail()
	}

	// Fetch some
	diseasesArray, err = QueryHPO(Queries.NewTamalouQuery("and", "", []Queries.ITamalouQuery{
		Queries.NewTamalouQuery("symptom", "Multicystic kidney dysplasia", nil),
	}))
	for _, disease := range diseasesArray {
		fmt.Println(disease.Name)
	}
	if err != nil {
		fmt.Println("Unit Test Error (disease_id=1 AND 1=1): \n	==> ", err, "\n	==> ", diseasesArray)
		t.Fail()
	}
}
