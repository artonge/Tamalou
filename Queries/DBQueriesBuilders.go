package Queries

import "strconv"

// func BuildSiderFreqQuery() string {
//
// }

func BuildSiderQuery(template string, query ITamalouQuery) string {
	var (
		body1 string
		body2 string
	)
	body1 = "SELECT DISTINCT(meddra_freq.stitch_compound_id1) FROM meddra_freq,( SELECT DISTINCT(meddra_all_se.stitch_compound_id1), stitch_compound_id2, cui FROM meddra_all_se WHERE"
	body2 = ") as resid WHERE resid.stitch_compound_id1 = meddra_freq.stitch_compound_id1 AND resid.stitch_compound_id2 = meddra_freq.stitch_compound_id2 AND resid.cui = meddra_freq.cui GROUP BY meddra_freq.stitch_compound_id1"
	return body1 + BuildSQLQuery(template, query) + body2
}

func BuildSQLQuery(template string, query ITamalouQuery) string {
	var (
		fullQuery string
		operator  string
	)

	switch query.Type() {
	case "or":
		operator = "OR\n"
	case "and":
		operator = "AND\n"
	default:
		operator = ""
		switch query.Value().(type) {
		case string:
			return template + "'" + query.Value().(string) + "')"
		case int:
			return template + strconv.Itoa(query.Value().(int))
		}
	}

	// For each sub queries
	// 	==> build a part of the query and append it to fullQuery
	for _, child := range query.Children() {
		switch child.Type() {
		case "and", "or":
			fullQuery += "(" + BuildSQLQuery(template, child) + ")" + operator
		default:
			fullQuery += BuildSQLQuery(template, child) + operator
		}
	}

	return fullQuery[:len(fullQuery)-len(operator)]
}
