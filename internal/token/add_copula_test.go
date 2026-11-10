//ff:func feature=token type=rule control=sequence
//ff:what TestAddCopula — the copula stem stands after a stem before a closing it takes, not alone and not before a closing it does not take; the dropped stem stands only after a vowel
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestAddCopula(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	graphs, edges, total := ix.graphs, ix.edges, ix.total
	ix.addCopula()
	said := func(p rule.Piece) float64 { v, _ := verdictOf(ix.copula.g, p, "x"); return v }
	gone := func(p rule.Piece) float64 { v, _ := verdictOf(ix.dropped.g, p, "x"); return v }
	if ix.graphs != graphs+2 || ix.edges != edges+6 || ix.total != total || ix.copula.count != copulaCount || ix.dropped.form != "" {
		t.Fatal(ix.graphs, ix.edges)
	}
	if said(rule.Piece{HasStem: true, Prev: '곰', After: "맺음:다"}) <= 0 || said(rule.Piece{Prev: '곰', After: "맺음:다"}) > 0 || said(rule.Piece{HasStem: true, Prev: '곰', After: "꾸밈:는"}) > 0 {
		t.Fatal("said")
	}
	if gone(rule.Piece{HasStem: true, Prev: '교', After: "맺음:다"}) <= 0 || gone(rule.Piece{HasStem: true, Prev: '곰', After: "맺음:다"}) > 0 {
		t.Fatal("dropped")
	}
	if adn := gone(rule.Piece{HasStem: true, Prev: '교', After: "꾸밈:ㄴ"}); adn <= 0 || adn >= gone(rule.Piece{HasStem: true, Prev: '교', After: "맺음:다"}) {
		t.Fatal("a dropped stem before a modifying closing stands lower", adn)
	}
}
