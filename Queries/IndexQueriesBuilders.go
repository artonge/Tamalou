package Queries

import "fmt"

func BuildIndexQuery(query ITamalouQuery) string {
	var (
		fullQuery string
	)

	if query.Type() == "symptom" {
		fullQuery = query.Value()
	}

	// For each value in the map
	// 	==> build a part of the query and append it to fullQuery
	for _, child := range query.Children() {
		fmt.Println("ok")
		switch child.Type() {
		case "and", "or":
			fullQuery += " " + BuildIndexQuery(child)
		default:
			fmt.Println("Error while building Index query\n	==> ", query)
		}
	}

	return fullQuery
}
