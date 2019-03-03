package database

import (
	"ngramdb/server/ngram"
	"sort"
)

type Database struct {
	sets map[string]*ngram.Set
}

func New() *Database {
	database := Database{}
	database.sets = make(map[string]*ngram.Set)
	return &database
}

func (db *Database) SetNames() []string {
	setNames := make([]string, len(db.sets))

	i := 0
	for setName := range db.sets {
		setNames[i] = setName
		i++
	}

	sort.Strings(setNames)
	return setNames
}

func (db *Database) GetSet(key string) (*ngram.Set, error) {
	set := db.sets[key]
	if set == nil || !set.Exists() {
		return nil, NotFoundError{ key }
	}

	return set, nil
}

func (db *Database) CountsForSize(key string, n int) (map[ngram.NGram]int64, error) {
	set, err := db.GetSet(key)
	if err != nil {
		return nil, err
	}

	if n > set.N {
		return nil, OutOfBoundsError{ key, n}
	}

	return set.CountsForSize(n), nil
}

func (db *Database) RemoveSet(key string) error {
	_, err := db.GetSet(key)
	if err != nil {
		return err
	}

	delete(db.sets, key)
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
	set := db.sets[key]
	if set != nil && db.sets[key].Exists() {
		return DuplicateKeyError{ key }
	}

	db.sets[key] = ngram.NewSet(n)
	return nil
}

