package internal

import (
	"unicode/utf8"
)

const (
	REGULAR_LOWER_MIN = 97
	REGULAR_LOWER_MAX = 122
	REGULAR_UPPER_MIN = 65
	REGULAR_UPPER_MAX = 90
	REGULAR_DIGIT_MIN = 48
	REGULAR_DIGIT_MAX = 57
)

type Transformer struct {
	rep              *replacer
	stockToTransform []byte
	stockToWrite     []byte
}

func (tr *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	return
}

func (tr *Transformer) Reset() {
}

type replacer struct {
	lowerFunc func(uint8) []byte
	upperFunc func(uint8) []byte
	digitFunc func(uint8) []byte
}

// rep.replace replaces a-zA-Z0-9 with regular style into specific styles.
// Invalid bytes, which could not be decoded into rune, are passed through.
func (rep *replacer) replace(p []byte) []byte {
	var replaced []byte

	for len(p) > 0 {
		r, n := utf8.DecodeRune(p)

		if r == utf8.RuneError {
			replaced = append(replaced, p[:n]...)
		} else {
			replaced = append(replaced, rep.doReplace(p[:n])...)
		}

		p = p[n:]
	}
	return replaced
}

// rep.doReplace focuses on a valid []byte which could be decoded into rune.
// Regular style a-zA-Z0-9 are replaced by specific functions.
func (rep *replacer) doReplace(src []byte) []byte {
	if isRegularLower(src) {
		return rep.lowerFunc(src[0])
	}
	if isRegularUpper(src) {
		return rep.upperFunc(src[0])
	}
	if isRegularDigit(src) {
		return rep.digitFunc(src[0])
	}
	return src
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
