//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestGluedBoundRuleGraph — the graph of gluedBoundRule evaluated by hand on rule.Around: 자+가 written onto a word that ends in a nominal (패소자가) stands, and it falls written apart, and written onto a word that does not end in a nominal
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGluedBoundRuleGraph(t *testing.T) {
	r := gluedBoundRule(tableBound)
	glued := rule.Near{Stem: "자", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가", Glued: true}
	apart := rule.Near{Stem: "자", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}
	noun := rule.Near{Stem: "패소", Kind: "열쇠:체언", Nominal: true}
	verb := rule.Near{Stem: "받", Kind: "열쇠:동사", Last: "꾸밈:은", Pred: true}
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"onto a nominal", rule.Around{Self: glued, Before: []rule.Near{noun}}, true},
		{"apart after a nominal", rule.Around{Self: apart, Before: []rule.Near{noun}}, false},
		{"onto a predicate", rule.Around{Self: glued, Before: []rule.Near{verb}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
