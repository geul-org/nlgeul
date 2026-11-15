//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestUnitWordRebuttal — the rebuttal is one edge; on a graph of a kind of stem it lowers the verdict of a stem told UnitWord to about a third, above 0 (the word stays a candidate behind the unit's reading), and leaves the verdict of any other stem as it was
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestUnitWordRebuttal(t *testing.T) {
	g := toulmin.NewGraph("시험")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "시험", Concludes: "열쇠"})
	if n := unitWordRebuttal(g, claim); n != 1 {
		t.Fatal(n)
	}
	plain, _ := verdictOf(g, rule.Piece{Form: "년간"}, "", false)
	rebutted, _ := verdictOf(g, rule.Piece{Form: "년간", UnitWord: true}, "", false)
	if plain != 1 || rebutted <= 0.3 || rebutted >= 0.34 {
		t.Fatal(plain, rebutted)
	}
	ix := &index{}
	ix.addKinds()
	for _, kind := range kindOrder {
		p := rule.Piece{Form: "가르", Tail: rule.TailBare, Count: seenEnough}
		free, _ := verdictOf(ix.kinds[kind], p, "", false)
		p.UnitWord = true
		lowered, _ := verdictOf(ix.kinds[kind], p, "", false)
		if free > 0 && (lowered <= 0 || lowered >= free) {
			t.Error(kindLabel[kind], free, lowered)
		}
	}
}
