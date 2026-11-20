//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestStoppedGraph — the graph of a date of two numbers stands for runs of its shape with a full stop right after them (1.28., 1961.12. by the second shape), and not without the full stop (1.28), nor for runs of another shape (13.28., 1.32., 1961.13., three numbers)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestStoppedGraph(t *testing.T) {
	md, ym := stoppedGraph(stopShapes[0]), stoppedGraph(stopShapes[1])
	run := func(e *entry, facts int, parts ...string) bool {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, rule.Mark{Parts: parts, Facts: facts})
		v, _ := evaluate(e.g, ctx, false)
		return v > 0
	}
	if md.name != "날짜:." || ym.name != "날짜:." {
		t.Fatal(md.name, ym.name)
	}
	for _, c := range []struct {
		e     *entry
		facts int
		parts []string
		want  bool
	}{
		{md, rule.MarkStopped, []string{"1", "28"}, true},
		{md, rule.MarkStopped, []string{"12", "1"}, true},
		{md, 0, []string{"1", "28"}, false},
		{md, rule.MarkStopped, []string{"13", "28"}, false},
		{md, rule.MarkStopped, []string{"1", "32"}, false},
		{md, rule.MarkStopped, []string{"1", "2", "3"}, false},
		{ym, rule.MarkStopped, []string{"1961", "12"}, true},
		{ym, 0, []string{"1961", "12"}, false},
		{ym, rule.MarkStopped, []string{"1961", "13"}, false},
		{ym, rule.MarkStopped, []string{"61", "12"}, false},
	} {
		if got := run(c.e, c.facts, c.parts...); got != c.want {
			t.Errorf("%v (facts %d): %v, want %v", c.parts, c.facts, got, c.want)
		}
	}
}
