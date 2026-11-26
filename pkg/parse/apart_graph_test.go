//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestApartGraph — a sign with no name stands apart between letters that are not both Latin letters or digits (값=3, 값=참); not between two such runs (a=1), in a row (==), next to another sign, nor at an edge of its unit; four edges
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestApartGraph(t *testing.T) {
	g, edges := apartGraph('=')
	if edges != 4 {
		t.Fatal(edges)
	}
	for _, c := range []struct {
		m     rule.Mark
		stand bool
	}{
		{rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideAlnum | rule.SideDigit}, true},
		{rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideHangul}, true},
		{rule.Mark{Run: 1, Left: rule.SideAlnum, Right: rule.SideHangul}, true},
		{rule.Mark{Run: 1, Left: rule.SideAlnum, Right: rule.SideAlnum | rule.SideDigit}, false},
		{rule.Mark{Run: 2, Left: rule.SideHangul, Right: rule.SideHangul}, false},
		{rule.Mark{Run: 1, Left: rule.SideHangul, Facts: rule.MarkBySign}, false},
		{rule.Mark{Run: 1, Left: rule.SideHangul, Facts: rule.MarkEnds}, false},
		{rule.Mark{Run: 1, Right: rule.SideHangul, Facts: rule.MarkLeads}, false},
		{rule.Mark{Run: 1, Facts: rule.MarkAlone}, false},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		if v, _ := evaluate(g, ctx, false); (v > 0) != c.stand {
			t.Errorf("%+v: verdict %.2f, want stands %v", c.m, v, c.stand)
		}
	}
}
