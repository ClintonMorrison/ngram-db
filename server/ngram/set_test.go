package ngram

import (
"testing"
)



func TestSet(t *testing.T) {
	set := NewSet(2)

	// Test Empty()
	isEmpty := set.Empty()
	if !isEmpty {
		t.Error("Expected Empty to return false")
	}

	// Test addNGram()
	set.addNGram("A")
	if set.counts[1]["A"] != 1 {
		t.Error("Expected addNGram to increment count")
	}

	if set.totals[1] != 1 {
		t.Error("Expected addNGram to increment total")
	}

	// Test Empty()
	isEmpty = set.Empty()
	if isEmpty {
		t.Error("Expected Empty to return false")
	}

	// Test Count()
	count := set.Count("A")
	if count != 1 {
		t.Error("Expected Count to return 1")
	}

	// Test Total()
	total := set.Total(1)
	if total != 1 {
		t.Error("Expected Total to return total")
	}

	// Test Add()
	set.Add("ABC")
	if set.counts[1]["A"] != 2 {
		t.Error("Expected Add() to increment 'A' count")
	}

	if set.counts[2]["BC"] != 1 {
		t.Error("Expected Add() to increment 'CD' count")
	}

	// Test Copy()
	set2 := set.Copy()
	if set2.counts[1]["A"] != 2 {
		t.Error("Expected Copy() to return an identical set")
	}

	// Test Union()
	setA := NewSet(2)
	setA.Add("AB")
	setB := NewSet(2)
	setB.Add("BC")
	setC := setA.Union(setB)
	if setC.Count("B") != 2 {
		t.Error("Expected Union() to combine counts from both sets")
	}
}
