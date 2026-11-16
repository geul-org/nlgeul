//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNominalBoundRuleGraph — the graph of nominalBoundRule evaluated by hand on rule.Around: 측+은 right after a word that ends in a nominal stands, written apart or onto it, and it falls after a modifying form and at the head of a sentence
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNominalBoundRuleGraph(t *testing.T) {
	r := nominalBoundRule(nominalBound)
	self := rule.Near{Stem: "측", Kind: "열쇠:체언", Last: "보조사:은", Marks: "보조사:은"}
	glued := self
	glued.Glued = true
	noun := rule.Near{Stem: "회사", Kind: "열쇠:체언", Nominal: true}
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"apart after a nominal", rule.Around{Self: self, Before: []rule.Near{noun}}, true},
		{"onto a nominal", rule.Around{Self: glued, Before: []rule.Near{noun}}, true},
		{"after a modifying form", rule.Around{Self: self, Before: []rule.Near{{Stem: "말하", Kind: "열쇠:동사", Last: "꾸밈:는", Pred: true}}}, false},
		{"at the head", rule.Around{Self: self}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
