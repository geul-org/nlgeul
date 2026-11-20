//ff:func feature=parse type=rule control=sequence
//ff:what TestOrderGraph — the graph of 제 before a number is named 차례:제 and stands before digits with nothing left of it; letters left of it or no digits right of it defeat it, each with its own reason
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestOrderGraph(t *testing.T) {
	e := orderGraph(&rule.Bits{Of: "sides", Mask: rule.SideDigit})
	run := func(m rule.Mark) (float64, []string) {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		return evaluate(e.g, ctx, true)
	}
	num := rule.SideAlnum | rule.SideDigit
	if e.name != "차례:제" {
		t.Fatal(e.name)
	}
	if v, why := run(rule.Mark{Right: num, Run: 1}); v <= 0 || len(why) != 0 {
		t.Fatal("at a word's head", v, why)
	}
	if v, why := run(rule.Mark{Left: rule.SideHangul, Right: num, Run: 1}); v > 0 || !slices.Equal(why, []string{whyOrderLeft}) {
		t.Fatal("after a syllable", v, why)
	}
	if v, why := run(rule.Mark{Left: rule.SideAlnum, Right: num, Run: 1}); v > 0 || !slices.Equal(why, []string{whyOrderLeft}) {
		t.Fatal("after a Latin letter", v, why)
	}
	if v, why := run(rule.Mark{Right: rule.SideHangul, Run: 1}); v > 0 || !slices.Equal(why, []string{whyOrderRight}) {
		t.Fatal("no digits right", v, why)
	}
}
