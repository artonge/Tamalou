package Queries

import (
	"testing"
)

func TestParseQuery(t *testing.T) {
	query := ParseQuery("ventre AND tete OR hand")
	subQuery := query.Children()[0]

	if query.Type() != "or" ||
		len(query.Children()) != 2 ||
		query.Children()[1].Type() != "symptom" ||
		query.Children()[1].Value() != "hand" ||
		subQuery.Type() != "and" ||
		len(subQuery.Children()) != 2 ||
		subQuery.Children()[0].Type() != "symptom" ||
		subQuery.Children()[1].Type() != "symptom" ||
		subQuery.Children()[0].Value() != "ventre" ||
		subQuery.Children()[1].Value() != "tete" ||
		len(subQuery.Children()[0].Children()) != 0 ||
		len(subQuery.Children()[1].Children()) != 0 {
		t.Fail()
	}
}
