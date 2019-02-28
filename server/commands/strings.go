package commands

import (
	"unicode"
)

func Tokenize(s string) []string {
	result := make([]string, 0)

	token := ""
	for _, c := range s {
		if unicode.IsSpace(rune(c)) {
			if len(token) > 0 {
				result = append(result, token)
			}
			token = ""
		} else {
			token = token + string(c)
		}
	}

	if len(token) > 0 {
		result = append(result, token)
	}

	return result
}