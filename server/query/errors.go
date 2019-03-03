package query

import "fmt"

type ParseError struct {
	query string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Could not parse query: %s", e.query)
}
