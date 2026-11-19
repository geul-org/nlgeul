//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestThingQuoteRule — the rule is named as the trace prints it, is a rebuttal of the force forceUngrammatical that adds nothing and reads nothing after the candidate, and speaks of 것 with 이라고 or 라고 as the particle of quotation (with a particle more too), not of 것 with the copula and -라고, nor of another nominal with 이라고
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestThingQuoteRule(t *testing.T) {
	r := thingQuoteRule()
	if r.name != ruleThingQuote || r.force != forceUngrammatical || r.bonus != 0 || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "것", Kind: "의존명사:것", Last: "부사격:이라고", Marks: "부사격:이라고"}, true},
		{rule.Near{Stem: "것", Kind: "의존명사:것", Last: "보조사:는", Marks: "부사격:이라고+보조사:는"}, true},
		{rule.Near{Stem: "것", Kind: "의존명사:것", Last: "부사격:라고", Marks: "접미:이+부사격:라고"}, true},
		{rule.Near{Stem: "것", Kind: "의존명사:것", Last: "맺음:라고", Marks: "서술격:이+맺음:라고", Pred: true}, false},
		{rule.Near{Stem: "책", Kind: "열쇠:체언", Last: "부사격:이라고", Marks: "부사격:이라고"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
