package handler

import (
	"ngramdb/responses"
	"ngramdb/server/database"
	"ngramdb/server/ngram"
	"ngramdb/server/query"
)

type QueryHandler struct {
	db *database.Database
}

func New(db *database.Database) *QueryHandler {
	return &QueryHandler{db}
}

func (handler *QueryHandler) Handle(q *query.Query, err error) interface{} {
	if err != nil {
		return responses.FromError(err)
	}

	switch q.Type {
	case query.GET_SETS:
		return handler.getSets()
	case query.DELETE_SET:
		return handler.deleteSet(q.SetFields[0])
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
	case query.GET_PROBABLE_SET:
		return handler.getProbableSet(q.TextFields[0])
	default:
		return responses.Generic{false}
	}
}

func (handler *QueryHandler) getSets() interface{} {
	sets := handler.db.SetNames()
	return responses.Sets{true, sets}
}

func (handler *QueryHandler) deleteSet(setName string) interface{} {
	err := handler.db.RemoveSet(setName)
	if err != nil {
		return responses.FromError(err)
	}

	return responses.Generic{true}
}

func (handler *QueryHandler) addSet(setName string, n int) interface{} {
	err := handler.db.AddSet(setName, n)
	if err != nil {
		return responses.FromError(err)
	}
	return responses.Generic{true}
}

func (handler *QueryHandler) addText(setName string, text string) interface{} {
	err := handler.db.AddText(setName, text)
	if err != nil {
		return responses.FromError(err)
	}
	return responses.Generic{true}
}

func (handler *QueryHandler) getNGrams(setName string, n int) interface{} {
	ngrams, err := handler.db.CountsForSize(setName, n)
	if err != nil {
		return responses.FromError(err)
	}

	return responses.NGrams{true, ngrams}
}

func (handler *QueryHandler) getCount(setName string, text string) interface{} {
	ngram := ngram.NGram(text)

	set, err := handler.db.GetSet(setName)
	if err != nil {
		return responses.FromError(err)
	}

	count := set.Count(ngram)

	return responses.Count{true, count}
}

func (handler *QueryHandler) getFrequency(setName string, text string) interface{} {
	ngram := ngram.NGram(text)

	set, err := handler.db.GetSet(setName)
	if err != nil {
		return responses.FromError(err)
	}

	freq := set.Freq(ngram)
	count := set.Count(ngram)
	total := set.Total(len(text))

	return responses.Frequency{true, freq, count, total}
}

func (handler *QueryHandler) getProbableSet(text string) interface{} {
	setName, prob := handler.db.ClosestSet(text)
	return responses.ProbableSet{true, setName, prob}
}
