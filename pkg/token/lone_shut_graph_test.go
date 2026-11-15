//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestLoneShutGraph — the exceptions of a quotation mark that closes with none of its kind open, counted and judged: it stands with a letter right before it and Hangul right after it and no sign of its kind open in any way; it falls when one may stand open in some way, when nothing or a space stands before it, when what stands after it is no Hangul, and (for ' and ’, which a row bars inside a word) between two runs of Latin letters and digits; the trace gives each reason
package token

import (
	"reflect"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestLoneShutGraph(t *testing.T) {
	onto := rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul, Run: 1}
	cases := []struct {
		sign  rune
		edges int
		m     rule.Mark
		stand bool
		why   []string
	}{
		{'"', 3, onto, true, nil},
		{'”', 3, onto, true, nil},
		{'\'', 4, onto, true, nil},
		{'’', 4, onto, true, nil},
		{'”', 3, rule.Mark{Left: rule.SideAlnum | rule.SideDigit, Right: rule.SideHangul}, true, nil},
		{'”', 3, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul, Facts: rule.MarkOpenedAny}, false, []string{whyOpenedAny}},
		{'"', 3, rule.Mark{Right: rule.SideHangul}, false, []string{whyNotOnWord}},
		{'"', 3, rule.Mark{Left: rule.SideHangul}, false, []string{whyNoHangulOn}},
		{'’', 4, rule.Mark{Left: rule.SideHangul, Right: rule.SideAlnum}, false, []string{whyNoHangulOn}},
		{'\'', 4, rule.Mark{Left: rule.SideAlnum, Right: rule.SideAlnum | rule.SideHangul}, false, []string{whyInWord}},
		{'"', 3, rule.Mark{Left: rule.SideAlnum, Right: rule.SideAlnum | rule.SideHangul}, true, nil},
		{'"', 3, rule.Mark{Facts: rule.MarkOpenedAny | rule.MarkApart}, false, []string{whyOpenedAny, whyNotOnWord, whyNoHangulOn}},
	}
	for _, c := range cases {
		g, edges := loneShutGraph(c.sign, pairRows[c.sign])
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		v, why := evaluate(g, ctx, true)
		if edges != c.edges || (v > 0) != c.stand || !reflect.DeepEqual(why, c.why) {
			t.Errorf("%c %+v: %d edges, verdict %.2f, why %q; want %d edges, stands %v, why %q", c.sign, c.m, edges, v, why, c.edges, c.stand, c.why)
		}
	}
}
