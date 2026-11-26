//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPairGraph — the exceptions of one candidate of a sign of a pair, counted and judged: a closing candidate falls when no sign of its kind stands open; the opening candidate of a quotation mark falls when one stands open, and right after a mark that ends a sentence unless a word is written right after it; a sign that bars it falls inside a word, apart with a space on both sides (but when the sign that pairs with it is written apart too: MarkMated), or next to a hyphen or an equals sign; a plain bracket and the backtick have no exception on the opening side; a quotation mark right after a digit stands; a quotation mark's candidate that stands is lowered (a verdict between 0 and 1) where its place speaks against it: the opening of " and ' onto a letter or digit before it or before a marker that takes a quotation, the opening of “ ‘ where a closing mark is written, the closing of any quotation mark where an opening mark is written; a bracket's is not
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
		low   bool
	}{
		{'(', openName, 0, word, true, false},
		{'(', openName, 0, rule.Mark{Facts: rule.MarkApart | rule.MarkOpened}, true, false},
		{')', shutName, 1, rule.Mark{}, false, false},
		{')', shutName, 1, rule.Mark{Facts: rule.MarkOpened}, true, false},
		{')', shutName, 1, rule.Mark{Facts: rule.MarkOpened | rule.MarkOpenSide}, true, false},
		{'"', openName, 4, rule.Mark{}, true, false},
		{'"', openName, 4, rule.Mark{Facts: rule.MarkOpened}, false, false},
		{'"', openName, 4, rule.Mark{Facts: rule.MarkOpened | rule.MarkOpenSide}, false, false},
		{'"', shutName, 2, rule.Mark{Facts: rule.MarkOpened}, true, false},
		{'"', shutName, 2, rule.Mark{}, false, false},
		{'"', shutName, 2, rule.Mark{Facts: rule.MarkOpened | rule.MarkOpenSide}, true, true},
		{'"', openName, 4, digit, true, true},
		{'"', openName, 4, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul}, true, true},
		{'"', openName, 4, rule.Mark{Right: rule.SideHangul, Facts: rule.MarkQuoted | rule.MarkShutSide}, true, true},
		{'"', openName, 4, rule.Mark{Right: rule.SideHangul, Facts: rule.MarkOpenSide}, true, false},
		{'"', openName, 4, rule.Mark{Facts: rule.MarkAfterStop}, false, false},
		{'"', openName, 4, rule.Mark{Right: rule.SideHangul, Facts: rule.MarkAfterStop}, true, false},
		{'\'', openName, 5, rule.Mark{Right: rule.SideAlnum, Facts: rule.MarkAfterStop}, true, false},
		{'\'', openName, 5, word, false, false},
		{'\'', openName, 5, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul}, true, true},
		{'\'', shutName, 3, digit, false, false},
		{'’', shutName, 3, word, false, false},
		{'’', shutName, 3, rule.Mark{Left: rule.SideHangul, Facts: rule.MarkOpened}, true, false},
		{'”', shutName, 2, rule.Mark{Right: rule.SideHangul, Facts: rule.MarkOpened | rule.MarkOpenSide}, true, true},
		{'“', openName, 1, rule.Mark{Right: rule.SideHangul, Facts: rule.MarkOpenSide}, true, false},
		{'“', openName, 1, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul}, true, false},
		{'“', openName, 1, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul, Facts: rule.MarkQuoted | rule.MarkShutSide}, true, true},
		{'‘', openName, 1, rule.Mark{Left: rule.SideHangul, Facts: rule.MarkShutSide}, true, true},
		{'「', openName, 0, rule.Mark{Left: rule.SideHangul, Facts: rule.MarkShutSide}, true, false},
		{'<', openName, 3, rule.Mark{Facts: rule.MarkApart}, false, false},
		{'<', openName, 3, rule.Mark{Facts: rule.MarkArrow}, false, false},
		{'<', openName, 3, rule.Mark{Right: rule.SideHangul}, true, false},
		{'>', shutName, 4, rule.Mark{Facts: rule.MarkOpened | rule.MarkArrow}, false, false},
		{'>', shutName, 4, rule.Mark{Facts: rule.MarkOpened}, true, false},
		{'<', openName, 3, rule.Mark{Facts: rule.MarkApart | rule.MarkMated}, true, false},
		{'<', openName, 3, rule.Mark{Facts: rule.MarkApart | rule.MarkMated | rule.MarkArrow}, false, false},
		{'>', shutName, 4, rule.Mark{Facts: rule.MarkOpened | rule.MarkApart}, false, false},
		{'>', shutName, 4, rule.Mark{Facts: rule.MarkOpened | rule.MarkApart | rule.MarkMated}, true, false},
		{'>', shutName, 4, rule.Mark{Facts: rule.MarkApart | rule.MarkMated}, false, false},
		{'`', openName, 0, word, true, false},
		{'`', openName, 0, rule.Mark{Facts: rule.MarkApart}, true, false},
	}
	for _, c := range cases {
		g, edges := pairGraph(c.sign, pairRows[c.sign], c.name)
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		v, _ := evaluate(g, ctx, false)
		if edges != c.edges || (v > 0) != c.stand || (v > 0 && v < 1) != c.low {
			t.Errorf("%c %s: %d edges, verdict %.2f, want %d edges, stands %v", c.sign, c.name, edges, v, c.edges, c.stand)
		}
	}
}
