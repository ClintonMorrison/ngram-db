package ngram

import (
	"fmt"
	"sort"
	"math"
)

type Set struct {
	N      int `json:"n"`
	Counts map[int]map[NGram]int64 `json:"counts"`
	Totals map[int]int64 `json:"totals"`
}

func NewSet(n int) *Set {
	collection := Set{}
	collection.N = n

	collection.Counts = make(map[int]map[NGram]int64)
	collection.Totals = make(map[int]int64)
	for i := 1; i <= n; i++ {
		collection.Counts[i] = make(map[NGram]int64)
	}

	return &collection
}

func (s *Set) Exists() bool {
	return s.N != 0
}

func (s *Set) Empty() bool {
	return s.N == 0 || len(s.Counts[1]) == 0
}

func (s *Set) Total(length int) int64 {
	return s.Totals[length]
}

func (s *Set) Count(ngram NGram) int64 {
	length := len(ngram)
	return s.Counts[length][ngram]
}

func (s *Set) CountsForSize(n int) map[NGram]int64 {
	return s.Counts[n]
}

func (s *Set) Freq(ngram NGram) float64 {
	length := len(ngram)
	return float64(s.Count(ngram)) / float64(s.Total(length))
}

func (s *Set) addNGram(ngram NGram) {
	length := len(ngram)
	s.Counts[length][ngram] += 1
	s.Totals[length] += 1
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
	countsForSize := s.Counts[n]
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
		set.Totals[i] = s.Totals[i]
		for ngram, count := range s.Counts[i] {
			set.Counts[i][ngram] = count
		}
	}

	return set
}

func (s1 *Set) DistanceTo(s2 *Set) float64 {
	n := int(math.Min(float64(s1.N), float64(s2.N)))

	ngrams1 := s1.NGrams(n)
	ngrams2 := s2.NGrams(n)
	ngrams := make(map[NGram]bool)

	for _, ngram := range ngrams1 {
		ngrams[ngram] = true
	}

	for _, ngram := range ngrams2 {
		ngrams[ngram] = true
	}

	distanceSquared := float64(0)

	for ngram := range ngrams {
		a := float64(s1.Freq(ngram))
		b := float64(s2.Freq(ngram))

		distanceSquared += math.Pow(a - b, 2)
	}

	return math.Sqrt(distanceSquared)
}
