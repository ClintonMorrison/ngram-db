package ngram

import (
	"testing"
)

type NGramTestCase struct {
	N      int
	Text   string
	NGrams []NGram
}

func (testCase *NGramTestCase) matches(ngrams []NGram) bool {
	if len(ngrams) != len(testCase.NGrams) {
		return false
	}

	for i := 0; i < len(ngrams); i++ {
		if ngrams[i] != testCase.NGrams[i] {
			return false
		}
	}

	return true
}

func (testCase *NGramTestCase) check(t *testing.T) {
	actual := ToNGrams(testCase.N, testCase.Text)
	matches := testCase.matches(actual)
	if !matches {
		t.Errorf("Expected %v but got %v", testCase.NGrams, actual)
	}
}

func TestToNGrams(t *testing.T) {
	testCases := []NGramTestCase{
		{1, "ABC", []NGram{"A", "B", "C"}},
		{2, "ABCD", []NGram{"AB", "BC", "CD"}},
		{3, "ABCD", []NGram{"ABC", "BCD"}},
		{4, "ABCD", []NGram{"ABCD"}},
		{5, "ABCD", []NGram{}},
		{6, "ABCD", []NGram{}},
	}

	for _, testCase := range testCases {
		testCase.check(t)
	}
}

func TestCastStringsToNGrams(t *testing.T) {
	strings := []string{"A"}
	ngrams := CastStringsToNGrams(strings)

	if string(ngrams[0]) != "A" {
		t.Error("Expected CastStringToNGrams to return the same strings")
	}

}
