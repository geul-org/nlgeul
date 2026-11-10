//ff:func feature=token type=rule control=sequence
//ff:what TestEvaluate — a claim that stands gives its verdict; one that is defeated gives the reasons; a graph that fails or claims nothing gives -1
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
	if v, why := evaluate(ix.stops[0].g, ctx); v != -1 || len(why) != 1 || why[0] != "양쪽이 모두 그 글자 열이다" {
		t.Fatal(v, why)
	}
	ctx.Set(rule.KeyMark, rule.Mark{Left: rule.SideHangul, Run: 1})
	if v, why := evaluate(ix.stops[0].g, ctx); v != 1 || why != nil {
		t.Fatal(v, why)
	}
	bad := toulmin.NewGraph("bad")
	bad.Rule(func(toulmin.Context, toulmin.Specs) (bool, any) { panic("x") })
	if v, _ := evaluate(bad, ctx); v != -1 {
		t.Fatal(v)
	}
	if v, _ := evaluate(toulmin.NewGraph("empty"), ctx); v != -1 {
		t.Fatal(v)
	}
}
