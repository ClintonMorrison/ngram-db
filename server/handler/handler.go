package handler

import (
	"ngramdb/server/query"
	"ngramdb/server/database"
	"ngramdb/server/ngram"
)

type QueryHandler struct {
	db *database.Database
}

func New(db *database.Database) *QueryHandler {
	return &QueryHandler{db}
}

func (handler *QueryHandler) Handle(q *query.Query, err error) interface{} {
	if err != nil {
			return responseFromError(err)
	}

	// TODO: DELETE SETS, LIST SETS
	switch q.Type {
	case query.ADD_SET:
		return handler.addSet(q.SetFields[0], q.NumberFields[0])
	case query.ADD_TEXT:
		return handler.addText(q.SetFields[0], q.TextFields[0])
	case query.GET_NGRAMS:
		return handler.getNGrams(q.SetFields[0], q.NumberFields[0])
	case query.GET_COUNT:
		return handler.getCount(q.SetFields[0], q.TextFields[0])
	case query.GET_FREQ:
		return handler.getFrequency(q.SetFields[0], q.TextFields[0])
	default:
		return GenericResponse{false}
	}
}

func (handler *QueryHandler) addSet(setName string, n int) interface{} {
	err := handler.db.AddSet(setName, n)
	if err != nil {
		return responseFromError(err)
	}
	return GenericResponse{true}
}

func (handler *QueryHandler) addText(setName string, text string) interface{} {
	err := handler.db.AddText(setName, text)
	if err != nil {
		return responseFromError(err)
	}
	return GenericResponse{true}
}

func (handler *QueryHandler) getNGrams(setName string, n int) interface{} {
	ngrams, err := handler.db.CountsForSize(setName, n)
	if err != nil {
		return responseFromError(err)
	}

	return NGramsResponse{true, ngrams}
}

func (handler *QueryHandler) getCount(setName string, text string) interface{} {
	ngram := ngram.NGram(text)

	set, err := handler.db.GetSet(setName)
	if err != nil {
		return responseFromError(err)
	}

	count := set.Count(ngram)

	return CountResponse{true, count}
}

func (handler *QueryHandler) getFrequency(setName string, text string) interface{} {
	ngram := ngram.NGram(text)

	set, err := handler.db.GetSet(setName)
	if err != nil {
		return responseFromError(err)
	}

	freq := set.Freq(ngram)
	count := set.Count(ngram)
	total := set.Total(len(text))

	return FreqResponse{true, freq, count, total}
}
