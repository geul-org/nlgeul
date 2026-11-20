//ff:func feature=parse type=reader control=sequence
//ff:what TestGated — a tail with no gate always passes; a tail with the predicate gate passes the stem of a verb and not a noun's, and with the nominal gate the noun's; a stem of two kinds passes by the kind its candidate is given and, with no kind given, by any of them; the gate of a suffix after a nominal passes a nominal of the table and not a new word, and not where the stem and the piece after it together are a word of the table (the nominal gate still passes it)
package parse

import "testing"

func TestGated(t *testing.T) {
	sc := wordsScan("가신 학교신 곰신", layerCost)
	stem := func(text string, kind int) step {
		return step{kind: kind, word: Eojeol{Parts: []Part{{Text: text, Stem: true}}}}
	}
	if !sc.gated(0, tail{start: 1}, stem("가", 0)) {
		t.Fatal("no gate")
	}
	if !sc.gated(0, tail{start: 1, gate: sc.ix.pred}, stem("가", 0)) || sc.gated(3, tail{start: 5, gate: sc.ix.pred}, stem("학교", 0)) || sc.gated(7, tail{start: 8, gate: sc.ix.pred}, stem("곰", 0)) {
		t.Fatal("pred")
	}
	if !sc.gated(3, tail{start: 5, gate: sc.ix.nominal}, stem("학교", 0)) {
		t.Fatal("nominal")
	}
	if !sc.gated(0, tail{start: 1, gate: sc.ix.pred}, stem("가", kindVerb)) || sc.gated(0, tail{start: 1, gate: sc.ix.pred}, stem("가", kindNominal)) {
		t.Fatal("a stem of two kinds passes by the kind its candidate is given")
	}
	suffixed := tail{start: 5, gate: sc.ix.suffixGate, pieces: []piece{{text: "신", names: "접미:신"}}}
	if sc.ix.suffixGate == nil || !sc.gated(3, suffixed, stem("학교", 0)) || sc.gated(7, tail{start: 8, gate: sc.ix.suffixGate, pieces: suffixed.pieces}, stem("곰", 0)) {
		t.Fatal("the gate of a suffix after a nominal passes a nominal of the table and not a new word")
	}
	sc.ix.lex["학교신"] = word{kinds: kindNominal, count: 10}
	if sc.gated(3, suffixed, stem("학교", 0)) || !sc.gated(3, tail{start: 5, gate: sc.ix.nominal, pieces: suffixed.pieces}, stem("학교", 0)) {
		t.Fatal("the stem and the piece after it together are a word of the table: only the gate of a suffix turns it down")
	}
}
