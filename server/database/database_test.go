package database

import (
"testing"
)



func TestDatabase_AddSet(t *testing.T) {
	db := New()
	db.AddSet("test", 3)

	if len(db.sets) != 1 {
		t.Error("Expected AddSet() to add a set")
	}
}

func TestDatabase_getSet(t *testing.T) {
	db := New()
	db.AddSet("test", 3)
	set, _ := db.getSet("test")
	if set.N != 3 {
		t.Error("Expected getSet() to return the specified set")
	}

	set, err := db.getSet("test2")
	if err == nil {
		t.Error("Expected getSet() to return an error when the requested set does not exist")
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

func TestDatabase_Remove(t *testing.T) {
	db := New()
	db.AddSet("test", 3)
	db.AddText("test", "ABC")
	db.Remove("test")
	if db.sets["test"] != nil {
		t.Error("Expected Remove() to delete the set")
	}
}