//ff:func feature=token type=rule control=sequence
//ff:what TestAddFused — the honorific under a fused letter stands after a vowel, the link vowel after a consonant, 잖 after a stem; a fused form with a fixed word is judged as the particle under it is in a chain, and is counted as that particle (내게 as 에게, 어디다 as 에다, 저기다가 as 에다가)
package token

import (
	"slices"
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
	ok := func(e *entry, p rule.Piece) bool { v, _ := verdictOf(e.g, p, "x", false); return v > 0 }
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
	fixedOf := func(form string) *entry {
		return ix.fixed[slices.IndexFunc(ix.contracts, func(c contract) bool { return c.form == form })]
	}
	to, at, atOn := ix.named('게', "부사격:에게"), ix.named('다', "부사격:에다"), ix.named('가', "부사격:에다가")
	if to.count == 0 || at.count == 0 || atOn.count == 0 {
		t.Fatal("the particles under 내게, 어디다 and 어디다가 are in the index", to.count, at.count, atOn.count)
	}
	mine, where, there := fixedOf("내게"), fixedOf("어디다"), fixedOf("저기다가")
	if mine.name != "줄임:내게" || mine.count != to.count || mine.slot != to.slot || where.count != at.count || there.name != "줄임:저기다가" || there.count != atOn.count {
		t.Fatal("a fused form of a pronoun is counted as the particle under it", mine, where, there)
	}
	if !ok(mine, rule.Piece{HasStem: true, Slot: rule.LinkAdv, Next: rule.LinkTopic}) || ok(mine, rule.Piece{HasStem: true, Slot: rule.LinkAdv, Next: rule.LinkCase}) || !ok(where, rule.Piece{HasStem: true, Slot: rule.LinkAdv, Next: -1}) {
		t.Fatal("내게 before 도 but not before a case particle; 어디다 alone")
	}
	first := ix.fixed[0]
	if first.name != "줄임:이게" || first.count == 0 || !ok(first, rule.Piece{HasStem: true, Slot: rule.LinkCase, Next: -1}) || ok(first, rule.Piece{HasStem: true, Slot: rule.LinkCase, Next: rule.LinkCase}) {
		t.Fatal(first)
	}
}
