//ff:func feature=parse type=reader control=sequence
//ff:what gated — does the word before a tail pass the tail's graph, when it has one: the kind its stem stands as in this candidate is handed to it (all the kinds the word table gives the stem when the candidate has no kind); every graph is told too whether the stem and the piece right after it together are a word of the word table (the graphs of a suffix read after a nominal, suffixGate, and of a shortened particle, joinedGateOf, ask it); the graph of a suffix is told so too where the stem and the suffix are two words of the word table written together (twoWords), a number before a suffix of numberSuffixes (사, 용) is handed to it as a nominal (suffixedNumber: 3사), and so is a run of Latin letters and digits before any suffix (suffixedRun: OTS용), and a stem of one syllable before a suffix of longSuffixes (용, 제) is handed to it as no nominal (shortStem), and it is told where the suffix's letter is the prefix 제 of a number that follows (orderJoint: 요추제1)
package parse

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
	if len(t.pieces) > 0 {
		_, p.Whole = sc.ix.lex[stem+t.pieces[0].text]
	}
	if t.gate == sc.ix.suffixGate && len(t.pieces) > 0 {
		p.Whole, p.Order = p.Whole || sc.ix.twoWords(stem, t.pieces[0].text), sc.orderJoint(t)
	}
	if t.gate == sc.ix.suffixGate && (sc.suffixedNumber(t, e) || suffixedRun(e)) {
		p.Stem = kindNominal
	} else if t.gate == sc.ix.suffixGate && shortStem(t, stem) {
		p.Stem = 0
	}
	return sc.judge(t.gate, b, t.start, p) > 0
}
