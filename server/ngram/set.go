package ngram

import (
	"sort"
	"fmt"
)

type Set struct {
	N int
	counts map[int]map[NGram]int64
	totals map[int]int64
}

func NewSet(n int) *Set {
	collection := Set{}
	collection.N = n

	collection.counts = make(map[int]map[NGram]int64)
	collection.totals = make(map[int]int64)
	for i := 1; i <= n; i++ {
		collection.counts[i] = make(map[NGram]int64)
	}

	return &collection
}

func (s *Set) Exists() bool {
	return s.N != 0
}


func (s *Set) Empty() bool {
	return s.N == 0 || len(s.counts[1]) == 0
}

func (s *Set) Total(length int) int64 {
	return s.totals[length]
}

func (s *Set) Count(ngram NGram) int64 {
	length := len(ngram)
	return s.counts[length][ngram]
}

func (s *Set) CountsForSize(n int) map[NGram]int64 {
	return s.counts[n]
}

func (s *Set) Freq(ngram NGram) float64 {
	length := len(ngram)
	return float64(s.Count(ngram)) / float64(s.Total(length))
}

func (s *Set) addNGram(ngram NGram) {
	length := len(ngram)
	s.counts[length][ngram] += 1
	s.totals[length] += 1
}

func (s *Set) Add(text string) {
	for n := 1; n <= s.N; n++ {
		ngrams := ToNGrams(n, text)
		for _, ngram := range ngrams {
			s.addNGram(ngram)
		}
	}
}

func (s *Set) NGrams(n int) []NGram {
	countsForSize := s.counts[n]
	strings := make([]string, len(countsForSize))

	i := 0
	for ngram := range countsForSize {
		strings[i] = string(ngram)
		i++
	}

	sort.Strings(strings)

	return CastStringsToNGrams(strings)
}


func (s *Set) Print() {
	for n := 1; n <= s.N; n++ {
		ngrams := s.NGrams(n)
		fmt.Printf("> n=%d\n", n)
		for _, ngram := range ngrams {
			count := s.Count(ngram)
			fmt.Printf("\t%s: %d\n", string(ngram), count)
		}
	}
}

func (s *Set) Copy() *Set {
	set := NewSet(s.N)
	for i := 1; i <= s.N; i++ {
		set.totals[i] = s.totals[i]
		for ngram, count := range s.counts[i] {
			set.counts[i][ngram] = count
		}
	}

	return set
}
