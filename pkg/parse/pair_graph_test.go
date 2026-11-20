//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPairGraph — the exceptions of one candidate of a sign of a pair, counted and judged: a closing candidate falls when no sign of its kind stands open; the opening candidate of a quotation mark falls when one stands open, and right after a mark that ends a sentence unless a word is written right after it; a sign that bars it falls inside a word, apart with a space on both sides, or next to a hyphen or an equals sign; a plain bracket has no exception on the opening side; a quotation mark right after a digit stands
package parse

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
		{'"', openName, 3, rule.Mark{}, true},
		{'"', openName, 3, rule.Mark{Facts: rule.MarkOpened}, false},
		{'"', shutName, 1, rule.Mark{Facts: rule.MarkOpened}, true},
		{'"', shutName, 1, rule.Mark{}, false},
		{'"', openName, 3, digit, true},
		{'"', openName, 3, rule.Mark{Facts: rule.MarkAfterStop}, false},
		{'"', openName, 3, rule.Mark{Right: rule.SideHangul, Facts: rule.MarkAfterStop}, true},
		{'\'', openName, 4, rule.Mark{Right: rule.SideAlnum, Facts: rule.MarkAfterStop}, true},
		{'\'', openName, 4, word, false},
		{'\'', openName, 4, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul}, true},
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
}
