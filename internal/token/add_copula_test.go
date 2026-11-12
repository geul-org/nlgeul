//ff:func feature=token type=rule control=sequence
//ff:what TestAddCopula — the copula stem stands after a stem before a closing it takes, not alone and not before a closing it does not take; the dropped stem stands after a vowel before a closing it is dropped before, and not before another closing or a slot; the closings of droppableKnown take it only with a word table: after a letter that is no Hangul syllable only before 야, and before the naming closing 기 only when a particle follows, and then lower than before 다; neither stem stands before 하 of 어하다 (copulaBars), and each is turned down for that reason
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddCopula(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	graphs, edges, total := ix.graphs, ix.edges, ix.total
	ix.addCopula()
	said := func(p rule.Piece) float64 { v, _ := verdictOf(ix.copula.g, p, "x", false); return v }
	gone := func(p rule.Piece) float64 { v, _ := verdictOf(ix.dropped.g, p, "x", false); return v }
	if ix.graphs != graphs+2 || ix.edges != edges+13 || ix.total != total || ix.copula.count != copulaCount || ix.dropped.form != "" {
		t.Fatal(ix.graphs, ix.edges)
	}
	if said(rule.Piece{HasStem: true, Prev: '곰', After: "맺음:다"}) <= 0 || said(rule.Piece{Prev: '곰', After: "맺음:다"}) > 0 || said(rule.Piece{HasStem: true, Prev: '곰', After: "꾸밈:는"}) > 0 {
		t.Fatal("said")
	}
	da := gone(rule.Piece{HasStem: true, Prev: '교', After: "맺음:다"})
	if da <= 0 || gone(rule.Piece{HasStem: true, Prev: '곰', After: "맺음:다"}) > 0 {
		t.Fatal("dropped")
	}
	if adn := gone(rule.Piece{HasStem: true, Prev: '교', After: "꾸밈:ㄴ"}); adn <= 0 || adn >= da {
		t.Fatal("a dropped stem before a modifying closing stands lower", adn)
	}
	if gone(rule.Piece{HasStem: true, Prev: '교', After: "맺음:어서"}) > 0 || gone(rule.Piece{HasStem: true, Prev: '교', After: "선어말:었"}) > 0 {
		t.Fatal("no dropped stem before a closing it is not dropped before, nor before a slot")
	}
	if gone(rule.Piece{HasStem: true, Prev: 'c', After: "맺음:야"}) > 0 || gone(rule.Piece{HasStem: true, Prev: '제', After: "이름:기", Right: rule.RightParticle}) > 0 {
		t.Fatal("without a word table no dropped stem before 야 or 기")
	}
	if said(rule.Piece{HasStem: true, Prev: '곰', After: "보조:하"}) > 0 || gone(rule.Piece{HasStem: true, Prev: '교', After: "보조:하"}) > 0 {
		t.Fatal("no copula stem, said or dropped, before an auxiliary written onto a main predicate")
	}
	if _, why := verdictOf(ix.copula.g, rule.Piece{HasStem: true, Prev: '곰', After: "보조:하"}, "x", true); !slices.Contains(why, whyHelper) {
		t.Fatal("turned down for the reason on its own node", why)
	}
	if _, why := verdictOf(ix.dropped.g, rule.Piece{HasStem: true, Prev: '교', After: "보조:하"}, "x", true); !slices.Contains(why, whyHelper) {
		t.Fatal("the dropped stem too is turned down for it", why)
	}
	lexed := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}, lexed: true}
	lexed.addClosings()
	lexed.addCopula()
	known := func(p rule.Piece) float64 { v, _ := verdictOf(lexed.dropped.g, p, "x", false); return v }
	if known(rule.Piece{HasStem: true, Prev: 'c', After: "맺음:야"}) <= 0 || known(rule.Piece{HasStem: true, Prev: '3', After: "맺음:야"}) <= 0 {
		t.Fatal("a dropped stem before 야 after a letter that is no Hangul syllable")
	}
	if known(rule.Piece{HasStem: true, Prev: 'c', After: "맺음:다"}) > 0 {
		t.Fatal("no dropped stem before 다 after a letter that is no Hangul syllable")
	}
	if known(rule.Piece{HasStem: true, Prev: '제', After: "이름:기"}) > 0 || known(rule.Piece{HasStem: true, Prev: '제', After: "이름:기", Right: rule.RightOther}) > 0 {
		t.Fatal("no dropped stem before 기 with no particle after it")
	}
	if gi := known(rule.Piece{HasStem: true, Prev: '제', After: "이름:기", Right: rule.RightParticle}); gi <= 0 || gi >= known(rule.Piece{HasStem: true, Prev: '교', After: "맺음:다"}) {
		t.Fatal("a dropped stem before 기 with a particle after it stands, lower than before 다", gi)
	}
}
