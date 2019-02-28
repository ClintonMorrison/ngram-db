package query

import (
	"regexp"
	"fmt"
	"strconv"
)

/*
Example queries:
ADD SET <SET_NAME>(5)`
ADD TEXT <text> TO <COLLECTIONS>
GET <ATTRIBUTE> IN <COLLECTIONS>
 */

const setNamePattern = "[a-zA-Z_\\-]+"
const propertiesPattern = "\\(([1-9][0-9]*)\\)"
const textPattern = "\"(.+)\""
const numberPattern = "[1-9][0-9]*"

var addSetQueryRegex = regexp.MustCompile(
	fmt.Sprintf("^ADD SET (%s)%s", setNamePattern, propertiesPattern))

var addTextQueryRegex = regexp.MustCompile(
	fmt.Sprintf("^ADD TEXT\\(%s\\) TO (%s)", textPattern, setNamePattern))

var getCountQueryRegex = regexp.MustCompile(
	fmt.Sprintf("^GET COUNT OF %s IN (%s)", textPattern, setNamePattern))

var getFreqQueryRegex = regexp.MustCompile(
	fmt.Sprintf("^GET FREQ OF %s IN (%s)", textPattern, setNamePattern))

var getNGramsQueryRegex = regexp.MustCompile(
	fmt.Sprintf("^GET NGRAMS\\((%s)\\) IN (%s)", numberPattern, setNamePattern))


func Parse(query string) interface{} {
	match := addSetQueryRegex.FindStringSubmatch(query)
	if match != nil {
		return parseAddSet(match)
	}

	match = addTextQueryRegex.FindStringSubmatch(query)
	if match != nil {
		return parseAddText(match)
	}

	match = getCountQueryRegex.FindStringSubmatch(query)
	if match != nil {
		return parseCountQuery(match)
	}

	match = getFreqQueryRegex.FindStringSubmatch(query)
	if match != nil {
		return parseFreqQuery(match)
	}

	match = getNGramsQueryRegex.FindStringSubmatch(query)
	if match != nil {
		return parseNGramsQuery(match)
	}

	return nil
}

func parseAddSet(match []string) AddSet {
	setName := match[1]
	n, _ := strconv.ParseInt(match[2], 10, 0)

	return AddSet{ setName, int(n), true}
}

func parseAddText(match []string) AddText {
	text := match[1]
	setName := match[2]

	return AddText{ setName,  text }
}

func parseCountQuery(match []string) GetCount {
	text := match[1]
	setName := match[2]

	return GetCount{ setName,  text }
}

func parseFreqQuery(match []string) GetFreq {
	text := match[1]
	setName := match[2]

	return GetFreq{ setName,  text }
}

func parseNGramsQuery(match []string) GetNGrams {
	n, _ := strconv.ParseInt(match[1], 10, 0)
	setName := match[2]

	return GetNGrams{ setName,  int(n) }
}