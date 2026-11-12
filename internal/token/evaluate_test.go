//ff:func feature=token type=rule control=sequence
//ff:what TestEvaluate — a claim that stands gives its verdict; one that is defeated gives the reasons of the exceptions that applied when a trace is kept, and the same verdict with no reasons when not; a graph that fails or claims nothing gives -1
package token

import (
	"testing"

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
	bad := toulmin.NewGraph("bad")
	bad.Rule(func(toulmin.Context, toulmin.Specs) (bool, any) { panic("x") })
	if v, _ := evaluate(bad, ctx, false); v != -1 {
		t.Fatal(v)
	}
	if v, _ := evaluate(toulmin.NewGraph("empty"), ctx, false); v != -1 {
		t.Fatal(v)
	}
}
