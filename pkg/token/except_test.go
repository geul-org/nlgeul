//ff:func feature=token type=rule control=sequence
//ff:what TestExcept — two exceptions of one rule function put on a graph each defeat the claim where their own criteria apply, and a trace gives the reason each was put with (that of the node, not of the function), in the order they were put; where neither applies the claim stands with no reason; on the node the criteria keep their order (the first is the one the rule function reads, fact before not) and a Claim of the reason that concludes an exception comes last
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestExcept(t *testing.T) {
	g := toulmin.NewGraph("맺음:다")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "맺음:다", Concludes: "표지"})
	except(g, rule.PieceIs, whyNoStem, fact(rule.PieceStem), not).Attacks(claim)
	except(g, rule.PieceIs, whyLast, fact(rule.PieceLast)).Attacks(claim)
	if v, why := verdictOf(g, rule.Piece{Form: "다", HasStem: true}, "가다", true); v != 1 || why != nil {
		t.Fatal("a stem before it and a piece after it", v, why)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "다"}, "다", true); v != -1 || len(why) != 1 || why[0] != whyNoStem {
		t.Fatal("no stem before it", v, why)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "다", HasStem: true, Next: -1}, "가다", true); v != -1 || len(why) != 1 || why[0] != whyLast {
		t.Fatal("nothing after it", v, why)
	}
	if v, why := verdictOf(g, rule.Piece{Form: "다", Next: -1}, "다", true); v != -1 || len(why) != 2 || why[0] != whyNoStem || why[1] != whyLast {
		t.Fatal("both apply, in the order they were put", v, why)
	}
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyPiece, rule.Piece{Form: "다"})
	results, err := g.Evaluate(ctx, toulmin.EvalOption{Trace: true})
	if err != nil || len(results) != 1 || len(results[0].Trace) != 3 {
		t.Fatal(results, err)
	}
	node := results[0].Trace[1]
	if node.Role != "except" || len(node.Specs) != 3 {
		t.Fatal(node)
	}
	b, first := node.Specs[0].(*rule.Bits)
	_, second := node.Specs[1].(*rule.Not)
	c, last := node.Specs[2].(*rule.Claim)
	if !first || b.Mask != rule.PieceStem || !second || !last || *c != (rule.Claim{Name: whyNoStem, Concludes: exceptLabel}) {
		t.Fatal(node.Specs)
	}
	if lone := results[0].Trace[2].Specs; len(lone) != 2 || whyOf(lone) != whyLast {
		t.Fatal("one criterion and the Claim", lone)
	}
}
