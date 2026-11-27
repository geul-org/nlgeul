//ff:func feature=parse type=helper control=sequence
//ff:what approxPair — are two runs of Sino-Korean numerals written in Hangul, one right after the other, the two numbers of an approximate number (육칠, 이삼십, 십오륙): the first ends in a digit and the second begins with the digit next after it (륙 for 육 right after 오, as it is written there)
package parse

import (
	"slices"
	"strings"
)

// The second number may go on to a greater number (이삼십 is 이 and
// 삼십, 육칠만 is 육 and 칠만): the whole of it is one number, as a run
// of numerals before a unit is one stem (삼십 before 명).
func approxPair(first, second string) bool {
	if rest, ok := strings.CutPrefix(second, sinoSixAfter); ok && strings.HasSuffix(first, sinoFive) {
		second = sinoSix + rest
	}
	if sinoTokens(first) == 0 || sinoTokens(second) == 0 {
		return false
	}
	a, digits := []rune(first), []rune(sinoDigits)
	last := slices.Index(digits, a[len(a)-1])
	return last >= 0 && last+1 < len(digits) && []rune(second)[0] == digits[last+1]
}
