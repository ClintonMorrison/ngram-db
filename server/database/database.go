package database

import (
	"ngramdb/server/ngram"
	"sort"
	"encoding/json"
	"io/ioutil"
)

type Database struct {
	Sets map[string]*ngram.Set `json:"Sets"`
}

func New() *Database {
	database := Database{}
	database.Sets = make(map[string]*ngram.Set)
	return &database
}

func (db *Database) SetNames() []string {
	setNames := make([]string, len(db.Sets))

	i := 0
	for setName := range db.Sets {
		setNames[i] = setName
		i++
	}

	sort.Strings(setNames)
	return setNames
}

func (db *Database) GetSet(key string) (*ngram.Set, error) {
	set := db.Sets[key]
	if set == nil || !set.Exists() {
		return nil, NotFoundError{key}
	}

	return set, nil
}

func (db *Database) CountsForSize(key string, n int) (map[ngram.NGram]int64, error) {
	set, err := db.GetSet(key)
	if err != nil {
		return nil, err
	}

	if n > set.N {
		return nil, OutOfBoundsError{key, n}
	}

	return set.CountsForSize(n), nil
}

func (db *Database) RemoveSet(key string) error {
	_, err := db.GetSet(key)
	if err != nil {
		return err
	}

	delete(db.Sets, key)
	return nil
}

func (db *Database) AddText(key string, text string) error {
	set, err := db.GetSet(key)
	if err != nil {
		return err
	}

	set.Add(text)

	return nil
}

func (db *Database) AddSet(key string, n int) error {
	set := db.Sets[key]
	if set != nil && db.Sets[key].Exists() {
		return DuplicateKeyError{key}
	}

	db.Sets[key] = ngram.NewSet(n)
	return nil
}

func (db *Database) ToFile(filename string) error {
	serialized, err := json.Marshal(db)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, serialized, 0666)
}

func FromFile(filename string) (*Database, error) {
	db := New()

	serialized, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(serialized, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}