package sider

import (
	"fmt"
	"testing"
)

func TestQueryMeddra(t *testing.T) {

	results, err := QueryMeddra(map[string]interface{}{
		"1": 1,
	})

	fmt.Println(results[0], err)

	if err != nil {
		t.Fail()
	}
}
