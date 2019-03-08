package query

import (
	"testing"
)

func expectQuery(t *testing.T, query *Query, queryType Type, set string, text string, number int) {
	if query.Type != queryType {
		t.Errorf("Expected query type %d but got %d", queryType, query.Type)
	}

	if set != "" && set != query.SetFields[0] {
		t.Errorf("Expected set %s but got %s", set, query.SetFields[0])
	}

	if text != "" && text != query.TextFields[0] {
		t.Errorf("Expected text %s but got %s", text, query.TextFields[0])
	}

	if number != 0 && number != query.NumberFields[0] {
		t.Errorf("Expected number %d but got %d", number, query.NumberFields[0])
	}
}

func TestParse_getSets(t *testing.T) {
	query, _ := Parse("GET SETS")
	expectQuery(
		t,
		query,
		GET_SETS,
		"",
		"",
		0)
}

func TestParse_deleteSet(t *testing.T) {
	query, _ := Parse("DELETE SET test")
	expectQuery(
		t,
		query,
		DELETE_SET,
		"test",
		"",
		0)
}

func TestParse_addSet(t *testing.T) {
	query, _ := Parse("ADD SET test(4)")
	expectQuery(
		t,
		query,
		ADD_SET,
		"test",
		"",
		4)
}

func TestParse_addText(t *testing.T) {
	query, _ := Parse("ADD TEXT 'ab cd' IN test")
	expectQuery(
		t,
		query,
		ADD_TEXT,
		"test",
		"ab cd",
		0)
}

func TestParse_getNGrams(t *testing.T) {
	query, _ := Parse("GET NGRAMS(3) IN test")
	expectQuery(
		t,
		query,
		GET_NGRAMS,
		"test",
		"",
		3)
}

func TestParse_getCount(t *testing.T) {
	query, _ := Parse("GET COUNT OF 'ab' IN test")
	expectQuery(
		t,
		query,
		GET_COUNT,
		"test",
		"ab",
		0)
}

func TestParse_getFreq(t *testing.T) {
	query, _ := Parse("GET FREQ OF 'ab' IN test")
	expectQuery(
		t,
		query,
		GET_FREQ,
		"test",
		"ab",
		0)
}

func TestParse_getCompletions(t *testing.T) {
	query, _ := Parse("GET COMPLETION OF 'ab' IN test")
	expectQuery(
		t,
		query,
		GET_COMPLETIONS,
		"test",
		"ab",
		0)
}

func TestParse_getProbableSet(t *testing.T) {
	query, _ := Parse("GET PROBABLE SET OF 'ab'")
	expectQuery(
		t,
		query,
		GET_PROBABLE_SET,
		"",
		"ab",
		0)
}

func TestParse_invalidQuery(t *testing.T) {
	query, err := Parse("GET INVALID OF 'ab' IN test")
	if query != nil {
		t.Errorf("Expected query to be nil")
	}

	if err == nil {
		t.Errorf("Expected parse error")
	}
}
