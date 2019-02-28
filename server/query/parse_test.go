package query

import (
	"testing"
)

func TestParse_addSet(t *testing.T) {
	rawQuery := "ADD SET test(10)"
	query := Parse(rawQuery).(AddSet)

	if query.N != 10 {
		t.Error("Expected query to have N=10")
	}

	if query.SetName != "test" {
		t.Error("Expected query to setName 'test'")
	}
}

func TestParse_addText(t *testing.T) {
	rawQuery := "ADD TEXT(\"test text\") TO test"
	query := Parse(rawQuery).(AddText)

	if query.Text != "test text" {
		t.Error("Expected query to have text 'test text'")
	}

	if query.SetName != "test" {
		t.Error("Expected query to setName 'test'")
	}
}

func TestParse_getCount(t *testing.T) {
	rawQuery := "GET COUNT OF \"AB\" IN test"
	query := Parse(rawQuery).(GetCount)

	if query.NGram != "AB" {
		t.Error("Expected query to have ngram 'AB'")
	}

	if query.SetName != "test" {
		t.Error("Expected query to setName 'test'")
	}
}

func TestParse_getFreq(t *testing.T) {
	rawQuery := "GET FREQ OF \"AB\" IN test"
	query := Parse(rawQuery).(GetFreq)

	if query.NGram != "AB" {
		t.Error("Expected query to have ngram 'AB'")
	}

	if query.SetName != "test" {
		t.Error("Expected query to setName 'test'")
	}
}

func TestParse_getNGrams(t *testing.T) {
	rawQuery := "GET NGRAMS(2) IN test"
	query := Parse(rawQuery).(GetNGrams)

	if query.N != 2 {
		t.Error("Expected query to have n '2'")
	}

	if query.SetName != "test" {
		t.Error("Expected query to setName 'test'")
	}
}