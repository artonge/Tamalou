package Queries

import "fmt"

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
		fullQuery = "symptom='" + query.Value() + "'"
	}

	// For each value in the map
	// 	==> build a part of the query and append it to fullQuery
	for _, child := range query.Children() {
		fmt.Println("ok")
		switch child.Type() {
		case "and", "or":
			fullQuery += "(" + BuildSQLQuery(child) + ")"
		default:
			fmt.Println("Error while building SQL query\n	==> ", query)
		}
	}

	return fullQuery[:len(fullQuery)-len(operator)]
}
