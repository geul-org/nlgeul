//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNumberQuoteRule — the rule is named as the trace prints it, adds bonusSameLetters, waits for nothing, and speaks only of a candidate whose stem is a number or an ordinal number and that has the particle 라고 or 이라고 among its markers (not of 라고 as an ending, nor of 이라고 after a noun)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNumberQuoteRule(t *testing.T) {
	r := numberQuoteRule()
	if r.name != ruleNumberQuote || r.bonus != bonusSameLetters || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "②", Kind: "수", Last: "부사격:라고", Marks: "부사격:라고"}, true},
		{rule.Near{Stem: "3", Kind: "수", Last: "부사격:이라고", Marks: "부사격:이라고"}, true},
		{rule.Near{Stem: "85", Kind: "수", Last: "부사격:라고", Marks: "단위:퍼센트+부사격:라고"}, true},
		{rule.Near{Stem: "제3", Kind: "차례", Last: "부사격:라고", Marks: "단위:자+부사격:라고"}, true},
		{rule.Near{Stem: "②", Kind: "수", Last: "맺음:라고", Marks: "맺음:라고"}, false},
		{rule.Near{Stem: "책", Kind: "열쇠:체언", Last: "부사격:이라고", Marks: "부사격:이라고"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
