//ff:func feature=parse type=rule control=sequence
//ff:what TestEvaluate — a claim that stands gives its verdict; one that is defeated gives the reasons of the exceptions that applied when a trace is kept, and the same verdict with no reasons when not; a case exception that applied gives its reason with the word it matched and its row; a graph that fails or claims nothing gives -1
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestEvaluate(t *testing.T) {
	ix := scanOf("x").ix
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyMark, rule.Mark{Left: rule.SideAlnum, Right: rule.SideAlnum, Run: 1})
	if v, why := evaluate(ix.stops[0].g, ctx, true); v != -1 || len(why) != 1 || why[0] != whyBoth {
		t.Fatal(v, why)
	}
	if v, why := evaluate(ix.stops[0].g, ctx, false); v != -1 || why != nil {
		t.Fatal("no trace kept", v, why)
	}
	ctx.Set(rule.KeyMark, rule.Mark{Left: rule.SideHangul, Run: 1})
	if v, why := evaluate(ix.stops[0].g, ctx, true); v != 1 || why != nil {
		t.Fatal(v, why)
	}
	g := toulmin.NewGraph("맺음:소")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: "맺음:소", Concludes: "표지"})
	rows := []caseException{{Word: "담소", GroundKind: groundNorm, Ground: "조항", Row: data.Row{Line: 5}}}
	except(g, rule.Listed, whyCase, listed("caseExceptions", rule.ListWord, caseWords(rows)), &caseRows{Rows: rows}).Attacks(claim)
	word := toulmin.NewContext()
	word.Set(rule.KeyPiece, rule.Piece{Form: "소"})
	word.Set(rule.KeyWord, &rule.Word{Text: "담소", Sent: &rule.Sentence{}})
	if v, why := evaluate(g, word, true); v != -1 || len(why) != 1 || why[0] != whyCase+"(담소) — 예외 표 5: 규범 조항" {
		t.Fatal(v, why)
	}
	bad := toulmin.NewGraph("bad")
	bad.Rule(func(toulmin.Context, toulmin.Specs) (bool, any) { panic("x") })
	if v, _ := evaluate(bad, ctx, false); v != -1 {
		t.Fatal(v)
	}
	if v, _ := evaluate(toulmin.NewGraph("empty"), ctx, false); v != -1 {
		t.Fatal(v)
	}
}
