package orpha

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {

	results, err := Query(map[string]interface{}{})

	if err != nil {
		fmt.Println("Error in Orpha TestQuery: \n   ==> ", err, "\n ==> ", results)
		t.Fail()
	}
}
