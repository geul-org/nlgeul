//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestLoneBracketGraph — the exceptions of a closing bracket that closes with none of its kind open, counted and judged: it stands with a letter right before it and none of its kind open in any way, whatever stands after it (a comma, a full stop, nothing, Hangul); it falls when one of its kind may stand open in some way, when nothing or a space stands before it, when a digit stands right before it (a list number), and > (which its row bars apart and next to an arrow) with spaces on both sides or next to an arrow; the trace gives each reason
package token

import (
	"reflect"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestLoneBracketGraph(t *testing.T) {
	cases := []struct {
		sign  rune
		edges int
		m     rule.Mark
		stand bool
		why   []string
	}{
		{')', 3, rule.Mark{Left: rule.SideHangul, Run: 1}, true, nil},
		{')', 3, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul, Run: 1}, true, nil},
		{'}', 3, rule.Mark{Left: rule.SideHangul}, true, nil},
		{'』', 3, rule.Mark{Left: rule.SideAlnum}, true, nil},
		{')', 3, rule.Mark{Left: rule.SideHangul, Facts: rule.MarkOpenedAny}, false, []string{whyOpenedAny}},
		{')', 3, rule.Mark{Right: rule.SideHangul}, false, []string{whyNotOnWord}},
		{')', 3, rule.Mark{Left: rule.SideAlnum | rule.SideDigit}, false, []string{whyAfterDigit}},
		{'>', 5, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul}, true, nil},
		{'>', 5, rule.Mark{Facts: rule.MarkApart}, false, []string{whyNotOnWord, whyApart}},
		{'>', 5, rule.Mark{Left: rule.SideAlnum, Facts: rule.MarkArrow}, false, []string{whyArrow}},
	}
	for _, c := range cases {
		g, edges := loneBracketGraph(c.sign, pairRows[c.sign])
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		v, why := evaluate(g, ctx, true)
		if edges != c.edges || (v > 0) != c.stand || !reflect.DeepEqual(why, c.why) {
			t.Errorf("%c %+v: %d edges, verdict %.2f, why %q; want %d edges, stands %v, why %q", c.sign, c.m, edges, v, why, c.edges, c.stand, c.why)
		}
	}
}
