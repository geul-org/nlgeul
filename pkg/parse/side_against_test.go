//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestSideAgainst — the facts that speak against a candidate of a quotation mark: a letter before or a quotation's marker after the opening of " and ', a closing mark's place for the opening of “ and ‘, an opening mark's place for the closing of every quotation mark; nothing for a bracket and for an opening that is not such a mark's
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestSideAgainst(t *testing.T) {
	for _, c := range []struct {
		sign rune
		name string
		want int
	}{
		{'"', openName, rule.MarkAfterWord | rule.MarkQuoted},
		{'\'', openName, rule.MarkAfterWord | rule.MarkQuoted},
		{'“', openName, rule.MarkShutSide},
		{'‘', openName, rule.MarkShutSide},
		{'"', shutName, rule.MarkOpenSide},
		{'\'', shutName, rule.MarkOpenSide},
		{'”', shutName, rule.MarkOpenSide},
		{'’', shutName, rule.MarkOpenSide},
		{'(', openName, 0},
		{')', shutName, 0},
		{'「', openName, 0},
		{'」', shutName, 0},
	} {
		if got := sideAgainst(c.sign, pairRows[c.sign], c.name); got != c.want {
			t.Errorf("%c %s: %b, want %b", c.sign, c.name, got, c.want)
		}
	}
}
