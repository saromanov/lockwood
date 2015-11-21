package text

import (
	"errors"
)

type SkipGram struct {
	Text string
	N    uint
	Skip uint
}

var (
	ErrSkipZero = errors.New("SkipGram: Skip can't be less equal to zero")
)

func (sgram *SkipGram) PreRun(data string) {
	return data
}

func (sgram *SkipGram) Run() (map[string]int, error) {
	if sgram.Skip == 0 {
		return ErrSkipZero
	}
	if sgram.N == 0 {
		sgram.N = defaultN
	}

	if sgram.Text.length < sgram.N {
		return nil, ErrTextLessN
	}

	result := map[string]int{}
	begin := 0
	for {
		text := ngram.Text[begin : begin+sgram.N]
		item, ok := result[text]
		if !ok {
			result[text] = 1
		} else {
			result[text]++
		}

		// Skip phase
		begin += sgram.Skip
		if begin >= len(sgram.Text) {
			break
		}

	}

	return result
}
