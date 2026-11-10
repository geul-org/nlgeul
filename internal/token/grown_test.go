//ff:func feature=token type=reader control=sequence
//ff:what TestGrown — each reading kept is grown by the word; none when the word before the tail does not pass the tail's graph; a tail read only after a modifying closing grows only the readings that end in one; where a word that ends in a modifying closing ends is kept; nothing grows where a dropped ㅡ is put back the stem does not drop (모르+아), where a word the table knows only as a nominal stands as a new word before a predicate tail (사과+다), or where a piece is taken out of a syllable right after a Latin letter (api고+아); the same words grow with a tail that breaks none of these
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestGrown(t *testing.T) {
	sc := scanOf("가는")
	kept := []*path{{verdict: 1}, {verdict: 1, last: step{adn: true}}}
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
	fused := []piece{{text: "(아)", names: "맺음:어(합침)", verdict: 1, tail: rule.TailPred}}
	plain := []*path{{verdict: 1}}
	eu := wordsScan("모라", layerCost)
	if got := eu.grown(nil, plain, tail{start: 2, base: '르', pieces: fused}, 0, 2); len(got) != 0 {
		t.Fatal("모르 is 르 irregular: no 모르+아", got)
	}
	if got := eu.grown(nil, plain, tail{start: 2, base: '라', pieces: fused}, 0, 2); len(got) != 1 {
		t.Fatal("모라+아 is no dropped ㅡ", got)
	}
	noun := wordsScan("사과다", layerCost)
	if got := noun.grown(nil, plain, tail{start: 2, pieces: []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}}, 0, 3); len(got) != 0 {
		t.Fatal("사과, known only as a nominal, as a new word before a predicate tail", got)
	}
	if got := noun.grown(nil, plain, tail{start: 2, pieces: []piece{{text: "다", verdict: 1, tail: rule.TailCase}}}, 0, 3); len(got) != 1 {
		t.Fatal("the same nominal before a particle", got)
	}
	latin := scanOf("api과")
	if got := latin.grown(nil, plain, tail{start: 4, base: '고', pieces: fused}, 0, 4); len(got) != 0 {
		t.Fatal("api고+아", got)
	}
	if got := latin.grown(nil, plain, tail{start: 4, base: '고', pieces: fused}, 3, 4); len(got) != 1 {
		t.Fatal("고+아 after the boundary", got)
	}
}
