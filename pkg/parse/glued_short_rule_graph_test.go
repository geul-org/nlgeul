//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestGluedShortRuleGraph — the graph of gluedShortRule evaluated by hand on rule.Around: a word glued to a word that ends in a nominal stands (a noun, a number with its unit); a word with a space before it falls, and so does a glued word after a word that ends in a particle or a closing, or with no word before it, or with a word it may modify written right onto it (명의신탁: the unit goes on with a nominal), but not when what is written onto it is a sign or a predicate
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestGluedShortRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "자", Kind: "열쇠:체언", Marks: "관형격:의", Text: "자의", Glued: true}
	spaced := self
	spaced.Glued = false
	noun := rule.Near{Stem: "패소", Kind: "열쇠:체언", Nominal: true}
	unit := rule.Near{Stem: "3", Kind: "수", Last: "단위:차", Nominal: true}
	marked := rule.Near{Stem: "패소", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}
	pred := rule.Near{Stem: "먹", Kind: "열쇠:동사", Last: "꾸밈:은", Pred: true}
	g := gluedShortRule().graph
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after a noun", rule.Around{Self: self, Before: []rule.Near{noun}}, true},
		{"after a unit", rule.Around{Self: self, Before: []rule.Near{unit}}, true},
		{"spaced", rule.Around{Self: spaced, Before: []rule.Near{noun}}, false},
		{"after a particle", rule.Around{Self: self, Before: []rule.Near{marked}}, false},
		{"after a modifying form", rule.Around{Self: self, Before: []rule.Near{pred}}, false},
		{"nothing before", rule.Around{Self: self}, false},
		{"a nominal written onto it", rule.Around{Self: self, Before: []rule.Near{noun}, After: []rule.Near{{Stem: "신탁", Kind: "열쇠:체언", Glued: true, Modifiable: true}}}, false},
		{"a nominal after a space", rule.Around{Self: self, Before: []rule.Near{noun}, After: []rule.Near{{Stem: "신탁", Kind: "열쇠:체언", Modifiable: true}}}, true},
		{"a sign written onto it", rule.Around{Self: self, Before: []rule.Near{noun}, After: []rule.Near{{Stem: ",", Kind: "쉼표", Glued: true}}}, true},
		{"a predicate written onto it", rule.Around{Self: self, Before: []rule.Near{noun}, After: []rule.Near{{Stem: "하", Kind: "열쇠:동사", Last: "맺음:다", Pred: true, Glued: true}}}, true},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, c.around)
		if v, _ := evaluate(g, ctx, false); (v > 0) != c.want {
			t.Errorf("%s: %v, want %v", c.name, v, c.want)
		}
	}
}
