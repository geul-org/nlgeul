//ff:func feature=token type=reader control=sequence
//ff:what TestExtend — a tail grows every reading kept before its stem, from each place before it; a tail that attaches to one stem only grows no other; a tail that asks for a predicate's stem grows only after one; a tail read only after a modifying closing grows only the readings whose last word ends in one, and where such a word ends is kept
package token

import "testing"

func TestExtend(t *testing.T) {
	sc := scanOf("곰이")
	best := [][]*path{{{verdict: 1}}, nil, nil}
	if got := sc.extend(nil, best, tail{start: 2, pieces: []piece{{text: "x"}}}, 2); len(got) != 1 {
		t.Fatal("a place with no reading grown", got)
	}
	inner := scanOf("ab는")
	if got := inner.extend(nil, [][]*path{{{verdict: 1}}, {{verdict: 1}}, nil, nil}, tail{start: 2, pieces: []piece{{text: "는"}}}, 3); len(got) != 1 {
		t.Fatal("a word beginning inside an atom", got)
	}
	if got := inner.extend(nil, [][]*path{{{verdict: 1}}, nil, nil, nil}, tail{start: 1, pieces: []piece{{text: "x"}}}, 3); len(got) != 0 {
		t.Fatal("a stem ending inside an atom", got)
	}
	if got := sc.extend(nil, best, tail{start: 2, only: "아니", pieces: []piece{{text: "x"}}}, 2); len(got) != 0 {
		t.Fatal("a tail of 아니 after another stem", got)
	}
	got := sc.extend(nil, best, tail{start: 1, pieces: []piece{{"이", "주격:이", 1, -1, 0, false}}}, 2)
	if len(got) != 1 || got[0].sig != "곰+이[주격:이]" {
		t.Fatal(got)
	}
	ws := wordsScan("학교래", layerCost)
	if got := ws.extend(nil, [][]*path{{{verdict: 1}}, nil, nil, nil}, tail{start: 2, gate: ws.ix.pred, pieces: []piece{{text: "래"}}}, 3); len(got) != 0 {
		t.Fatal("a noun before a closing that asks for a predicate", got)
	}
	adn := scanOf("간곳")
	after := [][]*path{{{verdict: 1}}, {{verdict: 1, last: step{end: 1}}, {verdict: 1, last: step{end: 1, adn: true}}}, nil}
	if got := adn.extend(nil, after, tail{start: 2, adn: true}, 2); len(got) != 1 || got[0].prev != after[1][1] {
		t.Fatal("a tail read only after a modifying closing", got)
	}
	mark := scanOf("가는")
	mark.extend(nil, [][]*path{{{verdict: 1}}, nil, nil}, tail{start: 1, pieces: []piece{{text: "는", names: "꾸밈:는", verdict: 1}}}, 2)
	if !mark.adnEnds[2] || mark.adnEnds[1] {
		t.Fatal("where a word that ends in a modifying closing ends", mark.adnEnds)
	}
}
