//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestLoneShutGraph — the exceptions of a quotation mark that closes with none of its kind open, counted and judged: it stands with a letter right before it and Hangul right after it and no sign of its kind open in any way; it stands too right after a mark that ends a sentence, whatever stands after it (다.”, 다.”라고); it stands with nothing or a space before it when a marker that takes a quotation stands right after it ("라고), and with no Hangul after it when neither a letter nor a digit stands there (했다' and a space) unless a digit stands right before it (3"); it falls when one may stand open in some way, when nothing or a space stands before it, when a Latin letter or digit stands right after it, and (for ' and ’, which a row bars inside a word) between two runs of Latin letters and digits; the trace gives each reason
package parse

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
		{'"', 8, onto, true, nil},
		{'”', 8, onto, true, nil},
		{'\'', 9, onto, true, nil},
		{'’', 9, onto, true, nil},
		{'”', 8, rule.Mark{Left: rule.SideAlnum | rule.SideDigit, Right: rule.SideHangul}, true, nil},
		{'”', 8, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul, Facts: rule.MarkOpenedAny}, false, []string{whyOpenedAny}},
		{'"', 8, rule.Mark{Right: rule.SideHangul}, false, []string{whyNotOnWord}},
		{'"', 8, rule.Mark{Right: rule.SideHangul, Facts: rule.MarkQuoted | rule.MarkShutSide}, true, []string{whyNotOnWord, whyQuoteOn}},
		{'"', 8, rule.Mark{Left: rule.SideHangul}, true, []string{whyNoHangulOn, whyNoWordOn}},
		{'"', 8, rule.Mark{Left: rule.SideAlnum | rule.SideDigit}, false, []string{whyNoHangulOn, whyNoWordOn, whyAfterDigit}},
		{'’', 9, rule.Mark{Left: rule.SideHangul, Right: rule.SideAlnum}, false, []string{whyNoHangulOn}},
		{'\'', 9, rule.Mark{Left: rule.SideAlnum, Right: rule.SideAlnum | rule.SideHangul}, false, []string{whyInWord}},
		{'"', 8, rule.Mark{Left: rule.SideAlnum, Right: rule.SideAlnum | rule.SideHangul}, true, nil},
		{'"', 8, rule.Mark{Facts: rule.MarkOpenedAny | rule.MarkApart}, false, []string{whyOpenedAny, whyNotOnWord, whyNoHangulOn, whyNoWordOn}},
		{'”', 8, rule.Mark{Facts: rule.MarkAfterStop}, true, []string{whyNotOnWord, whyEndsSentence, whyNoHangulOn, whyNoWordOn}},
		{'"', 8, rule.Mark{Right: rule.SideHangul, Facts: rule.MarkAfterStop}, true, []string{whyNotOnWord, whyEndsSentence}},
		{'’', 9, rule.Mark{Facts: rule.MarkAfterStop | rule.MarkOpenedAny}, false, []string{whyOpenedAny, whyNotOnWord, whyEndsSentence, whyNoHangulOn, whyNoWordOn}},
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
