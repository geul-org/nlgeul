//ff:func feature=token type=rule control=sequence
//ff:what TestMeasureGraph — the graph has one edge; a sign with nothing of its kind open stands as a unit fully, one with a sign of its kind open before it about a third (above 0)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMeasureGraph(t *testing.T) {
	g, edges := measureGraph()
	if edges != 1 || g == nil {
		t.Fatal(edges)
	}
	at := func(m rule.Mark) float64 {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		v, _ := evaluate(g, ctx, false)
		return v
	}
	if v := at(rule.Mark{}); v != 1 {
		t.Fatal("nothing open", v)
	}
	if v := at(rule.Mark{Facts: rule.MarkOpened}); v <= 0.3 || v >= 0.34 {
		t.Fatal("a quote open before it", v)
	}
}
