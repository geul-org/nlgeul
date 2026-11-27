//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what approxTails — the tails of a unit after each approximate number written in Hangul that ends where the unit begins (칠팔년: 칠, 팔 and 년): the first number names the stem and the second is a piece before the unit's, both named as a number is (수:한글십진 by their letters); the unit's graph is asked again with whether the two numbers and the unit together are a word (오륙도)
package parse

import (
	"math"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Each numeral of the two numbers is taken as one of the table's
// entries, as a numeral before a unit is (numerals).
func (sc *scan) approxTails(e *entry, start, j int, facts rule.Piece, pc piece) []tail {
	var out []tail
	for _, pair := range sc.approxAt(start) {
		first, second := []rune(pair[0]), []rune(pair[1])
		facts.Whole, facts.Lone = sc.ix.wholeWord(pair[0]+pair[1], e.form), false
		q := pc
		q.verdict = sc.judge(e, start, j, facts)
		if q.verdict <= 0 {
			continue
		}
		cost := math.Log10(1.0 / numeralCount)
		num := piece{pair[1], numName, sc.ix.number, float64(len(second)) * cost, rule.TailAux, false, ""}
		out = append(out, tail{start: start - len(second), pieces: []piece{num, q}, only: pair[0], as: numName, weight: float64(len(first)) * cost})
	}
	return out
}
