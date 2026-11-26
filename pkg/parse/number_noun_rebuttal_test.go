//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNumberNounRebuttal — the rebuttal and what lifts it are two edges; on a graph it lowers the claim of a piece right after a number before 면 or 군 to about a third (above 0), and leaves it as it was where no number stands before it or another closing follows (맺음:다, 맺음:라면)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestNumberNounRebuttal(t *testing.T) {
	g := toulmin.NewGraph("시험")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "시험", Concludes: "표지"})
	if n := numberNounRebuttal(g, claim); n != 2 {
		t.Fatal(n)
	}
	at := func(p rule.Piece) float64 { v, _ := verdictOf(g, p, "5면", false); return v }
	for _, after := range []string{"맺음:면", "맺음:군"} {
		if v := at(rule.Piece{Form: "이", HasStem: true, Prev: '5', Num: true, After: after}); v <= 0.3 || v >= 0.34 {
			t.Error("right after a number", after, v)
		}
		if v := at(rule.Piece{Form: "이", HasStem: true, Prev: '%', After: after}); v != 1 {
			t.Error("no number right before it", after, v)
		}
	}
	for _, after := range []string{"맺음:다", "맺음:라면", "맺음:니까"} {
		if v := at(rule.Piece{Form: "이", HasStem: true, Prev: '5', Num: true, After: after}); v != 1 {
			t.Error(after, v)
		}
	}
}
