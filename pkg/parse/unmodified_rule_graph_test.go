//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestUnmodifiedRuleGraph — the graph of unmodifiedRule evaluated by hand on rule.Around: the noun stands before an adverb, a sign, a predicate and at the end of the sentence; it falls before a word an adnominal may stand before (a noun, a determiner, a modifying form); only the word right after it counts
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestUnmodifiedRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "감정인", Kind: "열쇠:체언"}
	noun := rule.Near{Stem: "사람", Kind: "열쇠:체언", Nominal: true, Modifiable: true}
	det := rule.Near{Stem: "그", Kind: "관형사:그", Modifiable: true}
	adn := rule.Near{Stem: "착하", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Pred: true, Modifiable: true}
	adverb := rule.Near{Stem: "및", Kind: "열쇠:부사"}
	sign := rule.Near{Last: "쉼표", Marks: "쉼표"}
	pred := rule.Near{Stem: "오", Kind: "열쇠:동사", Last: "맺음:다", Pred: true}
	g := unmodifiedRule().graph
	for _, c := range []struct {
		name  string
		after []rule.Near
		want  bool
	}{
		{"before an adverb", []rule.Near{adverb, noun}, true},
		{"before a sign", []rule.Near{sign}, true},
		{"before a predicate", []rule.Near{pred}, true},
		{"at the end", nil, true},
		{"before a noun", []rule.Near{noun}, false},
		{"before a determiner", []rule.Near{det, noun}, false},
		{"before a modifying form", []rule.Near{adn, noun}, false},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, rule.Around{Self: self, After: c.after})
		if v, _ := evaluate(g, ctx, false); (v > 0) != c.want {
			t.Errorf("%s: %v, want %v", c.name, v, c.want)
		}
	}
}
