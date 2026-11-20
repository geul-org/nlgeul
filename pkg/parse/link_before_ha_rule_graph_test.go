//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestLinkBeforeHaRuleGraph — the graph of linkBeforeHaRule evaluated by hand on rule.Around: the -어야 form stands right before 하다 and 되다, and falls before another verb, before a noun and at the end of the sentence
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLinkBeforeHaRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "내", Kind: "열쇠:동사", Last: "맺음:어야(합침)", Marks: "맺음:어야(합침)", Pred: true}
	next := func(stem, kind string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: kind, Last: "맺음:ㄴ다", Marks: "맺음:ㄴ다", Pred: true}}
	}
	r := linkBeforeHaRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before 하다", rule.Around{Self: self, After: next("하", "열쇠:동사")}, true},
		{"before 되다", rule.Around{Self: self, After: next("되", "열쇠:동사")}, true},
		{"before 가다", rule.Around{Self: self, After: next("가", "열쇠:동사")}, false},
		{"before a noun", rule.Around{Self: self, After: []rule.Near{{Stem: "집", Kind: "열쇠:체언"}}}, false},
		{"at the end", rule.Around{Self: self}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
