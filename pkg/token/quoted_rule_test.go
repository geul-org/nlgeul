//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestQuotedRule — the rule is named as the trace prints it, adds bonusSameLetters, waits for nothing, and speaks only of a candidate that has the particle 라고 or 이라고 among its markers (not of 라고 as an ending, nor of a word with no marker)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestQuotedRule(t *testing.T) {
	r := quotedRule()
	if r.name != ruleQuoted || r.bonus != bonusSameLetters || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "책", Kind: "열쇠:체언", Last: "부사격:이라고", Marks: "부사격:이라고"}, true},
		{rule.Near{Stem: "", Last: "부사격:라고", Marks: "부사격:라고"}, true},
		{rule.Near{Stem: "책", Last: "접미:이", Marks: "접미:이+부사격:라고"}, true},
		{rule.Near{Last: "맺음:라고", Marks: "맺음:라고"}, false},
		{rule.Near{Stem: "책", Kind: "열쇠:체언"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
