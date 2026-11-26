//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestDashGraph — the candidates of the hyphen and of the box-drawing line by their place: ─ is 줄표 in a row of any length, between letters too, but not in a line of signs alone nor in a run with a sign that draws; the hyphen is 줄표 alone with a space on each side or two and more in a row (between letters and at a word's edge too), not one at a line's start, not one between letters; 빼기 is one hyphen alone only; where a sum is written 빼기 stands whole and 줄표 is lowered, elsewhere the other way round
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestDashGraph(t *testing.T) {
	low := 2/(1+forceDashSum) - 1
	cases := []struct {
		sign  rune
		name  string
		edges int
		m     rule.Mark
		want  float64 // 0: it does not stand
	}{
		{'─', "줄표", 2, rule.Mark{Run: 1, Facts: rule.MarkAlone | rule.MarkApart}, 1},
		{'─', "줄표", 2, rule.Mark{Run: 2, Left: rule.SideHangul, Right: rule.SideHangul}, 1},
		{'─', "줄표", 2, rule.Mark{Run: 1, Facts: rule.MarkLeads | rule.MarkHeads}, 1},
		{'─', "줄표", 2, rule.Mark{Run: 5, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkBare}, 0},
		{'─', "줄표", 2, rule.Mark{Run: 4, Facts: rule.MarkBySign | rule.MarkDrawn}, 0},
		{'-', "줄표", 7, rule.Mark{Run: 1, Facts: rule.MarkAlone | rule.MarkApart}, 1},
		{'-', "줄표", 7, rule.Mark{Run: 1, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkSum}, low},
		{'-', "줄표", 7, rule.Mark{Run: 1, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkHeads | rule.MarkSpaced}, 0},
		{'-', "줄표", 7, rule.Mark{Run: 1, Facts: rule.MarkAlone}, 0},
		{'-', "줄표", 7, rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideHangul}, 0},
		{'-', "줄표", 7, rule.Mark{Run: 1, Right: rule.SideHangul, Facts: rule.MarkLeads}, 0},
		{'-', "줄표", 7, rule.Mark{Run: 3, Left: rule.SideHangul, Right: rule.SideHangul}, 1},
		{'-', "줄표", 7, rule.Mark{Run: 3, Left: rule.SideHangul, Facts: rule.MarkEnds}, 1},
		{'-', "줄표", 7, rule.Mark{Run: 2, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkHeads | rule.MarkSpaced}, 1},
		{'-', "줄표", 7, rule.Mark{Run: 3, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkHeads | rule.MarkBare}, 0},
		{'-', "줄표", 7, rule.Mark{Run: 2, Facts: rule.MarkAlone | rule.MarkBySign | rule.MarkDrawn}, 0},
		{'-', minusName, 6, rule.Mark{Run: 1, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkSum}, 1},
		{'-', minusName, 6, rule.Mark{Run: 1, Facts: rule.MarkAlone | rule.MarkApart}, low},
		{'-', minusName, 6, rule.Mark{Run: 1, Facts: rule.MarkAlone}, 0},
		{'-', minusName, 6, rule.Mark{Run: 2, Facts: rule.MarkAlone | rule.MarkApart}, 0},
		{'-', minusName, 6, rule.Mark{Run: 1, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkHeads}, 0},
		{'-', minusName, 6, rule.Mark{Run: 1, Facts: rule.MarkAlone | rule.MarkApart | rule.MarkBare}, 0},
		{'-', minusName, 6, rule.Mark{Run: 1, Left: rule.SideDigit, Right: rule.SideDigit}, 0},
	}
	for _, c := range cases {
		g, edges := dashGraph(c.sign, c.name)
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		v, _ := evaluate(g, ctx, false)
		if edges != c.edges || (c.want == 0 && v > 0) || (c.want > 0 && (v < c.want-1e-9 || v > c.want+1e-9)) {
			t.Errorf("%c %s %+v: %d edges, verdict %.2f, want %d edges, verdict %.2f", c.sign, c.name, c.m, edges, v, c.edges, c.want)
		}
	}
}
