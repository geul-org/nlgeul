//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPairGraph — the exceptions of one reading of a sign of a pair, counted and judged: a closing reading falls when no sign of its kind stands open; the opening reading of a quotation mark falls when one stands open; a sign that bars it falls inside a word, apart with a space on both sides, or next to a hyphen or an equals sign; a plain bracket has no exception on the opening side; a row that bars a digit before the sign adds one that falls right after a digit
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPairGraph(t *testing.T) {
	word := rule.Mark{Left: rule.SideAlnum, Right: rule.SideAlnum, Run: 1}
	digit := rule.Mark{Left: rule.SideAlnum | rule.SideDigit, Run: 1, Facts: rule.MarkAfterDigit}
	cases := []struct {
		sign  rune
		name  string
		edges int
		m     rule.Mark
		stand bool
	}{
		{'(', openName, 0, word, true},
		{'(', openName, 0, rule.Mark{Facts: rule.MarkApart | rule.MarkOpened}, true},
		{')', shutName, 1, rule.Mark{}, false},
		{')', shutName, 1, rule.Mark{Facts: rule.MarkOpened}, true},
		{'"', openName, 1, rule.Mark{}, true},
		{'"', openName, 1, rule.Mark{Facts: rule.MarkOpened}, false},
		{'"', shutName, 1, rule.Mark{Facts: rule.MarkOpened}, true},
		{'"', shutName, 1, rule.Mark{}, false},
		{'"', openName, 1, digit, true},
		{'\'', openName, 2, word, false},
		{'\'', openName, 2, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul}, true},
		{'\'', shutName, 2, digit, false},
		{'’', shutName, 2, word, false},
		{'’', shutName, 2, rule.Mark{Left: rule.SideHangul, Facts: rule.MarkOpened}, true},
		{'<', openName, 2, rule.Mark{Facts: rule.MarkApart}, false},
		{'<', openName, 2, rule.Mark{Facts: rule.MarkArrow}, false},
		{'<', openName, 2, rule.Mark{Right: rule.SideHangul}, true},
		{'>', shutName, 3, rule.Mark{Facts: rule.MarkOpened | rule.MarkArrow}, false},
		{'>', shutName, 3, rule.Mark{Facts: rule.MarkOpened}, true},
	}
	for _, c := range cases {
		g, edges := pairGraph(c.sign, pairRows[c.sign], c.name)
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		v, _ := evaluate(g, ctx, false)
		if edges != c.edges || (v > 0) != c.stand {
			t.Errorf("%c %s: %d edges, verdict %.2f, want %d edges, stands %v", c.sign, c.name, edges, v, c.edges, c.stand)
		}
	}
	g, edges := pairGraph('"', pairRow{kind: '"', opens: true, bars: barDigit}, openName)
	for facts, stand := range map[int]bool{0: true, rule.MarkAfterDigit: false} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, rule.Mark{Facts: facts})
		if v, _ := evaluate(g, ctx, false); edges != 1 || (v > 0) != stand {
			t.Errorf("a row that bars a digit before it, facts %v: %d edges, verdict %.2f", facts, edges, v)
		}
	}
}
