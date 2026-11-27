//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what hanNumerals — the number written in Han characters that ends at a place: the whole run of Han numerals that ends there (二三, 十), named as a number is (수:한자십진 by its letters); none when a Han character that is no numeral stands right before the run (the numeral is a part of a word written in Han characters: 第一, 張三)
package parse

import (
	"math"
	"unicode"
)

// Each letter of the run is taken as one of the table's letters.
func (sc *scan) hanNumerals(end int) []numeral {
	start := end
	for start > 0 && hanRun(string(sc.s[start-1])) {
		start--
	}
	if start == end || (start > 0 && unicode.Is(unicode.Han, sc.s[start-1])) {
		return nil
	}
	return []numeral{{string(sc.s[start:end]), float64(end-start) * math.Log10(1.0/float64(len(hanNumerals))), numName}}
}
