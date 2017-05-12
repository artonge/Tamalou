package Queries

import "strconv"

func BuildSQLQuery(query DBQuery, queryType interface{}) string {
	var (
		fullQuery string
		operand   string
	)

	switch queryType {
	case "or":
		operand = " OR "
	case "and":
		operand = " AND "
	default:
		operand = ""
	}

	// For each value in the map
	// 	==> build a part of the query and append it to fullQuery
	for key, value := range query {
		switch key {
		case "and", "or":
			fullQuery += "(" + BuildSQLQuery(value.(DBQuery), key) + ")"
		default:
			switch value.(type) {
			case string:
				fullQuery += key + "='" + value.(string) + "'"
			case int:
				fullQuery += key + "=" + strconv.Itoa(value.(int))
			case float64:
				fullQuery += key + "=" + strconv.FormatFloat(value.(float64), 'f', 6, 64)
			default:
				fullQuery += key + "=" + strconv.Itoa(value.(int))
			}
		}
		fullQuery += operand
	}

	return fullQuery[:len(fullQuery)-len(operand)]
}
