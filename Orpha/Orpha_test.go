package orpha

import (
	"fmt"
	"testing"

	"github.com/artonge/Tamalou/Queries"
)

func TestQuery(t *testing.T) {

	results, err := Query(Queries.DBQuery{})

	if err != nil {
		fmt.Println("Error in Orpha TestQuery: \n   ==> ", err, "\n ==> ", results)
		t.Fail()
	}
}
