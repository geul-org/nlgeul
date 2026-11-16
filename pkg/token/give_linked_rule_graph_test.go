//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestGiveLinkedRuleGraph — the graph of giveLinkedRule evaluated by hand on rule.Around: 달다 stands after the 어/아 form, whole or fused (도와, 먹어), and falls after the linking 고, after a noun and at the head
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGiveLinkedRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "달", Kind: "열쇠:동사", Last: "맺음:라", Pred: true, Marks: "맺음:라"}
	before := func(stem, last string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: "열쇠:동사", Last: last, Pred: true, Marks: last}}
	}
	r := giveLinkedRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after a fused 아", rule.Around{Self: self, Before: before("돕", "맺음:어(합침)")}, true},
		{"after 어", rule.Around{Self: self, Before: before("먹", "맺음:어")}, true},
		{"after 고", rule.Around{Self: self, Before: before("읽", "맺음:고")}, false},
		{"after a noun", rule.Around{Self: self, Before: []rule.Near{{Stem: "돈", Kind: "열쇠:체언"}}}, false},
		{"at the head", rule.Around{Self: self}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
