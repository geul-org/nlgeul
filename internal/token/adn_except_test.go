//ff:func feature=token type=rule control=sequence
//ff:what TestAdnExcept — two edges, the exception of a modifying closing before the stem and the one against it: on a graph with them a stem stands with no modifying closing before it, does not stand right after one, and stands there when it is a listed auxiliary (듯하)
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAdnExcept(t *testing.T) {
	g := toulmin.NewGraph("열쇠:x")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "열쇠:x", Concludes: "열쇠"})
	if n := adnExcept(g, claim); n != 2 {
		t.Fatal("edges", n)
	}
	free, none := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred}, "")
	barred, why := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred, Adn: true}, "")
	if free <= 0 || len(none) != 0 {
		t.Fatal("no modifying closing before the stem", free, none)
	}
	if barred > 0 || !slices.Contains(why, reasons["PieceAdn"]) {
		t.Fatal("right after a modifying closing", barred, why)
	}
	aux, _ := verdictOf(g, rule.Piece{Form: "듯하", Tail: rule.TailPred, Adn: true}, "")
	plain, _ := verdictOf(g, rule.Piece{Form: "듯하", Tail: rule.TailPred}, "")
	if aux <= 0 || aux != plain || aux != free {
		t.Fatal("a listed auxiliary right after a modifying closing", aux, plain, free)
	}
	part, _ := verdictOf(g, rule.Piece{Form: "듯", Tail: rule.TailPred, Adn: true}, "")
	if part > 0 {
		t.Fatal("only the listed forms are let through", part)
	}
}
