package style

import (
	"unicode/utf8"

	"golang.org/x/text/transform"
)

const (
	REGULAR_LOWER_MIN = 97
	REGULAR_LOWER_MAX = 122
	REGULAR_UPPER_MIN = 65
	REGULAR_UPPER_MAX = 90
	REGULAR_DIGIT_MIN = 48
	REGULAR_DIGIT_MAX = 57
)

// Transformer is a implement of transform.Transformer.
// This aims to replace characters which is composed of one byte,
// so multi bytes characters or invalid bytes are passed through.
type Transformer struct {
	rep          Replacer
	stockToWrite []byte
}

func NewTransformer(r Replacer) *Transformer {
	return &Transformer{r, nil}
}

func (tr *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nSrc = len(src)

	// joined the remained bytes which were not written into dst in the previous loop.
	replaced := Replace(tr.rep, src)
	if tr.stockToWrite != nil {
		replaced = append(replaced, tr.stockToWrite...)
		tr.stockToWrite = nil
	}

	if len(dst) >= len(replaced) {
		copy(dst, replaced)
		nDst = len(replaced)
		err = nil
		return
	}

	tr.stockToWrite = replaced[len(dst):]
	copy(dst, replaced[:len(dst)])
	nDst = len(dst)
	err = transform.ErrShortDst
	return
}

func (tr *Transformer) Reset() {
	tr.stockToWrite = nil
}

type Replacer interface {
	LowerFunc(uint8) []byte
	UpperFunc(uint8) []byte
	DigitFunc(uint8) []byte
}

// replace replaces a-zA-Z0-9 with regular style into specific styles.
// Invalid bytes, which could not be decoded into rune, are passed through.
func Replace(rep Replacer, p []byte) []byte {
	var replaced []byte

	for len(p) > 0 {
		r, n := utf8.DecodeRune(p)

		if r == utf8.RuneError {
			replaced = append(replaced, p[:n]...)
		} else {
			replaced = append(replaced, replaceByRune(rep, p[:n])...)
		}

		p = p[n:]
	}
	return replaced
}

// replaceByRune focuses on a valid []byte which could be decoded into a rune.
// Regular style a-zA-Z0-9 are replaced by specific functions.
func replaceByRune(rep Replacer, p []byte) []byte {
	if isRegularLower(p) {
		return rep.LowerFunc(p[0])
	}
	if isRegularUpper(p) {
		return rep.UpperFunc(p[0])
	}
	if isRegularDigit(p) {
		return rep.DigitFunc(p[0])
	}
	return p
}

func inRange(src, min, max uint8) bool {
	return src >= min && src <= max
}

func isRegularLower(src []byte) bool {
	if len(src) != 1 {
		return false
	}
	return inRange(src[0], REGULAR_LOWER_MIN, REGULAR_LOWER_MAX)
}

func isRegularUpper(src []byte) bool {
	if len(src) != 1 {
		return false
	}
	return inRange(src[0], REGULAR_UPPER_MIN, REGULAR_UPPER_MAX)
}

func isRegularDigit(src []byte) bool {
	if len(src) != 1 {
		return false
	}
	return inRange(src[0], REGULAR_DIGIT_MIN, REGULAR_DIGIT_MAX)
}
