//ff:func feature=token type=reader control=sequence
//ff:what TestGated — a tail with no gate always passes; a tail with the predicate gate passes the stem of a verb and not a noun's, and with the nominal gate the noun's; a stem of two kinds passes by the kind its reading is given and, with no kind given, by any of them
package token

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
		t.Fatal("a stem of two kinds passes by the kind its reading is given")
	}
}
