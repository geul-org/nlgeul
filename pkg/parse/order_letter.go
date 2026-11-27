//ff:func feature=parse type=reader control=sequence
//ff:what orderLetter — the letter of orderLetters right before place start, taken as a number named 차례:한글 (orderLetterName: a Hangul letter that tells an order, before the unit 목), when a word begins at it (nothing or a sign before it) or it stands right after the end of a number (제1호가목)
package parse

import (
	"math"
	"strings"
	"unicode/utf8"
)

func (sc *scan) orderLetter(start int) []numeral {
	if start < 1 || !strings.ContainsRune(orderLetters, sc.s[start-1]) {
		return nil
	}
	if start >= 2 && sideOf(string(sc.s[start-2])) != 0 && !sc.numberBefore(start-1) {
		return nil
	}
	return []numeral{{string(sc.s[start-1]), math.Log10(1 / float64(utf8.RuneCountInString(orderLetters))), orderLetterName}}
}
