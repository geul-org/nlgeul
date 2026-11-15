//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNounListedRuleGraph — the graph of nounListedRule evaluated by hand on rule.Around: the noun 시 stands right before a middle dot, falls before a word, before another sign and at the end
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounListedRuleGraph(t *testing.T) {
	si := rule.Near{Stem: "시", Kind: "열쇠:체언"}
	r := nounListedRule(boundNext)
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before a middle dot", rule.Around{Self: si, After: []rule.Near{{Last: middleDot, Marks: middleDot}, {Stem: "군", Kind: "열쇠:체언"}}}, true},
		{"before a word", rule.Around{Self: si, After: []rule.Near{{Stem: "군", Kind: "열쇠:체언"}}}, false},
		{"before a comma", rule.Around{Self: si, After: []rule.Near{{Last: "쉼표", Marks: "쉼표"}}}, false},
		{"at the end", rule.Around{Self: si}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
