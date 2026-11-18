//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNamedBeforeHaRuleGraph — the graph of namedBeforeHaRule evaluated by hand on rule.Around: the -기 form with a particle stands right before 하다, and falls before 되다, before another verb, before a noun and at the end of the sentence
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNamedBeforeHaRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "서", Kind: "열쇠:동사", Last: "보조사:도", Marks: "이름:기+보조사:도", Pred: true}
	next := func(stem, kind string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: kind, Last: "맺음:ㄴ다", Marks: "맺음:ㄴ다", Pred: true}}
	}
	r := namedBeforeHaRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before 하다", rule.Around{Self: self, After: next("하", "열쇠:동사")}, true},
		{"before 되다", rule.Around{Self: self, After: next("되", "열쇠:동사")}, false},
		{"before 가다", rule.Around{Self: self, After: next("가", "열쇠:동사")}, false},
		{"before a noun", rule.Around{Self: self, After: []rule.Near{{Stem: "집", Kind: "열쇠:체언"}}}, false},
		{"at the end", rule.Around{Self: self}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
