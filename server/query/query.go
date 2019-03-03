package query

import "strconv"

type Type int
const (
	INVALID Type = iota
	ADD_SET
	ADD_TEXT
	GET_NGRAMS
	GET_COUNT
	GET_FREQ
	GET_COMPLETIONS
	GET_PROBABLE_SETS
)

type Query struct {
	Type Type
	SetFields []string
	TextFields []string
	NumberFields []int
}

func newQuery() *Query {
	query := Query{}
	query.Type = INVALID
	query.SetFields = make([]string, 0)
	query.TextFields = make([]string, 0)
	query.NumberFields = make([]int, 0)

	return &query
}

func (results *Query) AddField(fieldType string, value string) {
	switch fieldType {
	case "set":
		results.SetFields = append(results.SetFields, value)
		return
	case "text":
		results.TextFields = append(results.SetFields, value)
		return
	case "number":
		int64Value, _ := strconv.ParseInt(value, 10, 0)
		intValue := int(int64Value)
		results.NumberFields = append(results.NumberFields, intValue)
		return
	}
}