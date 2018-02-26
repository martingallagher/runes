// +build !std

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
		IsDigit(pos)
		IsDigit(neg)
	}
}

func BenchmarkIsDigitUnsafe(b *testing.B) {
	pos, neg := getTestCase(digit)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsDigitUnsafe(pos)
		IsDigitUnsafe(neg)
	}
}

func BenchmarkIsGraphic(b *testing.B) {
	pos, neg := getTestCase(graphic)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsGraphic(pos)
		IsGraphic(neg)
	}
}

func BenchmarkIsGraphicUnsafe(b *testing.B) {
	pos, neg := getTestCase(graphic)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsGraphicUnsafe(pos)
		IsGraphicUnsafe(neg)
	}
}

func BenchmarkIsLetter(b *testing.B) {
	pos, neg := getTestCase(letter)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsLetter(pos)
		IsLetter(neg)
	}
}

func BenchmarkIsLetterUnsafe(b *testing.B) {
	pos, neg := getTestCase(letter)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsLetterUnsafe(pos)
		IsLetterUnsafe(neg)
	}
}

func BenchmarkIsLower(b *testing.B) {
	pos, neg := getTestCase(lower)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsLower(pos)
		IsLower(neg)
	}
}

func BenchmarkIsLowerUnsafe(b *testing.B) {
	pos, neg := getTestCase(lower)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsLowerUnsafe(pos)
		IsLowerUnsafe(neg)
	}
}

func BenchmarkIsMark(b *testing.B) {
	pos, neg := getTestCase(mark)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsMark(pos)
		IsMark(neg)
	}
}

func BenchmarkIsMarkUnsafe(b *testing.B) {
	pos, neg := getTestCase(mark)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsMarkUnsafe(pos)
		IsMarkUnsafe(neg)
	}
}

func BenchmarkIsNumber(b *testing.B) {
	pos, neg := getTestCase(number)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsNumber(pos)
		IsNumber(neg)
	}
}

func BenchmarkIsNumberUnsafe(b *testing.B) {
	pos, neg := getTestCase(number)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsNumberUnsafe(pos)
		IsNumberUnsafe(neg)
	}
}

func BenchmarkIsPrint(b *testing.B) {
	pos, neg := getTestCase(print)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsPrint(pos)
		IsPrint(neg)
	}
}

func BenchmarkIsPrintUnsafe(b *testing.B) {
	pos, neg := getTestCase(print)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsPrintUnsafe(pos)
		IsPrintUnsafe(neg)
	}
}

func BenchmarkIsPunct(b *testing.B) {
	pos, neg := getTestCase(punct)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsPunct(pos)
		IsPunct(neg)
	}
}

func BenchmarkIsPunctUnsafe(b *testing.B) {
	pos, neg := getTestCase(punct)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsPunctUnsafe(pos)
		IsPunctUnsafe(neg)
	}
}

func BenchmarkIsSpace(b *testing.B) {
	pos, neg := getTestCase(space)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsSpace(pos)
		IsSpace(neg)
	}
}

func BenchmarkIsSpaceUnsafe(b *testing.B) {
	pos, neg := getTestCase(space)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsSpaceUnsafe(pos)
		IsSpaceUnsafe(neg)
	}
}

func BenchmarkIsSymbol(b *testing.B) {
	pos, neg := getTestCase(symbol)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsSymbol(pos)
		IsSymbol(neg)
	}
}

func BenchmarkIsSymbolUnsafe(b *testing.B) {
	pos, neg := getTestCase(symbol)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsSymbolUnsafe(pos)
		IsSymbolUnsafe(neg)
	}
}

func BenchmarkIsTitle(b *testing.B) {
	pos, neg := getTestCase(title)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsTitle(pos)
		IsTitle(neg)
	}
}

func BenchmarkIsTitleUnsafe(b *testing.B) {
	pos, neg := getTestCase(title)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsTitleUnsafe(pos)
		IsTitleUnsafe(neg)
	}
}

func BenchmarkIsUpper(b *testing.B) {
	pos, neg := getTestCase(upper)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsUpper(pos)
		IsUpper(neg)
	}
}

func BenchmarkIsUpperUnsafe(b *testing.B) {
	pos, neg := getTestCase(upper)

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		IsUpperUnsafe(pos)
		IsUpperUnsafe(neg)
	}
}

func BenchmarkSimpleFold(b *testing.B) {
	pos, neg := '\u212A', 'a'

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		SimpleFold(pos)
		SimpleFold(neg)
	}
}

func BenchmarkSimpleFoldUnsafe(b *testing.B) {
	pos, neg := '\u212A', 'a'

	b.SetBytes(int64(utf8.RuneLen(pos) + utf8.RuneLen(neg)))
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		SimpleFoldUnsafe(pos)
		SimpleFoldUnsafe(neg)
	}
}

func BenchmarkToLower(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(ToLower, testString)
	}
}

func BenchmarkToLowerUnsafe(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(ToLowerUnsafe, testString)
	}
}

func BenchmarkToTitle(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(ToTitle, testString)
	}
}

func BenchmarkToTitleUnsafe(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(ToTitleUnsafe, testString)
	}
}

func BenchmarkToUpper(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(ToUpper, testString)
	}
}

func BenchmarkToUpperUnsafe(b *testing.B) {
	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(ToUpperUnsafe, testString)
	}
}

func BenchmarkStringToLowerCase(b *testing.B) {
	f := func(r rune) rune {
		return To(unicode.LowerCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToLowerCaseUnsafe(b *testing.B) {
	f := func(r rune) rune {
		return ToUnsafe(unicode.LowerCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToTitleCase(b *testing.B) {
	f := func(r rune) rune {
		return To(unicode.TitleCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToTitleCaseUnsafe(b *testing.B) {
	f := func(r rune) rune {
		return ToUnsafe(unicode.TitleCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToUpperCase(b *testing.B) {
	f := func(r rune) rune {
		return To(unicode.UpperCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}

func BenchmarkStringToUpperCaseUnsafe(b *testing.B) {
	f := func(r rune) rune {
		return ToUnsafe(unicode.UpperCase, r)
	}

	b.SetBytes(testStringLen)
	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		strings.Map(f, testString)
	}
}
