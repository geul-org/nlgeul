//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestCountedShortRuleGraph — the graph of countedShortRule evaluated by hand on rule.Around: 대+가 right after the numeral determiner 한 or 두 read as the determiner of the closed set stands, and it falls after 한 read as the noun (恨), after another determiner (그), after a noun and at the head of a sentence
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCountedShortRuleGraph(t *testing.T) {
	r := countedShortRule(numDeterminers)
	self := rule.Near{Stem: "대", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가", Text: "대가"}
	for _, c := range []struct {
		name   string
		before []rule.Near
		want   bool
	}{
		{"after 한", []rule.Near{{Stem: "한", Kind: "관형사:한"}}, true},
		{"after 두", []rule.Near{{Stem: "두", Kind: "관형사:두"}}, true},
		{"after the noun 한", []rule.Near{{Stem: "한", Kind: "열쇠:체언"}}, false},
		{"after 그", []rule.Near{{Stem: "그", Kind: "관형사:그"}}, false},
		{"after a noun", []rule.Near{{Stem: "차", Kind: "열쇠:체언"}}, false},
		{"at the head", nil, false},
	} {
		if got := graphHolds(t, r, rule.Around{Self: self, Before: c.before}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
