package database

import (
	"ngramdb/server/ngram"
	"ngramdb/server/query"
	"fmt"
)


type Database struct {
	sets map[string]*ngram.Set
}

func New() *Database {
	database := Database{}
	database.sets = make(map[string]*ngram.Set)
	return &database
}

func (db *Database) getSet(key string) (*ngram.Set, error) {
	set := db.sets[key]
	if set == nil || !set.Exists() {
		return nil, NotFoundError{ key }
	}

	return set, nil
}

func (db *Database) Remove(key string) error {
	_, err := db.getSet(key)
	if err != nil {
		return err
	}

	delete(db.sets, key)
	return nil
}

func (db *Database) AddText(key string, text string) error {
	set, err := db.getSet(key)
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

func (db *Database) AnswerQuery(i interface{}) (string, error) {
	switch q := i.(type) {
	case query.AddSet:
		return "", db.AddSet(q.SetName, q.N)
	case query.AddText:
		return "", db.AddText(q.SetName, q.Text)
	case query.GetCount:
		set, err := db.getSet(q.SetName)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%d", set.Count(ngram.NGram(q.NGram))), nil
	}
	return "", nil // TODO: return error
}
