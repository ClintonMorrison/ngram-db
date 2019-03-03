package database

import (
	"testing"
)

func TestDatabase_SetNames(t *testing.T) {
	db := New()
	db.AddSet("C", 3)
	db.AddSet("B", 3)
	db.AddSet("A", 3)
	names := db.SetNames()
	expected := []string{"A", "B", "C"}

	for i := 0; i < len(expected); i++ {
		if names[i] != expected[i] {
			t.Error("Exected SetNames() return sorted names")
		}
	}
}

func TestDatabase_GetSet(t *testing.T) {
	db := New()
	db.AddSet("test", 3)
	set, _ := db.GetSet("test")
	if set.N != 3 {
		t.Error("Expected getSet() to return the specified set")
	}

	set, err := db.GetSet("test2")
	if err == nil {
		t.Error("Expected getSet() to return an error when the requested set does not exist")
	}
}

func TestDatabase_CountsForSize(t *testing.T) {
	db := New()
	db.AddSet("test", 3)
	db.AddText("test", "ABC")

	counts, _ := db.CountsForSize("test", 1)
	if counts["A"] != 1 {
		t.Error("Expected CountsForSize() to return the counts")
	}

	_, err := db.CountsForSize("test", 6)
	if err == nil {
		t.Error("Expected CountsForSize() to return an error if n is too big")
	}

	_, err = db.CountsForSize("test12", 2)
	if err == nil {
		t.Error("Expected CountsForSize() to return an error if set does not exist")
	}
}

func TestDatabase_RemoveSet(t *testing.T) {
	db := New()
	db.AddSet("test", 3)
	db.RemoveSet("test")

	if len(db.sets) != 0 {
		t.Error("Expected RemoveSet() to remove set")
	}

	err := db.RemoveSet("test")
	if err == nil {
		t.Error("Expected RemoveSet() to return an error if set does not exist")
	}
}

func TestDatabase_AddText(t *testing.T) {
	db := New()
	db.AddSet("test", 3)
	db.AddText("test", "ABC")
	if db.sets["test"].Count("ABC") != 1 {
		t.Error("Expected add to add the ngram to the correct set")
	}

	err := db.AddText("test2", "ABCD")
	if err == nil {
		t.Error("Expected AddText() to reutn an error when the requested set does not exist")
	}
}

func TestDatabase_AddSet(t *testing.T) {
	db := New()
	db.AddSet("test", 3)

	if len(db.sets) != 1 {
		t.Error("Expected AddSet() to add a set")
	}
}
