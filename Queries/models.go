package Queries

type DBQuery map[string]interface{}

// type TamalouQuery map[string]interface{}
type TamalouQueryType string

type ITamalouQuery interface {
	Children() []ITamalouQuery // Return The query children
	Type() TamalouQueryType    // OR/AND/SYMPTOM
	Value() interface{}        // For symptoms only
}

type TamalouQuery struct {
	children  []ITamalouQuery
	queryType TamalouQueryType
	value     interface{}
}

func (q TamalouQuery) Children() []ITamalouQuery {
	return q.children
}

func (q TamalouQuery) Type() TamalouQueryType {
	return q.queryType
}

func (q TamalouQuery) Value() interface{} {
	return q.value
}

func NewTamalouQuery(queryType string, value interface{}, children []ITamalouQuery) TamalouQuery {
	return TamalouQuery{
		value:     value,
		queryType: TamalouQueryType(queryType),
		children:  children,
	}
}
