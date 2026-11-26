//ff:func feature=parse type=reader control=sequence
//ff:what TestGated — a tail with no gate always passes; a tail with the predicate gate passes the stem of a verb and not a noun's, and with the nominal gate the noun's; a stem of two kinds passes by the kind its candidate is given and, with no kind given, by any of them; the gate of a suffix after a nominal passes a nominal of the table and not a new word, and not where the stem and the piece after it together are a word of the table (the nominal gate still passes it), nor where they are two words of the table written together (twoWords); a number stem before the suffix 사 passes the gate of a suffix and before 률 does not (suffixedNumber)
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
	delete(sc.ix.lex, "학교신")
	sc.ix.lex["학교곰"], sc.ix.lex["곰신"] = word{kinds: kindNominal, count: 10}, word{kinds: kindNominal, count: 10}
	two := tail{start: 6, gate: sc.ix.suffixGate, pieces: suffixed.pieces}
	if sc.gated(3, two, stem("학교곰", 0)) || !sc.gated(3, tail{start: 6, gate: sc.ix.nominal, pieces: suffixed.pieces}, stem("학교곰", 0)) || !sc.gated(3, tail{start: 6, gate: sc.ix.suffixGate}, stem("학교곰", 0)) {
		t.Fatal("the stem and the suffix are two words of the table written together: only the gate of a suffix turns it down, and only with a piece")
	}
	sc = wordsScan("3사 3률", layerCost)
	number := step{word: Eojeol{Parts: []Part{{Text: "3", Stem: true, Name: Name{Kind: "수"}}}}}
	sa := tail{start: 1, gate: sc.ix.suffixGate, pieces: []piece{{text: "사", names: "접미:사"}}}
	ryul := tail{start: 4, gate: sc.ix.suffixGate, pieces: []piece{{text: "률", names: "접미:률"}}}
	if !sc.gated(0, sa, number) || sc.gated(3, ryul, number) {
		t.Fatal("a number stem passes the gate of a suffix before 사 and not before 률")
	}
}
