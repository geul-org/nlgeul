//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestNounModifiedRuleGraph — the graph of nounModifiedRule evaluated by hand on rule.Around: a noun stands after a modifying closing, after the adnominal case particle 의 and after a determiner of the closed set (관형사:한), falls after a bare noun, after another particle, after a linking closing and at the head
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounModifiedRuleGraph(t *testing.T) {
	bogo := rule.Near{Stem: "보고", Kind: "열쇠:체언"}
	before := func(n rule.Near) rule.Around { return rule.Around{Self: bogo, Before: []rule.Near{n}} }
	r := nounModifiedRule(modifiedSkip)
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after a modifying closing", before(rule.Near{Stem: "따르", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true, Marks: "꾸밈:ㄴ"}), true},
		{"after 의", before(rule.Near{Stem: "새", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의"}), true},
		{"after a determiner", before(rule.Near{Stem: "한", Kind: "관형사:한"}), true},
		{"after a bare noun", before(rule.Near{Stem: "인구", Kind: "열쇠:체언"}), false},
		{"after another particle", before(rule.Near{Stem: "새", Kind: "열쇠:체언", Last: "목적격:를", Marks: "목적격:를"}), false},
		{"after a linking closing", before(rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:고", Pred: true, Marks: "맺음:고"}), false},
		{"at the head", rule.Around{Self: bogo}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
