//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what approxLone — the tails of an approximate number written in Hangul that ends at j with no unit written onto it (육칠: 육 and 칠): the first number names the stem and the second is the one piece, both named as a number is (수:한글십진 by their letters), with the verdict of the graph loneNumberGraph, a new word's (it stands behind the word of the word table: before a unit written apart the letters may be its number, 육칠 명); read only with a word table (with none, nothing tells the numerals from the syllables of any word)
package parse

import (
	"math"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Each letter is valued as a letter of a new word is: the reading
// stands with the new word of the same letters, not before a word of
// the word table.
func (sc *scan) approxLone(j int) []tail {
	var out []tail
	if !sc.ix.lexed {
		return nil
	}
	cost := math.Log10(1.0 / syllables)
	for _, pair := range sc.approxAt(j) {
		first, second := []rune(pair[0]), []rune(pair[1])
		num := piece{pair[1], numName, sc.ix.loneNumber, float64(len(second)) * cost, rule.TailAux, false, ""}
		out = append(out, tail{start: j - len(second), pieces: []piece{num}, only: pair[0], as: numName, weight: float64(len(first)) * cost})
	}
	return out
}
