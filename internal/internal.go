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
	rep              Replacer
	stockToTransform []byte
	stockToWrite     []byte
}

func (tr *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	return
}

func (tr *Transformer) Reset() {
}

type Replacer interface {
	LowerFunc(uint8) []byte
	UpperFunc(uint8) []byte
	DigitFunc(uint8) []byte
}

// replace replaces a-zA-Z0-9 with regular style into specific styles.
// Invalid bytes, which could not be decoded into rune, are passed through.
func replace(rep Replacer, p []byte) []byte {
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
