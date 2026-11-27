//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestFormGraph — the graph of a name of a sign of the tables of special letters by its place: a touched name does not stand in a run of signs that draw; 가림 is lowered where the sign stands alone, one ○, and stands whole at a unit's edge or in a row; 글머리 stands only as one sign at an input line's start with a space after it; any other name stands anywhere; the numbers of edges returned
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestFormGraph(t *testing.T) {
	low := 2/(1+forceNoHide) - 1
	head := rule.MarkHeads | rule.MarkSpaced
	cases := []struct {
		name  string
		drawn bool
		edges int
		m     rule.Mark
		want  float64 // 0: it does not stand
	}{
		{"기호:원", false, 0, rule.Mark{Run: 1, Facts: rule.MarkTouched}, 1},
		{"기호:원", true, 1, rule.Mark{Run: 1, Facts: rule.MarkTouched}, 0},
		{"기호:원", true, 1, rule.Mark{Run: 1}, 1},
		{"가림:원", false, 1, rule.Mark{Run: 1}, low},
		{"가림:원", false, 1, rule.Mark{Run: 2}, 1},
		{"가림:원", false, 1, rule.Mark{Run: 1, Facts: rule.MarkEnds}, 1},
		{"가림:원", true, 2, rule.Mark{Run: 1, Facts: rule.MarkTouched | rule.MarkRepeated}, 0},
		{"글머리:검은 원", false, 3, rule.Mark{Run: 1, Facts: head}, 1},
		{"글머리:검은 원", false, 3, rule.Mark{Run: 2, Facts: head | rule.MarkRepeated}, 0},
		{"글머리:검은 원", false, 3, rule.Mark{Run: 1, Facts: rule.MarkSpaced}, 0},
		{"글머리:검은 원", false, 3, rule.Mark{Run: 1, Facts: rule.MarkHeads}, 0},
		{"글머리:검은 원", true, 4, rule.Mark{Run: 1, Facts: head | rule.MarkTouched}, 0},
	}
	for _, c := range cases {
		g, edges := formGraph(c.name, c.drawn)
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		v, _ := evaluate(g, ctx, false)
		if edges != c.edges || (c.want == 0 && v > 0) || (c.want > 0 && (v < c.want-1e-9 || v > c.want+1e-9)) {
			t.Errorf("%s drawn=%v %+v: %d edges, verdict %.2f, want %d edges, verdict %.2f", c.name, c.drawn, c.m, edges, v, c.edges, c.want)
		}
	}
}
