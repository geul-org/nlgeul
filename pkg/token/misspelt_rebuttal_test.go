//ff:func feature=token type=rule control=sequence
//ff:what TestMisspeltRebuttal — one edge, the rebuttal that the stem's letters are a wrong way to write another word, against a claim with no exception to lift it: a stem not misspelt stands as before and a misspelt one stands lower and above 0
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMisspeltRebuttal(t *testing.T) {
	g := toulmin.NewGraph("열쇠:x")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "열쇠:x", Concludes: "열쇠"})
	if n := misspeltRebuttal(g, claim); n != 1 {
		t.Fatal("edges", n)
	}
	judge := func(misspelt bool) float64 {
		v, _ := verdictOf(g, rule.Piece{Form: "가르", Count: seenEnough, Misspelt: misspelt}, "", false)
		return v
	}
	if v := judge(false); v != 1 {
		t.Fatal("not misspelt", v)
	}
	if v := judge(true); v <= 0 || v >= 1 {
		t.Fatal("misspelt", v)
	}
}
