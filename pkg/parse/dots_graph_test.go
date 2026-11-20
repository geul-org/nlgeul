//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestDotsGraph — the graph of a run of middle dots stands for a run at a word's end, right before the letters of its unit or alone, and not for one dot, nor for a run between letters; it has two edges; end to end 그는······ ends in ······[줄임표] and 가··나 keeps its dots inside the word
package parse

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestDotsGraph(t *testing.T) {
	g, edges := dotsGraph()
	if edges != 2 {
		t.Fatal(edges)
	}
	for _, c := range []struct {
		m    rule.Mark
		want bool
	}{
		{rule.Mark{Run: 6, Facts: rule.MarkEnds}, true},
		{rule.Mark{Run: 3, Facts: rule.MarkLeads}, true},
		{rule.Mark{Run: 2, Facts: rule.MarkAlone}, true},
		{rule.Mark{Run: 1, Facts: rule.MarkEnds}, false},
		{rule.Mark{Run: 2}, false},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		if v, _ := evaluate(g, ctx, false); (v > 0) != c.want {
			t.Errorf("%+v: %v, want %v", c.m, v, c.want)
		}
	}
	r, _ := New(Options{})
	if got := streamText(r, "그는······"); !strings.Contains(got, "······[줄임표]") {
		t.Errorf("a run of middle dots at a word's end\n%s", got)
	}
	if got := streamText(r, "가··나"); strings.Contains(got, "[줄임표]") {
		t.Errorf("a run of two between letters\n%s", got)
	}
}
