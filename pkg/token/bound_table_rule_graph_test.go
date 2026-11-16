//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestBoundTableRuleGraph — the graph of boundTableRule evaluated by hand on rule.Around: 리+가 stands right after a word that ends in a modifying closing (들릴, 아는), and falls after a noun, after a particle and at the head of a sentence
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundTableRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "리", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}
	r := boundTableRule(tableBound)
	for _, c := range []struct {
		name   string
		before []rule.Near
		want   bool
	}{
		{"after the modifying ㄹ", []rule.Near{{Stem: "들리", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Pred: true}}, true},
		{"after the modifying 는", []rule.Near{{Stem: "알", Kind: "열쇠:동사", Last: "꾸밈:는", Pred: true}}, true},
		{"after a noun", []rule.Near{{Stem: "회사", Kind: "열쇠:체언"}}, false},
		{"after a particle", []rule.Near{{Stem: "회사", Kind: "열쇠:체언", Last: "주격:가"}}, false},
		{"at the head", nil, false},
	} {
		if got := graphHolds(t, r, rule.Around{Self: self, Before: c.before}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
