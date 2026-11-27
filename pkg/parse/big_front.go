//ff:func feature=parse type=reader control=sequence
//ff:what bigFront — the candidate of a stem that is one greater number written in Hangul (십, 백, 천, 만, 억, 조) right before a run of digits of the same number (천500억, 천6백여): named as a greater number after digits is (큰 수:천), with the verdict of the graph bigFrontGraph; the letter is read so only where no numeral stands right before it, nothing is written onto it, and a greater number is written right after the digits (the number goes on: 500억, 6백)
package parse

import "strings"

// A number written in digits and greater numbers of Hangul together
// (1천500억) may leave the 1 before its first greater number unwritten.
// The greater number after the digits tells that number from 만 the
// determiner written onto an age or a span (만19세) and from 조 the noun
// (조1위): there a unit follows the digits.
func (sc *scan) bigFront(b int, stem string, t tail) []kindCandidate {
	a, ok := sc.atoms[t.start]
	if !ok || !a.num || len(t.pieces) > 0 || len([]rune(stem)) != 1 || !strings.Contains(sinoBigs, stem) || a.end >= len(sc.s) || !strings.ContainsRune(sinoBigs, sc.s[a.end]) || sc.numEnds[b] || sc.numeralBefore(b) {
		return nil
	}
	return []kindCandidate{{0, bigName + stem, sc.ix.bigFront}}
}
