package Queries

import "strings"

func GetClinicalSigns(rawQuery string) []string {
	var replacer = strings.NewReplacer("OR", ",", "AND", ",")
	formattedQuery := replacer.Replace(rawQuery)
	splitQuery := strings.Split(formattedQuery, ",")

	for _, value := range splitQuery {
		value = strings.TrimLeft(value, " ")
		value = strings.TrimRight(value, " ")
	}

	return splitQuery
}

func ParseQuery(rawQuery string) ITamalouQuery {

	splitedQuery := strings.SplitN(rawQuery, "OR", 2)

	var query ITamalouQuery

	if len(splitedQuery) == 2 {
		query = TamalouQuery{
			queryType: "or",
			children: []ITamalouQuery{
				ParseQuery(splitedQuery[0]),
				ParseQuery(splitedQuery[1]),
			},
		}
	} else {
		splitedQuery := strings.SplitN(rawQuery, "AND", 2)
		if len(splitedQuery) == 2 {
			query = TamalouQuery{
				queryType: "and",
				children: []ITamalouQuery{
					ParseQuery(splitedQuery[0]),
					ParseQuery(splitedQuery[1]),
				},
			}
		} else {
			query = TamalouQuery{
				queryType: "symptom",
				value:     strings.TrimSpace(rawQuery),
			}
		}
	}

	return query
}
