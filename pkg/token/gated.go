//ff:func feature=token type=reader control=sequence
//ff:what gated — does the word before a tail pass the tail's graph, when it has one: the kind its stem stands as in this candidate is handed to it (all the kinds the word table gives the stem when the candidate has no kind); the graph of a suffix read after a nominal (suffixGate) is told too whether the stem and the piece right after it together are a word of the word table
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) gated(b int, t tail, e step) bool {
	if t.gate == nil {
		return true
	}
	stem := e.word.Parts[0].Text
	kinds := sc.ix.lex[stem].kinds | sc.ix.closed[stem]
	if e.kind != 0 {
		kinds = e.kind
	}
	p := rule.Piece{Form: stem, HasStem: true, Next: -1, Stem: kinds}
	if t.gate == sc.ix.suffixGate {
		_, p.Whole = sc.ix.lex[stem+t.pieces[0].text]
	}
	return sc.judge(t.gate, b, t.start, p) > 0
}
