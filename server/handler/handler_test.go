package handler

import (
	"testing"
	"ngramdb/server/database"
	"ngramdb/server/query"
)

func TestQueryHandler_Handle(t *testing.T) {
	db := database.New()
	db.AddSet("test", 3)

	q := query.New()
	q.Type = query.GET_SETS

	handler := New(db)
	response := handler.Handle(q, nil)
	_, isSetsResponse := response.(SetsResponse)

	if !isSetsResponse {
		t.Error("Expected Handle() to return the appropriate response")
	}
}