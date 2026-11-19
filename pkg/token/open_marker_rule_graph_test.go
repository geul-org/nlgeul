//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestOpenMarkerRuleGraph — the graph of openMarkerRule evaluated by hand on rule.Around: markers alone are rebutted right after an opening quotation mark or bracket; not after a closing one, a nominal, nor with nothing before them
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestOpenMarkerRuleGraph(t *testing.T) {
	r := openMarkerRule()
	self := rule.Near{Last: "목적격:을", Marks: "목적격:을", Glued: true}
	for _, c := range []struct {
		name   string
		before []rule.Near
		want   bool
	}{
		{"after an opening bracket", []rule.Near{{Stem: "(", Kind: "열기:1"}}, true},
		{"after an opening quotation mark", []rule.Near{{Stem: "'", Kind: "열기:2"}}, true},
		{"after a closing bracket", []rule.Near{{Stem: ")", Kind: "닫기:1"}}, false},
		{"after a nominal", []rule.Near{{Stem: "약정", Kind: "열쇠:체언"}}, false},
		{"with nothing before", nil, false},
	} {
		if got := graphHolds(t, r, rule.Around{Self: self, Before: c.before}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
