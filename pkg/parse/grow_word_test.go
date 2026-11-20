//ff:func feature=parse type=reader control=sequence
//ff:what TestGrowWord — each kept candidate is grown by the word that tail t makes; a word of no verdict grows none; a tail read only after a modifying closing grows only the candidates on which its graph stands; the candidates that end in a modifying closing or a determiner are not grown by a predicate; a word that ends in a modifying closing marks where it ends, and so does one that ends in the linking 어/아 (and neither marks the other's); a number in binary has two candidates
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGrowWord(t *testing.T) {
	kept := []*path{{verdict: 1}, {verdict: 1, last: step{adn: true}}, {verdict: 1, last: step{det: true}}}
	da := tail{start: 4, pieces: []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}}
	verb := wordsScan("먹을마시다", layerCost)
	verb.adnEnds[2], verb.detEnds[2] = true, true
	e := verb.stepsOf(2, 5, da)[0]
	if got := verb.growWord(nil, kept, da, e, 2, 5); len(got) != 1 || got[0].prev != kept[0] || verb.adnEnds[5] {
		t.Fatal("a predicate grows only the candidate that ends in no modifying closing and no determiner", got)
	}
	dead := e
	dead.verdict = 0
	if got := verb.growWord(nil, kept, da, dead, 2, 5); len(got) != 0 {
		t.Fatal("a word of no verdict", got)
	}
	turn := tail{start: 4, pieces: []piece{{text: "는", names: "꾸밈:는", verdict: 1, tail: rule.TailPred}}}
	e = verb.stepsOf(2, 5, turn)[0]
	if got := verb.growWord(nil, kept[:1], turn, e, 2, 5); len(got) != 1 || !verb.adnEnds[5] || verb.connEnds[5] {
		t.Fatal("a word that ends in a modifying closing is marked where it ends", got, verb.adnEnds, verb.connEnds)
	}
	link := wordsScan("먹을마시어", layerCost)
	eo := tail{start: 4, pieces: []piece{{text: "어", names: "맺음:어", verdict: 1, tail: rule.TailPred}}}
	e = link.stepsOf(2, 5, eo)[0]
	if got := link.growWord(nil, kept[:1], eo, e, 2, 5); len(got) != 1 || !link.connEnds[5] || link.adnEnds[5] {
		t.Fatal("a word that ends in the linking 어/아 is marked where it ends", got, link.connEnds, link.adnEnds)
	}
	if got := link.growWord(nil, nil, eo, e, 2, 4); len(got) != 0 || link.connEnds[4] {
		t.Fatal("no candidate kept before it, no mark", got, link.connEnds)
	}
	noun := wordsScan("먹을사과를", layerCost)
	noun.adnEnds[2], noun.detEnds[2] = true, true
	nt := tail{start: 4, pieces: []piece{{text: "를", names: "목적격:를", verdict: 1, tail: rule.TailCase}}}
	e = noun.stepsOf(2, 5, nt)[0]
	if got := noun.growWord(nil, kept, nt, e, 2, 5); len(got) != 3 {
		t.Fatal("a nominal grows every candidate", got)
	}
	bound := scanOf("한건")
	bound.adnEnds[1] = true
	geon := bound.fixed(2, -1)[0]
	e = bound.stepsOf(1, 2, geon)[0]
	if got := bound.growWord(nil, kept[:2], geon, e, 1, 2); len(got) != 1 || got[0].prev != kept[1] {
		t.Fatal("a tail read only after a modifying closing", got)
	}
	bin := scanOf("1010")
	e = bin.stepsOf(0, 4, tail{start: 4})[0]
	if got := bin.growWord(nil, kept[:1], tail{start: 4}, e, 0, 4); len(got) != 2 {
		t.Fatal("a number in binary has two candidates", got)
	}
}
