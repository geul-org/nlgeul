//ff:func feature=token type=rule control=sequence
//ff:what TestNameGraph — a sign between two runs of letters stands as the name's; not when a side is no run, nor when both sides are numbers
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestNameGraph(t *testing.T) {
	e := nameGraph("경로:/")
	judge := func(m rule.Mark) float64 {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		v, _ := evaluate(e.g, ctx)
		return v
	}
	digits := rule.SideAlnum | rule.SideDigit
	if judge(rule.Mark{Left: rule.SideAlnum, Right: digits}) <= 0 || judge(rule.Mark{Left: rule.SideAlnum, Right: rule.SideHangul}) > 0 || judge(rule.Mark{Left: digits, Right: digits}) > 0 || e.name != "경로:/" {
		t.Fatal("name graph")
	}
}
