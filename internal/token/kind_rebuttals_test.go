//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestKindRebuttals — four edges, the rebuttals "rarely seen" and "one syllable long" against a claim and what lifts each: on a graph with them a form of more syllables seen enough times stands fully; one syllable long it stands lower, and fully as a sole word; seen once it stands lower, and fully when it is a spoken form of the closed table (spoken) as a sole word; with both rebuttals a form stands below either one alone and above 0
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestKindRebuttals(t *testing.T) {
	g := toulmin.NewGraph("열쇠:x")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "열쇠:x", Concludes: "열쇠"})
	if n := kindRebuttals(g, claim); n != 4 {
		t.Fatal("edges", n)
	}
	judge := func(form string, count int, sole bool) float64 {
		v, _ := verdictOf(g, rule.Piece{Form: form, Count: count, Sole: sole}, "", false)
		return v
	}
	if v := judge("가르", seenEnough, false); v != 1 {
		t.Fatal("a form of more syllables seen enough", v)
	}
	short := judge("솥", seenEnough, false)
	if short >= 1 || judge("솥", seenEnough, true) != 1 {
		t.Fatal("a one-syllable form seen enough", short)
	}
	few := judge("가르", 1, false)
	if few >= 1 {
		t.Fatal("a form of more syllables seen once", few)
	}
	for _, form := range spoken {
		if v := judge(form, 1, true); v != 1 {
			t.Fatal(form, "a spoken form of the table seen once as a sole word", v)
		}
	}
	both := judge("솥", 1, false)
	if both <= 0 || both >= short || both >= few {
		t.Fatal("a one-syllable form seen once", both, short, few)
	}
}
