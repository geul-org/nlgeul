//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestCompoundFrontRuleGraph — the graph of compoundFrontRule evaluated by hand on rule.Around: the nominal 우선 stands with a noun written onto it, with or without its particle; it falls with a noun written apart after it, with a predicate's modifying form or closing written onto it, with a sign written onto it, and at the end of the sentence
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCompoundFrontRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "우선", Kind: "열쇠:체언"}
	r := compoundFrontRule(map[string]word{"우선": {kinds: kindNominal | kindAdverb}})
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"a noun written onto it", rule.Around{Self: self, After: []rule.Near{{Stem: "출자", Kind: "열쇠:체언", Modifiable: true, Glued: true}}}, true},
		{"a noun and its particle written onto it", rule.Around{Self: self, After: []rule.Near{{Stem: "출자", Kind: "열쇠:체언", Last: "목적격:를", Marks: "목적격:를", Modifiable: true, Glued: true}}}, true},
		{"a noun written apart", rule.Around{Self: self, After: []rule.Near{{Stem: "출자", Kind: "열쇠:체언", Modifiable: true}}}, false},
		{"a modifying form written onto it", rule.Around{Self: self, After: []rule.Near{{Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:는", Marks: "꾸밈:는", Pred: true, Modifiable: true, Glued: true}}}, false},
		{"a closing written onto it", rule.Around{Self: self, After: []rule.Near{{Stem: "하", Kind: "열쇠:동사", Last: "맺음:다", Marks: "맺음:다", Pred: true, Glued: true}}}, false},
		{"a sign written onto it", rule.Around{Self: self, After: []rule.Near{{Stem: ",", Kind: "쉼표", Glued: true}}}, false},
		{"at the end", rule.Around{Self: self}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
