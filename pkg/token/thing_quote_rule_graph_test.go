//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestThingQuoteRuleGraph — the graph of thingQuoteRule evaluated by hand on rule.Around: the rebuttal stands for 것 with the particle of quotation whatever stands around it (after a modifying form, at the head of a sentence, before a verb), and not for a candidate with no such particle
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestThingQuoteRuleGraph(t *testing.T) {
	quote := rule.Near{Stem: "것", Kind: "의존명사:것", Last: "부사격:이라고", Marks: "부사격:이라고"}
	copula := rule.Near{Stem: "것", Kind: "의존명사:것", Last: "맺음:라고", Marks: "서술격:이+맺음:라고", Pred: true}
	adn := []rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "꾸밈:은", Marks: "꾸밈:은", Pred: true}}
	verb := []rule.Near{{Stem: "하", Kind: "열쇠:동사", Last: "맺음:다", Marks: "맺음:다", Pred: true}}
	r := thingQuoteRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after a modifying form", rule.Around{Self: quote, Before: adn, After: verb}, true},
		{"at the head", rule.Around{Self: quote}, true},
		{"the copula with -라고", rule.Around{Self: copula, Before: adn, After: verb}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
