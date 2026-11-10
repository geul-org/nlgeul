//ff:func feature=token type=rule control=sequence
//ff:what TestShapeGraph — the graph stands for runs of its shape and not for others
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestShapeGraph(t *testing.T) {
	e := shapeGraph(rule.Shape{Of: "비율::", Min: 2, Max: 2})
	run := func(parts ...string) float64 {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, rule.Mark{Parts: parts})
		v, _ := evaluate(e.g, ctx)
		return v
	}
	if e.name != "비율::" || run("16", "9") <= 0 || run("10", "30", "00") > 0 {
		t.Fatal(e.name)
	}
}
