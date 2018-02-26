// +build std

package runes

import (
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
)

func BenchmarkIsDigit(b *testing.B) {
	pos, neg := getTestCase(digit)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsDigit(pos)
		unicode.IsDigit(neg)
	}
}

func BenchmarkIsDigitUnsafe(b *testing.B) {
	pos, neg := getTestCase(digit)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsDigit(pos)
		unicode.IsDigit(neg)
	}
}

func BenchmarkIsGraphic(b *testing.B) {
	pos, neg := getTestCase(graphic)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsGraphic(pos)
		unicode.IsGraphic(neg)
	}
}

func BenchmarkIsGraphicUnsafe(b *testing.B) {
	pos, neg := getTestCase(graphic)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsGraphic(pos)
		unicode.IsGraphic(neg)
	}
}

func BenchmarkIsLetter(b *testing.B) {
	pos, neg := getTestCase(letter)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsLetter(pos)
		unicode.IsLetter(neg)
	}
}

func BenchmarkIsLetterUnsafe(b *testing.B) {
	pos, neg := getTestCase(letter)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsLetter(pos)
		unicode.IsLetter(neg)
	}
}

func BenchmarkIsLower(b *testing.B) {
	pos, neg := getTestCase(lower)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsLower(pos)
		unicode.IsLower(neg)
	}
}

func BenchmarkIsLowerUnsafe(b *testing.B) {
	pos, neg := getTestCase(lower)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsLower(pos)
		unicode.IsLower(neg)
	}
}

func BenchmarkIsMark(b *testing.B) {
	pos, neg := getTestCase(mark)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsMark(pos)
		unicode.IsMark(neg)
	}
}

func BenchmarkIsMarkUnsafe(b *testing.B) {
	pos, neg := getTestCase(mark)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsMark(pos)
		unicode.IsMark(neg)
	}
}

func BenchmarkIsNumber(b *testing.B) {
	pos, neg := getTestCase(number)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsNumber(pos)
		unicode.IsNumber(neg)
	}
}

func BenchmarkIsNumberUnsafe(b *testing.B) {
	pos, neg := getTestCase(number)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsNumber(pos)
		unicode.IsNumber(neg)
	}
}

func BenchmarkIsPrint(b *testing.B) {
	pos, neg := getTestCase(print)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsPrint(pos)
		unicode.IsPrint(neg)
	}
}

func BenchmarkIsPrintUnsafe(b *testing.B) {
	pos, neg := getTestCase(print)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsPrint(pos)
		unicode.IsPrint(neg)
	}
}

func BenchmarkIsPunct(b *testing.B) {
	pos, neg := getTestCase(punct)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsPunct(pos)
		unicode.IsPunct(neg)
	}
}

func BenchmarkIsPunctUnsafe(b *testing.B) {
	pos, neg := getTestCase(punct)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsPunct(pos)
		unicode.IsPunct(neg)
	}
}

func BenchmarkIsSpace(b *testing.B) {
	pos, neg := getTestCase(space)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsSpace(pos)
		unicode.IsSpace(neg)
	}
}

func BenchmarkIsSpaceUnsafe(b *testing.B) {
	pos, neg := getTestCase(space)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsSpace(pos)
		unicode.IsSpace(neg)
	}
}

func BenchmarkIsSymbol(b *testing.B) {
	pos, neg := getTestCase(symbol)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsSymbol(pos)
		unicode.IsSymbol(neg)
	}
}

func BenchmarkIsSymbolUnsafe(b *testing.B) {
	pos, neg := getTestCase(symbol)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsSymbol(pos)
		unicode.IsSymbol(neg)
	}
}

func BenchmarkIsTitle(b *testing.B) {
	pos, neg := getTestCase(title)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsTitle(pos)
		unicode.IsTitle(neg)
	}
}

func BenchmarkIsTitleUnsafe(b *testing.B) {
	pos, neg := getTestCase(title)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsTitle(pos)
		unicode.IsTitle(neg)
	}
}

func BenchmarkIsUpper(b *testing.B) {
	pos, neg := getTestCase(upper)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsUpper(pos)
		unicode.IsUpper(neg)
	}
}

func BenchmarkIsUpperUnsafe(b *testing.B) {
	pos, neg := getTestCase(upper)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.IsUpper(pos)
		unicode.IsUpper(neg)
	}
}

func BenchmarkSimpleFold(b *testing.B) {
	pos, neg := '\u212A', 'a'

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.SimpleFold(pos)
		unicode.SimpleFold(neg)
	}
}

func BenchmarkSimpleFoldUnsafe(b *testing.B) {
	pos, neg := '\u212A', 'a'

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		unicode.SimpleFold(pos)
		unicode.SimpleFold(neg)
	}
}

func BenchmarkToLower(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(unicode.ToLower, testString)
	}
}

func BenchmarkToLowerUnsafe(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(unicode.ToLower, testString)
	}
}

func BenchmarkToTitle(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(unicode.ToTitle, testString)
	}
}

func BenchmarkToTitleUnsafe(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(unicode.ToTitle, testString)
	}
}

func BenchmarkToUpper(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(unicode.ToUpper, testString)
	}
}

func BenchmarkToUpperUnsafe(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(unicode.ToUpper, testString)
	}
}

func BenchmarkStringToLowerCase(b *testing.B) {
	f := func(r rune) rune {
		return unicode.To(unicode.LowerCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToLowerCaseUnsafe(b *testing.B) {
	f := func(r rune) rune {
		return unicode.To(unicode.LowerCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToTitleCase(b *testing.B) {
	f := func(r rune) rune {
		return unicode.To(unicode.TitleCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToTitleCaseUnsafe(b *testing.B) {
	f := func(r rune) rune {
		return unicode.To(unicode.TitleCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToUpperCase(b *testing.B) {
	f := func(r rune) rune {
		return unicode.To(unicode.UpperCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToUpperCaseUnsafe(b *testing.B) {
	f := func(r rune) rune {
		return unicode.To(unicode.UpperCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}
