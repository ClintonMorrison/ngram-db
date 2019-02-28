package query

type MatcherType int
const (
	LITERAL MatcherType = iota
	REGEX
)

type Matcher struct {
	Value string
	Type MatcherType
}

type Collection struct {
	Sets []string
	Matchers []Matcher
}

