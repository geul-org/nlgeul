//ff:func feature=token type=rule control=sequence
//ff:what TestAddFused — the honorific under a fused letter stands after a vowel, the link vowel after a consonant, 잖 after a stem; a fused form with a fixed word is judged as the particle under it is in a chain
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddFused(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	ix.addClosings()
	ix.addSlots()
	graphs, edges, total := ix.graphs, ix.edges, ix.total
	ix.addFused()
	ok := func(e *entry, p rule.Piece) bool { v, _ := verdictOf(e.g, p, "x"); return v > 0 }
	if ix.graphs != graphs+2+len(contracts) || ix.edges != edges+4+len(contracts)+3 || ix.total != total || len(ix.fixed) != len(contracts) {
		t.Fatal(ix.graphs-graphs, ix.edges-edges)
	}
	if !ok(ix.honor, rule.Piece{HasStem: true, Prev: '가', Next: -1}) || ok(ix.honor, rule.Piece{HasStem: true, Prev: '읽', Next: -1}) || ok(ix.honor, rule.Piece{Prev: '가', Next: -1}) {
		t.Fatal("honor")
	}
	if !ok(ix.eu, rule.Piece{HasStem: true, Prev: '읽', Next: -1, After: honorName}) || ok(ix.eu, rule.Piece{HasStem: true, Prev: '가', Next: -1, After: honorName}) || ok(ix.eu, rule.Piece{HasStem: true, Prev: '살', Next: -1, After: honorName}) {
		t.Fatal("eu")
	}
	if !ok(ix.eu, rule.Piece{HasStem: true, Prev: '읽', Next: -1, After: "맺음:니"}) || ok(ix.eu, rule.Piece{HasStem: true, Prev: '읽', Next: -1, After: "맺음:고"}) {
		t.Fatal("eu stands only before what takes it as a piece")
	}
	if !ok(ix.neg, rule.Piece{HasStem: true, Next: -1}) || ok(ix.neg, rule.Piece{Next: -1}) || ix.ji.g == nil {
		t.Fatal("neg")
	}
	lexed := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}, lexed: true}
	lexed.addParticles()
	lexed.addClosings()
	lexed.addSlots()
	lexed.addFused()
	if len(lexed.fixed) != len(contracts)+len(shortPronouns) || len(lexed.contracts) != len(lexed.fixed) || len(ix.contracts) != len(contracts) || lexed.fixed[len(contracts)].name != "줄임:난" {
		t.Fatal("with a word table the one-syllable pronouns are read too", len(lexed.fixed))
	}
	first := ix.fixed[0]
	if first.name != "줄임:이게" || first.count == 0 || !ok(first, rule.Piece{HasStem: true, Slot: rule.LinkCase, Next: -1}) || ok(first, rule.Piece{HasStem: true, Slot: rule.LinkCase, Next: rule.LinkCase}) {
		t.Fatal(first)
	}
}
