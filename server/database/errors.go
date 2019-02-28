package database

import "fmt"

type NotFoundError struct {
	Key string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("No such key: '%s'", e.Key)
}

type DuplicateKeyError struct {
	Key string
}

func (e DuplicateKeyError) Error() string {
	return fmt.Sprintf("Key already exists: '%s'", e.Key)
}
