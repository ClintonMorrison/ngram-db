package query

type AddSet struct {
	SetName string
	N int
	CaseSensitive bool
}

type AddText struct {
	SetName string
	Text string
}

type GetCount struct {
	SetName string
	NGram string
}

type GetFreq struct {
	SetName string
	NGram string
}

type GetNGrams struct {
	SetName string
	N int
}