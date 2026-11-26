//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestSwapGraph — a round quotation mark read against its shape: one drawn as a closing mark (” ’) opens only where an opening mark is written (MarkOpenSide), one drawn as an opening mark (“ ‘) closes only where a closing mark is written (MarkShutSide); each graph has one exception and its trace gives the reason
package parse

import (
	"reflect"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSwapGraph(t *testing.T) {
	for _, c := range []struct {
		sign  rune
		name  string
		facts int
		stand bool
		why   []string
	}{
		{'”', openName, rule.MarkOpenSide, true, nil},
		{'’', openName, rule.MarkOpenSide, true, nil},
		{'”', openName, rule.MarkShutSide, false, []string{whyNotOpenSide}},
		{'’', openName, 0, false, []string{whyNotOpenSide}},
		{'“', shutName, rule.MarkShutSide | rule.MarkQuoted, true, nil},
		{'‘', shutName, rule.MarkShutSide, true, nil},
		{'“', shutName, rule.MarkOpenSide, false, []string{whyNotShutSide}},
		{'‘', shutName, 0, false, []string{whyNotShutSide}},
	} {
		g, name, edges := swapGraph(c.sign)
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, rule.Mark{Run: 1, Facts: c.facts})
		v, why := evaluate(g, ctx, true)
		if name != c.name || edges != 1 || (v > 0) != c.stand || !reflect.DeepEqual(why, c.why) {
			t.Errorf("%c with %b: %s, %d edges, verdict %.2f, why %q", c.sign, c.facts, name, edges, v, why)
		}
	}
}
