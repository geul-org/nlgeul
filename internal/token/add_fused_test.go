//ff:func feature=token type=rule control=sequence
//ff:what TestAddFused — the honorific under a fused letter stands after a vowel, the link vowel after a consonant, 잖 after a stem; a fused form with a fixed word is judged as the particle under it is in a chain
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestAddFused(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	ix.addClosings()
	graphs, edges, total := ix.graphs, ix.edges, ix.total
	ix.addFused()
	ok := func(e *entry, p rule.Piece) bool { v, _ := verdictOf(e.g, p, "x"); return v > 0 }
	if ix.graphs != graphs+3+len(contracts) || ix.edges != edges+5+len(contracts) || ix.total != total || len(ix.fixed) != len(contracts) {
		t.Fatal(ix.graphs-graphs, ix.edges-edges)
	}
	if !ok(ix.honor, rule.Piece{HasStem: true, Prev: '가', Next: -1}) || ok(ix.honor, rule.Piece{HasStem: true, Prev: '읽', Next: -1}) || ok(ix.honor, rule.Piece{Prev: '가', Next: -1}) {
		t.Fatal("honor")
	}
	if !ok(ix.eu, rule.Piece{HasStem: true, Prev: '읽', Next: -1}) || ok(ix.eu, rule.Piece{HasStem: true, Prev: '가', Next: -1}) || ok(ix.eu, rule.Piece{HasStem: true, Prev: '살', Next: -1}) {
		t.Fatal("eu")
	}
	if !ok(ix.neg, rule.Piece{HasStem: true, Next: -1}) || ok(ix.neg, rule.Piece{Next: -1}) || ix.ji.g == nil {
		t.Fatal("neg")
	}
	first := ix.fixed[0]
	if first.name != "줄임:이게" || first.count == 0 || !ok(first, rule.Piece{HasStem: true, Slot: rule.LinkCase, Next: -1}) || ok(first, rule.Piece{HasStem: true, Slot: rule.LinkCase, Next: rule.LinkCase}) {
		t.Fatal(first)
	}
}
