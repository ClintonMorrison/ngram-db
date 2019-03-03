package responses

import (
	"ngramdb/server/ngram"
	"reflect"
	"fmt"
)

type Generic struct {
	Success bool `json:"success"`
}

type Error struct {
	Success bool   `json:"success"`
	ErrorType   string `json:"error"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorType, e.Message)
}

func FromError(e error) Error {
	typeName := reflect.TypeOf(e).Name()
	return Error{false, typeName, e.Error()}
}

type Sets struct {
	Success bool     `json:"success"`
	Sets    []string `json:"sets"`
}

type NGrams struct {
	Success bool                  `json:"success"`
	NGrams  map[ngram.NGram]int64 `json:"ngrams"`
}

type Count struct {
	Success bool  `json:"success"`
	Count   int64 `json:"count"`
}

type Frequency struct {
	Success   bool    `json:"success"`
	Frequency float64 `json:"frequency"`
	Count     int64   `json:"count"`
	Total     int64   `json:"total"`
}

type Completion struct {
	NGram       ngram.NGram `json:"ngram"`
	Probability float64     `json:"probability"`
}

type Completions struct {
	Success     bool         `json:"success"`
	Completions []Completion `json:"completions"`
}

type SetProbability struct {
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
}

type ProbableSets struct {
	Success bool             `json:"success"`
	Sets    []SetProbability `json:"sets"`
}
