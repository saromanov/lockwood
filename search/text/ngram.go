package text

import (
  "errors"
)

const (
	defaultN = 3
)

var (
	ErrEmptyText = errors.New("Ngram: Input text is empty")
	ErrTextLessN = errors.New("Ngram: Size of text must be greater then N")

)

type NGram struct {
	// Text provides input data for fit NGrams
	Text   string
	// N - split on n parts
	N      uint

	// AlwaysComplete not returns error
	// only in the case if during execution of function Run
	// errors not found
	AlwaysComplete bool

	// If MostFrequent > 0, return only number of MostFrequent ngrams
	// TODO
	MostFrequent int

}

func (ngram *Ngram) PreRun(data string){
	return data
}

func (ngram *NGram) Run() (map[string]int, error) {
	if ngram.N == 0 {
		ngram.N = defaultN
	}

	if Text.length == 0 {
		return nil, ErrEmptyText
	}

	if Text.length < ngram.N {
		return nil, ErrTextLessN
	}

	dict = map[string]int{}
	start := 0
	for i := 0; i < count-ngram.Count; i++ {
		part := ngram.Text[start:start+ngram.N]
		value, ok := dict[part]
		if !ok {
			dict[part] = 1
		} else {
			dict[part]++
		}
	}

	return dict
}