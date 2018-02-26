package runes

import (
	"io/ioutil"
	"strings"
	"testing"
	"unicode"
)

var (
	testString    string
	testStringLen int64
)

func init() {
	b, err := ioutil.ReadFile("./testdata/corpus.txt")

	if err != nil {
		panic(err)
	}

	testString = string(b)
	testStringLen = int64(len(testString))
}

func TestFuncs(t *testing.T) {
	for r := rune(0); r <= unicode.MaxRune; r++ {
		if IsDigit(r) != unicode.IsDigit(r) {
			t.Fatalf("IsDigit failed for: %U, expected %v, got %v",
				r, unicode.IsDigit(r), IsDigit(r))
		}

		if IsDigitUnsafe(r) != unicode.IsDigit(r) {
			t.Fatalf("IsDigitUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsDigit(r), IsDigitUnsafe(r))
		}

		if IsGraphic(r) != unicode.IsGraphic(r) {
			t.Fatalf("IsGraphic failed for: %U, expected %v, got %v",
				r, unicode.IsGraphic(r), IsGraphic(r))
		}

		if IsGraphicUnsafe(r) != unicode.IsGraphic(r) {
			t.Fatalf("IsGraphicUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsGraphic(r), IsGraphicUnsafe(r))
		}

		if IsLetter(r) != unicode.IsLetter(r) {
			t.Fatalf("IsLetter failed for: %U, expected %v, got %v",
				r, unicode.IsLetter(r), IsLetter(r))
		}

		if IsLetterUnsafe(r) != unicode.IsLetter(r) {
			t.Fatalf("IsLetterUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsLetter(r), IsLetterUnsafe(r))
		}

		if IsLower(r) != unicode.IsLower(r) {
			t.Fatalf("IsLower failed for: %U, expected %v, got %v",
				r, unicode.IsLower(r), IsLower(r))
		}

		if IsLowerUnsafe(r) != unicode.IsLower(r) {
			t.Fatalf("IsLowerUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsLower(r), IsLowerUnsafe(r))
		}

		if IsMark(r) != unicode.IsMark(r) {
			t.Fatalf("IsMark failed for: %U, expected %v, got %v",
				r, unicode.IsMark(r), IsMark(r))
		}

		if IsMarkUnsafe(r) != unicode.IsMark(r) {
			t.Fatalf("IsMarkUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsMark(r), IsMarkUnsafe(r))
		}

		if IsNumber(r) != unicode.IsNumber(r) {
			t.Fatalf("IsNumber failed for: %U, expected %v, got %v",
				r, unicode.IsNumber(r), IsNumber(r))
		}

		if IsNumberUnsafe(r) != unicode.IsNumber(r) {
			t.Fatalf("IsNumberUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsNumber(r), IsNumberUnsafe(r))
		}

		if IsPrint(r) != unicode.IsPrint(r) {
			t.Fatalf("IsPrint failed for: %U, expected %v, got %v",
				r, unicode.IsPrint(r), IsPrint(r))
		}

		if IsPrintUnsafe(r) != unicode.IsPrint(r) {
			t.Fatalf("IsPrintUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsPrint(r), IsPrintUnsafe(r))
		}

		if IsPunct(r) != unicode.IsPunct(r) {
			t.Fatalf("IsPunct failed for: %U, expected %v, got %v",
				r, unicode.IsPunct(r), IsPunct(r))
		}

		if IsPunctUnsafe(r) != unicode.IsPunct(r) {
			t.Fatalf("IsPunctUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsPunct(r), IsPunctUnsafe(r))
		}

		if IsSpace(r) != unicode.IsSpace(r) {
			t.Fatalf("IsSpace failed for: %U, expected %v, got %v",
				r, unicode.IsSpace(r), IsSpace(r))
		}

		if IsSpaceUnsafe(r) != unicode.IsSpace(r) {
			t.Fatalf("IsSpaceUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsSpace(r), IsSpaceUnsafe(r))
		}

		if IsSymbol(r) != unicode.IsSymbol(r) {
			t.Fatalf("IsSymbol failed for: %U, expected %v, got %v",
				r, unicode.IsSymbol(r), IsSymbol(r))
		}

		if IsSymbolUnsafe(r) != unicode.IsSymbol(r) {
			t.Fatalf("IsSymbolUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsSymbol(r), IsSymbolUnsafe(r))
		}

		if IsTitle(r) != unicode.IsTitle(r) {
			t.Fatalf("IsTitle failed for: %U, expected %v, got %v",
				r, unicode.IsTitle(r), IsTitle(r))
		}

		if IsTitleUnsafe(r) != unicode.IsTitle(r) {
			t.Fatalf("IsTitleUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsTitle(r), IsTitleUnsafe(r))
		}

		if IsUpper(r) != unicode.IsUpper(r) {
			t.Fatalf("IsUpper failed for: %U, expected %v, got %v",
				r, unicode.IsUpper(r), IsUpper(r))
		}

		if IsUpperUnsafe(r) != unicode.IsUpper(r) {
			t.Fatalf("IsUpperUnsafe failed for: %U, expected %v, got %v",
				r, unicode.IsUpper(r), IsUpperUnsafe(r))
		}

		if SimpleFold(r) != unicode.SimpleFold(r) {
			t.Fatalf("SimpleFold failed for: %U, expected %v, got %v",
				r, unicode.SimpleFold(r), SimpleFold(r))
		}

		if ToLower(r) != unicode.ToLower(r) {
			t.Fatalf("Lower failed for: %U, expected %U, got %U",
				r, unicode.ToLower(r), ToLower(r))
		}

		if To(unicode.LowerCase, r) != unicode.To(unicode.LowerCase, r) {
			t.Fatalf("To failed for: %U, expected %U, got %U",
				r, unicode.To(unicode.LowerCase, r), To(unicode.LowerCase, r))
		}

		if ToTitle(r) != unicode.ToTitle(r) {
			t.Fatalf("Title failed for: %U, expected %U, got %U",
				r, unicode.ToTitle(r), ToTitle(r))
		}

		if To(unicode.TitleCase, r) != unicode.To(unicode.TitleCase, r) {
			t.Fatalf("To failed for: %U, expected %U, got %U",
				r, unicode.To(unicode.TitleCase, r), To(unicode.TitleCase, r))
		}

		if ToUpper(r) != unicode.ToUpper(r) {
			t.Fatalf("Upper failed for: %U, expected %U, got %U",
				r, unicode.ToUpper(r), ToUpper(r))
		}

		if To(unicode.UpperCase, r) != unicode.To(unicode.UpperCase, r) {
			t.Fatalf("To failed for: %U, expected %U, got %U",
				r, unicode.To(unicode.UpperCase, r), To(unicode.UpperCase, r))
		}
	}

	const (
		_case = 99
		r     = 'A'
	)

	if To(_case, r) != unicode.To(_case, r) {
		t.Fatalf("To failed for: %U, expected %U, got %U",
			r, unicode.To(_case, r), To(_case, r))
	}

	if strings.Map(ToLower, testString) != strings.Map(unicode.ToLower, testString) {
		t.Fatal("ToLower failed ")
	}

	if strings.Map(ToTitle, testString) != strings.Map(unicode.ToTitle, testString) {
		t.Fatal("ToTitle failed ")
	}

	if strings.Map(ToUpper, testString) != strings.Map(unicode.ToUpper, testString) {
		t.Fatal("ToUpper failed ")
	}
}
