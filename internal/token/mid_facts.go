//ff:func feature=token type=reader control=sequence
//ff:what midFacts — the facts of a stretch of one sign between letters of a unit, from start to stop in the unit's letters: how many stand in a row, what stands on each side, and whether a punctuation mark or symbol stands right next to it
package token

import (
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) midFacts(start, stop int) rule.Mark {
	m := rule.Mark{Run: stop - start, Left: sideOf(string(sc.at(start - 1))), Right: sideOf(string(sc.at(stop)))}
	for _, r := range []rune{sc.at(start - 1), sc.at(stop)} {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			m.Facts |= rule.MarkBySign
		}
	}
	return m
}
