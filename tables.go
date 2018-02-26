package runes

import (
	"unicode"
	"unicode/utf8"
)

const (
	digit = 1 << iota
	graphic
	letter
	lower
	mark
	number
	print
	punct
	space
	symbol
	title
	upper
)

type char struct {
	lower  rune
	title  rune
	upper  rune
	folded rune
	flags  uint
}

var properties [unicode.MaxRune + 1]char

func init() {
	for r := rune(0); r <= unicode.MaxRune; r++ {
		if !utf8.ValidRune(r) {
			continue
		}

		var v char

		if unicode.IsDigit(r) {
			v.flags |= digit
		}

		if unicode.IsGraphic(r) {
			v.flags |= graphic
		}

		if unicode.IsLetter(r) {
			v.flags |= letter
		}

		if unicode.IsLower(r) {
			v.flags |= lower
		}

		if unicode.IsMark(r) {
			v.flags |= mark
		}

		if unicode.IsNumber(r) {
			v.flags |= number
		}

		if unicode.IsPrint(r) {
			v.flags |= print
		}

		if unicode.IsPunct(r) {
			v.flags |= punct
		}

		if unicode.IsSpace(r) {
			v.flags |= space
		}

		if unicode.IsSymbol(r) {
			v.flags |= symbol
		}

		if unicode.IsTitle(r) {
			v.flags |= title
		}

		if unicode.IsUpper(r) {
			v.flags |= upper
		}

		v.lower = unicode.ToLower(r)
		v.title = unicode.ToTitle(r)
		v.upper = unicode.ToUpper(r)
		v.folded = unicode.SimpleFold(r)
		properties[r] = v
	}
}
