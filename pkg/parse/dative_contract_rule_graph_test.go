//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestDativeContractRuleGraph — the graph of dativeContractRule evaluated by hand on rule.Around: 내게 stands before a verb other than 하다, 되다 and 만들다, before a noun and at the end of the sentence, and falls right before 하다, 되다 or 만들다 (-게 하다)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDativeContractRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "내", Kind: "열쇠:체언", Last: "부사격:에게", Marks: "부사격:에게"}
	verb := func(stem string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: "열쇠:동사", Last: "맺음:다", Marks: "맺음:다", Pred: true}}
	}
	r := dativeContractRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before 주다", rule.Around{Self: self, After: verb("주")}, true},
		{"before a noun", rule.Around{Self: self, After: []rule.Near{{Stem: "책", Kind: "열쇠:체언", Last: "목적격:을", Marks: "목적격:을"}}}, true},
		{"at the end", rule.Around{Self: self}, true},
		{"before 하다", rule.Around{Self: self, After: verb("하")}, false},
		{"before 되다", rule.Around{Self: self, After: verb("되")}, false},
		{"before 만들다", rule.Around{Self: self, After: verb("만들")}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
