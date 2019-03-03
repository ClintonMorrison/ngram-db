package ngram

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	n := 4
	set := NewSet(n)

	if set.N != n {
		t.Error("Expected NewSet to set correct n")
	}

	if set.Counts[4] == nil {
		t.Error("Expected NewSet to initialize Counts maps")
	}

	if set.Totals == nil {
		t.Error("Expected NewSet to initialize Totals map")
	}
}

func TestSet_Exists(t *testing.T) {
	set := &Set{}

	if set.Exists() {
		t.Error("Expected Exists() to return false for uninitialized set")
	}

	set = NewSet(3)
	if !set.Exists() {
		t.Error("Expected Exists() to return true for an initialized set")
	}
}

func TestSet_Empty(t *testing.T) {
	set := NewSet(3)

	if !set.Empty() {
		t.Error("Expected Empty() to return true for empty set")
	}

	set.Add("ABC")
	if set.Empty() {
		t.Error("Expected Empty() to return false for non-empty set")
	}
}

func TestSet_Total(t *testing.T) {
	set := NewSet(3)
	set.Add("AAAB")
	total := set.Total(1)
	if total != 4 {
		t.Error("Expected Total() to return 4")
	}
}

func TestSet_Count(t *testing.T) {
	set := NewSet(3)
	set.Add("AAAB")
	total := set.Count("A")
	if total != 3 {
		t.Error("Expected Count() to return 3")
	}
}

func TestSet_CountsForForSize(t *testing.T) {
	set := NewSet(3)
	set.Add("AAAB")
	counts := set.CountsForSize(1)
	if counts["A"] != 3 {
		t.Error("Expected CountsForForSize() to return 3")
	}
}

func TestSet_Freq(t *testing.T) {
	set := NewSet(3)
	set.Add("AAAB")
	freq := set.Freq("A")
	if freq != 0.75 {
		t.Error("Expected CountsForForSize() to return .75")
	}
}

func TestSet_addNGram(t *testing.T) {
	set := NewSet(3)
	set.addNGram("A")
	set.addNGram("A")
	count := set.Count("A")

	if count != 2 {
		t.Error("Expected addNGram to add the ngram")
	}
}

func TestSet_NGrams(t *testing.T) {
	set := NewSet(3)
	set.Add("CABC")
	ngrams := set.NGrams(2)
	expected := []NGram{"AB", "BC", "CA"}

	for i := 0; i < len(expected); i++ {
		if ngrams[i] != expected[i] {
			t.Error("Expected NGrams() to return ngrams in alphabetical order")
		}
	}
}
