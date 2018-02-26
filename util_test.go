package runes

import "unicode"

func getTestCase(t uint) (rune, rune) {
	var pos, neg rune = -1, -1

	for r := unicode.MaxRune; r > -1; r-- {
		if pos == -1 && is(t, r) {
			pos = r
		}

		if neg == -1 && !is(t, r) {
			neg = r
		}

		if pos != -1 && neg != -1 {
			break
		}
	}

	return pos, neg
}
