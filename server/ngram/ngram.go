package ngram

type NGram string

func ToNGrams(n int, text string) []NGram {
	ngrams := make([]NGram, 0)
	end := len(text) - n + 1

	for i := 0; i < end; i++ {
		j := i+n
		ngram := NGram(text[i:j])
		ngrams = append(ngrams, ngram)
	}

	return ngrams
}

func CastStringsToNGrams(strings []string) []NGram {
	ngrams := make([]NGram, len(strings))
	for i, str := range strings {
		ngrams[i] = NGram(str)
	}

	return ngrams
}
