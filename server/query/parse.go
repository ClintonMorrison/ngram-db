package query

import (
	"regexp"
	"strings"
)

var placeholderPatterns = map[string]string{
"<set>": "(?P<set>[a-zA-Z_\\-]+)",
"<text>":   "'(?P<text>.+)'",
"<number>": "(?P<number>[1-9][0-9]*)",
}

var queryPatterns = map[Type]string{
	GET_SETS: "GET SETS",
	DELETE_SET: "DELETE SET <set>",
	ADD_SET: "ADD SET <set>\\(<number>\\)",
	ADD_TEXT: "ADD TEXT <text> IN <set>",
	GET_NGRAMS: "GET NGRAMS\\(<number>\\) IN <set>",
	GET_COUNT: "GET COUNT OF <text> IN <set>",
	GET_FREQ: "GET FREQ OF <text> IN <set>",
	GET_COMPLETIONS: "GET COMPLETIONS OF <text> IN <set>",
	GET_PROBABLE_SETS: "GET PROBABLE SETS OF <text> IN <set>",
}

func queryPatternToRegex(pattern string) *regexp.Regexp {
	for placeholder, replacement := range placeholderPatterns {
		pattern = strings.Replace(pattern, placeholder, replacement, -1)
	}

	return regexp.MustCompile(pattern)
}

func getQueryRegex() map[Type]*regexp.Regexp {
	queryRegex := map[Type]*regexp.Regexp{}

	for queryType, pattern := range queryPatterns {
		queryRegex[queryType] = queryPatternToRegex(pattern)
	}

	return queryRegex
}

var queryRegex = getQueryRegex()

func Parse(rawQuery string) (*Query, error) {
	parsedQuery := New()
	for queryType, regex := range queryRegex {
		matches := regex.FindStringSubmatch(rawQuery)

		if len(matches) == 0 {
			continue
		}

		parsedQuery.Type = queryType
		matchNames := regex.SubexpNames()

		for i := range matches {
			parsedQuery.AddField(matchNames[i], matches[i])
		}
	}

	if parsedQuery.Type == INVALID {
		return nil, ParseError{rawQuery}
	}


	return parsedQuery, nil
}