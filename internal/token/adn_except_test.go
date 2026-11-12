//ff:func feature=token type=rule control=sequence
//ff:what TestAdnExcept — five edges, the exception of a modifying closing before the stem with the two against it and the exception of a determiner before the stem with the one against it: on a graph with them a stem stands with no modifying closing before it, does not stand right after one (with that exception's reason), and stands there when it is a listed auxiliary (듯하) or when its own word is a modifying or a naming form; right after a determiner it does not stand (with that exception's reason and not the other's), a listed auxiliary does not stand there either, and a word that is itself a modifying or a naming form does
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
	if n := adnExcept(g, claim); n != 5 {
		t.Fatal("edges", n)
	}
	free, none := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred}, "", true)
	barred, barredWhy := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred, Adn: true}, "", true)
	if free <= 0 || len(none) != 0 {
		t.Fatal("no modifying closing before the stem", free, none)
	}
	if barred > 0 || !slices.Contains(barredWhy, whyAdn) {
		t.Fatal("right after a modifying closing", barred, barredWhy)
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
	det, why := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred, Det: true}, "", true)
	if det > 0 || !slices.Contains(why, whyDet) || slices.Contains(why, whyAdn) || slices.Contains(barredWhy, whyDet) {
		t.Fatal("right after a determiner", det, why, barredWhy)
	}
	if aux, _ := verdictOf(g, rule.Piece{Form: "듯하", Tail: rule.TailPred, Det: true}, "", false); aux > 0 {
		t.Fatal("a listed auxiliary is not written onto a determiner", aux)
	}
	if turned, _ := verdictOf(g, rule.Piece{Form: "가르", Tail: rule.TailPred, Det: true, Turned: true}, "", false); turned != free {
		t.Fatal("a word that is itself a modifying or a naming form right after a determiner", turned)
	}
}
