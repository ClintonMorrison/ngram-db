package handler

import (
	"ngramdb/server/ngram"
	"reflect"
)

type GenericResponse struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	Message string `json:"message"`
}

func responseFromError(e error) ErrorResponse {
	typeName := reflect.TypeOf(e).Name()
	return ErrorResponse{false, typeName, e.Error()}
}

type NGramsResponse struct {
	Success bool `json:"success"`
	NGrams map[ngram.NGram]int64 `json:"ngrams"`
}

type CountResponse struct {
	Success bool `json:"success"`
	Count int64 `json:"count"`
}

type FreqResponse struct {
	Success bool `json:"success"`
	Frequency float64 `json:"frequency"`
	Count int64 `json:"count"`
	Total int64 `json:"total"`
}

type Completion struct {
	NGram ngram.NGram `json:"ngram"`
	Probability float64 `json:"probability"`
}

type CompletionResponse struct {
	Success bool `json:"success"`
	Completions []Completion `json:"completions"`
}

type SetProbability struct {
	Name string `json:"name"`
	Probability float64 `json:"probability"`
}

type ProbableSetsResponse struct {
	Success bool `json:"success"`
	Sets []SetProbability `json:"sets"`
}
