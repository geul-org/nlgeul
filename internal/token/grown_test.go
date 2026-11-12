//ff:func feature=token type=reader control=sequence
//ff:what TestGrown — each reading kept is grown by the word; none when the word before the tail does not pass the tail's graph; a tail read only after a modifying closing grows only the readings that end in one; where a word that ends in a modifying closing ends is kept; nothing grows where a dropped ㅡ is put back the stem does not drop (모르+아), where a word the table knows only as a nominal stands as a new word before a predicate tail (사과+다), or where a piece is taken out of a syllable right after a Latin letter (api고+아); the same words grow with a tail that breaks none of these; nothing grows before a new word's stem that ends in ㅆ with markers after it (곰졌+다, which grows without a word table); a verb of the table right after a modifying closing grows only the kept reading that does not end in one (마시+다), a noun grows both (사과+를)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGrown(t *testing.T) {
	sc := scanOf("가는")
	kept := []*path{{verdict: 1}, {verdict: 1, last: step{adn: true}}}
	adn := tail{start: 1, pieces: []piece{{text: "는", names: "꾸밈:는", verdict: 1}}}
	if got := sc.grown(nil, kept, adn, 0, 2); len(got) != 2 || !sc.adnEnds[2] {
		t.Fatal(got, sc.adnEnds)
	}
	bound := scanOf("가건")
	bound.adnEnds[1] = true
	if got := bound.grown(nil, kept, bound.fixed(2, -1)[0], 1, 2); len(got) != 1 || got[0].prev != kept[1] {
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
	da := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	if got := wordsScan("곰졌다", layerCost).grown(nil, kept, tail{start: 2, pieces: da}, 0, 3); len(got) != 0 {
		t.Fatal("곰졌, a new word that ends in ㅆ, before 다", got)
	}
	if got := scanOf("곰졌다").grown(nil, plain, tail{start: 2, pieces: da}, 0, 3); len(got) != 1 {
		t.Fatal("the same new word without a word table", got)
	}
	verb := wordsScan("먹을마시다", layerCost)
	verb.adnEnds[2] = true
	if got := verb.grown(nil, kept, tail{start: 4, pieces: da}, 2, 5); len(got) != 1 || got[0].prev != kept[0] || verb.adnEnds[5] {
		t.Fatal("마시+다 grows only the reading that does not end in a modifying closing", got)
	}
	noun2 := wordsScan("먹을사과를", layerCost)
	noun2.adnEnds[2] = true
	if got := noun2.grown(nil, kept, tail{start: 4, pieces: []piece{{text: "를", names: "목적격:를", verdict: 1, tail: rule.TailCase}}}, 2, 5); len(got) != 2 {
		t.Fatal("사과+를 grows both readings", got)
	}
}
