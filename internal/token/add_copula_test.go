//ff:func feature=token type=rule control=sequence
//ff:what TestAddCopula — the copula stem stands after a stem before a closing it takes, not alone and not before a closing it does not take; the dropped stem stands after a vowel, and after a letter that is no Hangul syllable only before 야; before the naming closing 기 it stands only when a particle follows, and then lower than before 다
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddCopula(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	graphs, edges, total := ix.graphs, ix.edges, ix.total
	ix.addCopula()
	said := func(p rule.Piece) float64 { v, _ := verdictOf(ix.copula.g, p, "x"); return v }
	gone := func(p rule.Piece) float64 { v, _ := verdictOf(ix.dropped.g, p, "x"); return v }
	if ix.graphs != graphs+2 || ix.edges != edges+10 || ix.total != total || ix.copula.count != copulaCount || ix.dropped.form != "" {
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
	if gone(rule.Piece{HasStem: true, Prev: 'c', After: "맺음:야"}) <= 0 || gone(rule.Piece{HasStem: true, Prev: '3', After: "맺음:야"}) <= 0 {
		t.Fatal("a dropped stem before 야 after a letter that is no Hangul syllable")
	}
	if gone(rule.Piece{HasStem: true, Prev: 'c', After: "맺음:다"}) > 0 {
		t.Fatal("no dropped stem before 다 after a letter that is no Hangul syllable")
	}
	if gone(rule.Piece{HasStem: true, Prev: '제', After: "이름:기"}) > 0 || gone(rule.Piece{HasStem: true, Prev: '제', After: "이름:기", Right: rule.RightOther}) > 0 {
		t.Fatal("no dropped stem before 기 with no particle after it")
	}
	if gi := gone(rule.Piece{HasStem: true, Prev: '제', After: "이름:기", Right: rule.RightParticle}); gi <= 0 || gi >= da {
		t.Fatal("a dropped stem before 기 with a particle after it stands, lower than before 다", gi)
	}
}
