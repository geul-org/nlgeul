//ff:func feature=token type=rule control=sequence
//ff:what TestAloneExcept — two edges, the exceptions of a modifying closing and of a determiner before the word: on a graph with them a word stands with neither before it, and does not stand right after a modifying closing or right after a determiner, each for its own reason; being a modifying or a naming form itself lifts neither
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAloneExcept(t *testing.T) {
	g := toulmin.NewGraph("열쇠:x")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "열쇠:x", Concludes: "열쇠"})
	if n := aloneExcept(g, claim); n != 2 {
		t.Fatal("edges", n)
	}
	free, none := verdictOf(g, rule.Piece{Form: "아이고", Tail: rule.TailBare}, "", true)
	if free <= 0 || len(none) != 0 {
		t.Fatal("nothing that modifies a nominal before the word", free, none)
	}
	adn, why := verdictOf(g, rule.Piece{Form: "아이고", Tail: rule.TailBare, Adn: true}, "", true)
	if adn > 0 || !slices.Contains(why, whyAdn) || slices.Contains(why, whyDet) {
		t.Fatal("right after a modifying closing", adn, why)
	}
	det, why := verdictOf(g, rule.Piece{Form: "아이고", Tail: rule.TailBare, Det: true}, "", true)
	if det > 0 || !slices.Contains(why, whyDet) || slices.Contains(why, whyAdn) {
		t.Fatal("right after a determiner", det, why)
	}
	if turned, _ := verdictOf(g, rule.Piece{Form: "아이고", Tail: rule.TailBare, Det: true, Turned: true}, "", false); turned > 0 {
		t.Fatal("nothing lifts the exception", turned)
	}
}
