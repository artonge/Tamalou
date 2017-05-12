package Queries

type DBQuery map[string]interface{}

// type TamalouQuery map[string]interface{}
type TamalouQueryType string

type ITamalouQuery interface {
	Children() []ITamalouQuery // Return The query children
	Type() TamalouQueryType    // OR/AND/SYMPTOM
	Value() string             // For symptoms only
}

type TamalouQuery struct {
	children  []ITamalouQuery
	queryType TamalouQueryType
	value     string
}

func (q TamalouQuery) Children() []ITamalouQuery {
	return q.children
}

func (q TamalouQuery) Type() TamalouQueryType {
	return q.queryType
}

func (q TamalouQuery) Value() string {
	return q.value
}
