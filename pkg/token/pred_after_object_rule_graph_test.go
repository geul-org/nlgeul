//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPredAfterObjectRuleGraph — the graph of predAfterObjectRule evaluated by hand on rule.Around: a predicate stands right after a word that ends in 을 or 를, and falls after a subject, after a word with an auxiliary particle, after an adverb and at the head of the sentence; a modifying form falls right after an object too (argumentAdnRule judges it)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPredAfterObjectRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "보", Kind: "열쇠:동사", Last: "맺음:니", Marks: "맺음:니", Pred: true}
	before := func(last string) []rule.Near {
		return []rule.Near{{Stem: "살림", Kind: "열쇠:체언", Last: last, Marks: last}}
	}
	r := predAfterObjectRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after 을", rule.Around{Self: self, Before: before("목적격:을")}, true},
		{"after 를", rule.Around{Self: self, Before: before("목적격:를")}, true},
		{"after a subject", rule.Around{Self: self, Before: before("주격:이")}, false},
		{"after 도", rule.Around{Self: self, Before: before("보조사:도")}, false},
		{"after an adverb", rule.Around{Self: self, Before: []rule.Near{{Stem: "다시", Kind: "열쇠:부사"}}}, false},
		{"at the head", rule.Around{Self: self}, false},
		{"a modifying form after 을", rule.Around{Self: rule.Near{Stem: "뜨", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Marks: "꾸밈:ㄹ", Pred: true}, Before: before("목적격:을")}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
