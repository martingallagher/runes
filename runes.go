//go:generate go run cmd/bootstrap/main.go

// Package runes provides functions to test some properties of
// Unicode code points.
package runes

import (
	"unicode"
	"unicode/utf8"
)

// IsControl reports whether the rune is a control character.
// The C (Other) Unicode category includes more code points
// such as surrogates; use Is(C, r) to test for them.
func IsControl(r rune) bool {
	// Included for symmetry
	return unicode.IsControl(r)
}

// IsDigit reports whether the rune is a decimal digit.
func IsDigit(r rune) bool {
	return is(digit, r)
}

// IsDigitUnsafe is the unsafe version of IsDigit.
func IsDigitUnsafe(r rune) bool {
	return isUnsafe(digit, r)
}

// IsGraphic reports whether the rune is defined as a Graphic by Unicode.
// Such characters include letters, marks, numbers, punctuation, symbols, and
// spaces, from categories L, M, N, P, S, Zs.
func IsGraphic(r rune) bool {
	return is(graphic, r)
}

// IsGraphicUnsafe is the unsafe version of IsGraphic.
func IsGraphicUnsafe(r rune) bool {
	return isUnsafe(graphic, r)
}

// IsLetter reports whether the rune is a letter (category L).
func IsLetter(r rune) bool {
	return is(letter, r)
}

// IsLetterUnsafe is the unsafe version of IsLetter.
func IsLetterUnsafe(r rune) bool {
	return isUnsafe(letter, r)
}

// IsLower reports whether the rune is a lower case letter.
func IsLower(r rune) bool {
	return is(lower, r)
}

// IsLowerUnsafe is the unsafe version of IsLower.
func IsLowerUnsafe(r rune) bool {
	return isUnsafe(lower, r)
}

// IsMark reports whether the rune is a mark character (category M).
func IsMark(r rune) bool {
	return is(mark, r)
}

// IsMarkUnsafe is the unsafe version of IsMark.
func IsMarkUnsafe(r rune) bool {
	return isUnsafe(mark, r)
}

// IsNumber reports whether the rune is a number (category N).
func IsNumber(r rune) bool {
	return is(number, r)
}

// IsNumberUnsafe is the unsafe version of IsNumber.
func IsNumberUnsafe(r rune) bool {
	return isUnsafe(number, r)
}

// IsPrint reports whether the rune is defined as printable by Go. Such
// characters include letters, marks, numbers, punctuation, symbols, and the
// ASCII space character, from categories L, M, N, P, S and the ASCII space
// character. This categorization is the same as IsGraphic except that the
// only spacing character is ASCII space, U+0020.
func IsPrint(r rune) bool {
	return is(print, r)
}

// IsPrintUnsafe is the unsafe version of IsPrint.
func IsPrintUnsafe(r rune) bool {
	return isUnsafe(print, r)
}

// IsPunct reports whether the rune is a Unicode punctuation character
// (category P).
func IsPunct(r rune) bool {
	return is(punct, r)
}

// IsPunctUnsafe is the unsafe version of IsPunct.
func IsPunctUnsafe(r rune) bool {
	return isUnsafe(punct, r)
}

// IsSpace reports whether the rune is a space character as defined
// by Unicode's White Space property; in the Latin-1 space
// this is
//	'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
// Other definitions of spacing characters are set by category
// Z and property Pattern_White_Space.
func IsSpace(r rune) bool {
	return is(space, r)
}

// IsSpaceUnsafe is the unsafe version of IsSpace.
func IsSpaceUnsafe(r rune) bool {
	return isUnsafe(space, r)
}

// IsSymbol reports whether the rune is a symbolic character.
func IsSymbol(r rune) bool {
	return is(symbol, r)
}

// IsSymbolUnsafe is the unsafe version of IsSymbol.
func IsSymbolUnsafe(r rune) bool {
	return isUnsafe(symbol, r)
}

// IsTitle reports whether the rune is a title case letter.
func IsTitle(r rune) bool {
	return is(title, r)
}

// IsTitleUnsafe is the unsafe version of IsTitle.
func IsTitleUnsafe(r rune) bool {
	return isUnsafe(title, r)
}

// IsUpper reports whether the rune is an upper case letter.
func IsUpper(r rune) bool {
	return is(upper, r)
}

// IsUpperUnsafe is the unsafe version of IsUpper.
func IsUpperUnsafe(r rune) bool {
	return isUnsafe(upper, r)
}

// SimpleFold iterates over Unicode code points equivalent under
// the Unicode-defined simple case folding. Among the code points
// equivalent to rune (including rune itself), SimpleFold returns the
// smallest rune > r if one exists, or else the smallest rune >= 0.
// If r is not a valid Unicode code point, SimpleFold(r) returns r.
//
// For example:
//	SimpleFold('A') = 'a'
//	SimpleFold('a') = 'A'
//
//	SimpleFold('K') = 'k'
//	SimpleFold('k') = '\u212A' (Kelvin symbol, K)
//	SimpleFold('\u212A') = 'K'
//
//	SimpleFold('1') = '1'
//
//	SimpleFold(-2) = -2
//
func SimpleFold(r rune) rune {
	if !utf8.ValidRune(r) {
		return r
	}

	return properties[r].folded
}

// SimpleFoldUnsafe is the unsafe version of SimpleFold.
func SimpleFoldUnsafe(r rune) rune {
	return properties[r].folded
}

func isUnsafe(t uint, r rune) bool {
	return properties[r].flags&t != 0
}

func is(t uint, r rune) bool {
	return r <= unicode.MaxRune && properties[r].flags&t != 0
}

// To maps the rune to the specified case: UpperCase, LowerCase, or TitleCase.
func To(_case int, r rune) rune {
	switch _case {
	case unicode.LowerCase:
		return ToLower(r)
	case unicode.TitleCase:
		return ToTitle(r)
	case unicode.UpperCase:
		return ToUpper(r)
	}

	// Emulate standard library behaviour
	return unicode.ReplacementChar
}

// ToUnsafe is the unsafe version of To.
func ToUnsafe(_case int, r rune) rune {
	switch _case {
	case unicode.LowerCase:
		return ToLowerUnsafe(r)
	case unicode.TitleCase:
		return ToTitleUnsafe(r)
	case unicode.UpperCase:
		return ToUpperUnsafe(r)
	}

	// Emulate standard library behaviour
	return unicode.ReplacementChar
}

// ToLower maps the rune to lower case.
func ToLower(r rune) rune {
	if !utf8.ValidRune(r) {
		return r
	}

	return properties[r].lower
}

// ToLowerUnsafe is the unsafe version of ToLower.
func ToLowerUnsafe(r rune) rune {
	return properties[r].lower
}

// ToTitle maps the rune to title case.
func ToTitle(r rune) rune {
	if !utf8.ValidRune(r) {
		return r
	}

	return properties[r].title
}

// ToTitleUnsafe is the unsafe version of ToTitle.
func ToTitleUnsafe(r rune) rune {
	return properties[r].title
}

// ToUpper maps the rune to upper case.
func ToUpper(r rune) rune {
	if !utf8.ValidRune(r) {
		return r
	}

	return properties[r].upper
}

// ToUpperUnsafe is the unsafe version of ToUpper.
func ToUpperUnsafe(r rune) rune {
	return properties[r].upper
}
