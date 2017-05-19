package Queries

import (
	"fmt"
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

func TestBuildSQLQuery(t *testing.T) {

	SQLQuery := BuildSiderQuery(" stitch_compound_id1 IN (SELECT stitch_compound_id1 FROM meddra_all_se WHERE side_effect_name =", ParseQuery("Abdominal pain OR Gastrointestinal pain AND anorexia"))
	//fmt.Println(fullQuery)
	//SQLQuery := BuildSQLQuery("SELECT * FROM meddra_all_se, meddra_freq WHERE meddra_all_se.stitch_compound_id1 = meddra_freq.stitch_compound_id1 AND meddra_all_se.side_effect_name=", ParseQuery("Anaemie OR Abdomen acute"))
	if SQLQuery != "symptom='Anaemie'" {

		fmt.Println("query: ", SQLQuery)
		t.Fail()
	}
}

func TestBuildIndexQuery(t *testing.T) {
	indexQuery := BuildIndexQuery(ParseQuery("Anaemie"))
	fmt.Println(indexQuery)
}
