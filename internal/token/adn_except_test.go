//ff:func feature=token type=rule control=sequence
//ff:what TestAdnExcept — three edges, the exception of a modifying closing before the stem and the two against it: on a graph with them a stem stands with no modifying closing before it, does not stand right after one (with that exception's reason), and stands there when it is a listed auxiliary (듯하) or when its own word is a modifying or a naming form
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
	if n := adnExcept(g, claim); n != 3 {
		t.Fatal("edges", n)
	}
	free, none := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred}, "", true)
	barred, why := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred, Adn: true}, "", true)
	if free <= 0 || len(none) != 0 {
		t.Fatal("no modifying closing before the stem", free, none)
	}
	if barred > 0 || !slices.Contains(why, whyAdn) {
		t.Fatal("right after a modifying closing", barred, why)
	}
	aux, _ := verdictOf(g, rule.Piece{Form: "듯하", Tail: rule.TailPred, Adn: true}, "", false)
	plain, _ := verdictOf(g, rule.Piece{Form: "듯하", Tail: rule.TailPred}, "", false)
	if aux <= 0 || aux != plain || aux != free {
		t.Fatal("a listed auxiliary right after a modifying closing", aux, plain, free)
	}
	part, _ := verdictOf(g, rule.Piece{Form: "듯", Tail: rule.TailPred, Adn: true}, "", false)
	if part > 0 {
		t.Fatal("only the listed forms are let through", part)
	}
	if turned, _ := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred, Adn: true, Turned: true}, "", false); turned != free {
		t.Fatal("a word that is itself a modifying or a naming form right after a modifying closing", turned)
	}
}
