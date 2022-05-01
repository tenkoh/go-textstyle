package textstyle

import (
	"unicode/utf8"

	"golang.org/x/text/transform"
)

var altMap = map[rune]rune{}

// Transformer is a implement of transform.Transformer.
// This aims to replace characters which is composed of one byte,
// so multi bytes characters or invalid bytes are passed through.
type Transformer struct {
	Rep          Replacer
	stockToWrite []byte
}

func NewTransformer(r Replacer) *Transformer {
	return &Transformer{r, nil}
}

// Transform conducts transforming following its Replacer.
// Other specifications follow transform.Transformer.
func (tr *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nSrc = len(src)

	// joined the remained bytes which were not written into dst in the previous loop.
	replaced := Replace(tr.Rep, src)
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

// Replacer defines replacing method for a-z, A-Z and 0-9.
type Replacer interface {
	LowerFunc(rune) rune
	UpperFunc(rune) rune
	DigitFunc(rune) rune
}

// SimpleReplacer is an implementation of Replacer
// which just offsets a-z, A-Z and 0-9.
type SimpleReplacer struct {
	LowerOffset rune
	UpperOffset rune
	DigitOffset rune
}

func NewSimpleReplacer(lo, uo, do rune) *SimpleReplacer {
	return &SimpleReplacer{lo, uo, do}
}

func simpleReplace(src rune, offset rune) rune {
	replaced := src + offset
	// check whethere special replace is required
	alt, exist := altMap[replaced]
	if exist {
		replaced = alt
	}
	// check a minimal valid condition
	if !utf8.ValidRune(replaced) {
		return src
	}
	return replaced
}

func (sr *SimpleReplacer) LowerFunc(src rune) rune {
	return simpleReplace(src, sr.LowerOffset)
}
func (sr *SimpleReplacer) UpperFunc(src rune) rune {
	return simpleReplace(src, sr.UpperOffset)
}
func (sr *SimpleReplacer) DigitFunc(src rune) rune {
	return simpleReplace(src, sr.DigitOffset)
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
			rr := replaceByRune(rep, r)
			replaced = append(replaced, []byte(string(rr))...)
		}

		p = p[n:]
	}
	return replaced
}

// replaceByRune focuses on a valid []byte which could be decoded into a rune.
// Regular style a-zA-Z0-9 are replaced by specific functions.
func replaceByRune(rep Replacer, r rune) rune {
	if 'a' <= r && r <= 'z' {
		return rep.LowerFunc(r)
	}
	if 'A' <= r && r <= 'Z' {
		return rep.UpperFunc(r)
	}
	if '0' <= r && r <= '9' {
		return rep.DigitFunc(r)
	}
	return r
}
