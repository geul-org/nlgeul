//ff:func feature=token type=reader control=sequence
//ff:what TestGrown — each reading kept is grown by the word; none when the word before the tail does not pass the tail's graph; a tail read only after a modifying closing grows only the readings that end in one; where a word that ends in a modifying closing ends is kept
package token

import "testing"

func TestGrown(t *testing.T) {
	sc := scanOf("가는")
	kept := []*path{{verdict: 1}, {verdict: 1, last: eojeol{adn: true}}}
	adn := tail{start: 1, pieces: []piece{{text: "는", names: "꾸밈:는", verdict: 1}}}
	if got := sc.grown(nil, kept, adn, 0, 2); len(got) != 2 || !sc.adnEnds[2] {
		t.Fatal(got, sc.adnEnds)
	}
	if got := sc.grown(nil, kept, tail{start: 2, adn: true}, 0, 2); len(got) != 1 || got[0].prev != kept[1] {
		t.Fatal("only after a modifying closing", got)
	}
	ws := wordsScan("학교래", layerCost)
	if got := ws.grown(nil, kept, tail{start: 2, gate: ws.ix.pred, pieces: []piece{{text: "래"}}}, 0, 3); len(got) != 0 {
		t.Fatal("a noun before a tail that asks for a predicate", got)
	}
}
