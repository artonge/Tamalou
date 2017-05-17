package Queries

import "strconv"

func BuildSQLQuery(query ITamalouQuery) string {
	var (
		fullQuery string
		operator  string
	)

	switch query.Type() {
	case "or":
		operator = " OR "
	case "and":
		operator = " AND "
	default:
		operator = ""
		switch query.Value().(type) {
		case string:
			return string(query.Type()) + "='" + query.Value().(string) + "'"
		case int:
			return string(query.Type()) + "=" + strconv.Itoa(query.Value().(int))
		}
	}

	// For each sub queries
	// 	==> build a part of the query and append it to fullQuery
	for _, child := range query.Children() {
		switch child.Type() {
		case "and", "or":
			fullQuery += "(" + BuildSQLQuery(child) + ")" + operator
		default:
			fullQuery += BuildSQLQuery(child) + operator
		}
	}

	return fullQuery[:len(fullQuery)-len(operator)]
}
