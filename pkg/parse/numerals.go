//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what numerals — the numbers that end at a place and are no run of digits: every run of Hangul numerals that ends there, however long (as far back as the letters of numerals go), a Latin letter of the table that is a run by itself, and the run of Han numerals that ends there (hanNumerals)
package parse

import (
	"math"
	"slices"
)

// Each numeral of a run is taken as one of the table's entries, and a
// letter as one of the letters.
func (sc *scan) numerals(end int) []numeral {
	var out []numeral
	for n := 1; n <= end && numeralLetter(sc.s[end-n]); n++ {
		text := string(sc.s[end-n : end])
		if k := max(nativeTokens(text), sinoTokens(text)); k > 0 {
			out = append(out, numeral{text, float64(k) * math.Log10(1.0/numeralCount), numName})
		}
	}
	a, ok := sc.atoms[end-1]
	if ok && a.end == end && slices.Contains(numberLetters, string(sc.s[end-1:end])) {
		out = append(out, numeral{string(sc.s[end-1 : end]), math.Log10(1.0 / float64(len(numberLetters))), numName})
	}
	return append(out, sc.hanNumerals(end)...)
}
