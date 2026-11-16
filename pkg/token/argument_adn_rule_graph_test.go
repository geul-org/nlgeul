//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestArgumentAdnRuleGraph — the graph of argumentAdnRule evaluated by hand on rule.Around: a modifying form stands after a word ending in a subject, complement, object or adverbial case particle and before a noun, a determiner or another word a modifier may stand before; it falls after a topic particle, after a bare noun and at the head, before an adverb, a predicate and at the end, and before a word that ends in a modifying closing itself
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestArgumentAdnRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "팔", Kind: "열쇠:동사", Last: "꾸밈:는", Pred: true, Marks: "꾸밈:는"}
	arg := func(last string) []rule.Near {
		return []rule.Near{{Stem: "물", Kind: "열쇠:체언", Last: last, Marks: last}}
	}
	noun := []rule.Near{{Stem: "가게", Kind: "열쇠:체언", Modifiable: true}}
	r := argumentAdnRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after an object, before a noun", rule.Around{Self: self, Before: arg("목적격:을"), After: noun}, true},
		{"after a subject", rule.Around{Self: self, Before: arg("주격:이"), After: noun}, true},
		{"after a complement", rule.Around{Self: self, Before: arg("보격:이"), After: noun}, true},
		{"after an adverbial", rule.Around{Self: self, Before: arg("부사격:에"), After: noun}, true},
		{"before a determiner", rule.Around{Self: self, Before: arg("목적격:을"), After: []rule.Near{{Stem: "그", Kind: "관형사:그", Modifiable: true}}}, true},
		{"after a topic", rule.Around{Self: self, Before: arg("보조사:은"), After: noun}, false},
		{"after a bare noun", rule.Around{Self: self, Before: []rule.Near{{Stem: "물", Kind: "열쇠:체언"}}, After: noun}, false},
		{"at the head", rule.Around{Self: self, After: noun}, false},
		{"before an adverb", rule.Around{Self: self, Before: arg("목적격:을"), After: []rule.Near{{Stem: "빨리", Kind: "열쇠:부사"}}}, false},
		{"before a predicate", rule.Around{Self: self, Before: arg("목적격:을"), After: []rule.Near{{Stem: "가", Kind: "열쇠:동사", Last: "맺음:다", Pred: true}}}, false},
		{"at the end", rule.Around{Self: self, Before: arg("목적격:을")}, false},
		{"before a modifying form", rule.Around{Self: self, Before: arg("목적격:을"), After: []rule.Near{{Stem: "쓰", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true, Modifiable: true}}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
