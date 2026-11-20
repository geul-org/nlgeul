//ff:func feature=parse type=rule control=sequence
//ff:what TestFrontGraph — a sign at a number's front stands with nothing left of it and falls after letters (what stands right of it is not its graph's to ask: it is judged only where a number begins); the range sign stands after letters before digits and falls before no digits
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestFrontGraph(t *testing.T) {
	digits := &rule.Bits{Of: "sides", Mask: rule.SideDigit}
	run := func(e *entry, m rule.Mark) float64 {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		v, _ := evaluate(e.g, ctx, false)
		return v
	}
	front, any := frontGraph("부호:-", digits, true), frontGraph("범위:~", digits, false)
	num := rule.SideAlnum | rule.SideDigit
	if run(front, rule.Mark{Right: num}) <= 0 || run(front, rule.Mark{Left: rule.SideHangul, Right: num}) > 0 || run(front, rule.Mark{Right: rule.SideHangul}) <= 0 {
		t.Fatal("front")
	}
	if run(any, rule.Mark{Left: rule.SideHangul, Right: num}) <= 0 || run(any, rule.Mark{Left: rule.SideHangul}) > 0 {
		t.Fatal("any")
	}
}
