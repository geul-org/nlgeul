//ff:func feature=token type=rule control=sequence
//ff:what TestFrontGraph — a sign at a number's front stands before digits with nothing left of it; after letters only when it is not asked to be at the front
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestFrontGraph(t *testing.T) {
	digits := &rule.Bits{Of: "sides", Mask: rule.SideDigit}
	run := func(e *entry, m rule.Mark) float64 {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		v, _ := evaluate(e.g, ctx)
		return v
	}
	front, any := frontGraph("부호:-", digits, true), frontGraph("범위:~", digits, false)
	num := rule.SideAlnum | rule.SideDigit
	if run(front, rule.Mark{Right: num}) <= 0 || run(front, rule.Mark{Left: rule.SideHangul, Right: num}) > 0 || run(front, rule.Mark{Right: rule.SideHangul}) > 0 {
		t.Fatal("front")
	}
	if run(any, rule.Mark{Left: rule.SideHangul, Right: num}) <= 0 || run(any, rule.Mark{Left: rule.SideHangul}) > 0 {
		t.Fatal("any")
	}
}
