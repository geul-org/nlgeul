//ff:func feature=token type=reader control=sequence
//ff:what gated — does the word before a tail pass the tail's graph, when it has one: the word's kinds in the word table are handed to it
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) gated(b int, t tail) bool {
	if t.gate == nil {
		return true
	}
	stem, _, _ := sc.stemOf(b, t)
	kinds := sc.ix.lex[stem].kinds | sc.ix.closed[stem]
	return sc.judge(t.gate, b, t.start, rule.Piece{Form: stem, HasStem: true, Next: -1, Stem: kinds}) > 0
}
