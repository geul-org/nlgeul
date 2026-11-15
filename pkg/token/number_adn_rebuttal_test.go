//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNumberAdnRebuttal — the rebuttal and what lifts it are two edges; on a graph it lowers the claim of a piece right after a number before the modifying ㄴ to about a third (above 0), and leaves it as it was where no number stands before it or another closing follows (맺음:다, 꾸밈:ㄹ, 맺음:므로)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestNumberAdnRebuttal(t *testing.T) {
	g := toulmin.NewGraph("시험")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "시험", Concludes: "표지"})
	if n := numberAdnRebuttal(g, claim); n != 2 {
		t.Fatal(n)
	}
	at := func(p rule.Piece) float64 { v, _ := verdictOf(g, p, "1인", false); return v }
	if v := at(rule.Piece{Form: "이", HasStem: true, Prev: '1', Num: true, After: "꾸밈:ㄴ"}); v <= 0.3 || v >= 0.34 {
		t.Fatal("right after a number before ㄴ", v)
	}
	if v := at(rule.Piece{Form: "이", HasStem: true, Prev: '%', After: "꾸밈:ㄴ"}); v != 1 {
		t.Fatal("no number right before it", v)
	}
	for _, after := range []string{"맺음:다", "꾸밈:ㄹ", "맺음:므로"} {
		if v := at(rule.Piece{Form: "이", HasStem: true, Prev: '1', Num: true, After: after}); v != 1 {
			t.Error(after, v)
		}
	}
}
