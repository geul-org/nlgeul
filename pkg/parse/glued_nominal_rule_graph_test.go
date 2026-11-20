//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestGluedNominalRuleGraph — the graph of gluedNominalRule evaluated by hand on rule.Around: the nominal 전 written onto a noun stands at the end of its word, with a noun written onto it, and with a predicate written apart after it; it falls written apart from the noun, written onto a word that ends in a particle, and with a predicate written onto it
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGluedNominalRuleGraph(t *testing.T) {
	glued := rule.Near{Stem: "전", Kind: "열쇠:체언", Glued: true}
	apart := rule.Near{Stem: "전", Kind: "열쇠:체언"}
	noun := rule.Near{Stem: "발병", Kind: "열쇠:체언", Nominal: true}
	pred := rule.Near{Stem: "하", Kind: "열쇠:동사", Last: "맺음:고", Marks: "맺음:고", Pred: true}
	predOnto := rule.Near{Stem: "하", Kind: "열쇠:동사", Last: "맺음:고", Marks: "맺음:고", Pred: true, Glued: true}
	r := gluedNominalRule(map[string]word{"전": {kinds: kindNominal | kindDeterminer}}, map[string]int{})
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"onto a noun", rule.Around{Self: glued, Before: []rule.Near{noun}}, true},
		{"a noun written onto it", rule.Around{Self: glued, Before: []rule.Near{noun}, After: []rule.Near{{Stem: "학교", Kind: "열쇠:체언", Modifiable: true, Nominal: true, Glued: true}}}, true},
		{"a predicate written apart", rule.Around{Self: glued, Before: []rule.Near{noun}, After: []rule.Near{pred}}, true},
		{"apart from a noun", rule.Around{Self: apart, Before: []rule.Near{noun}}, false},
		{"onto a particle", rule.Around{Self: glued, Before: []rule.Near{{Stem: "학교", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}}}, false},
		{"a predicate written onto it", rule.Around{Self: glued, Before: []rule.Near{noun}, After: []rule.Near{predOnto}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
