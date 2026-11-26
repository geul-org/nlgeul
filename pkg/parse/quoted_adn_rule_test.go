//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestQuotedAdnRule — the rule is named as the trace prints it, adds bonusSameLetters, waits for nothing, and speaks only of a candidate that has the modifying 는 among its markers (not of the particle 는, nor of another modifying form, nor of a word with no marker)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestQuotedAdnRule(t *testing.T) {
	r := quotedAdnRule()
	if r.name != ruleQuotedAdn || r.bonus != bonusSameLetters || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Last: "꾸밈:는", Marks: "꾸밈:는"}, true},
		{rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "꾸밈:는", Marks: "꾸밈:는"}, true},
		{rule.Near{Last: "보조사:는", Marks: "보조사:는"}, false},
		{rule.Near{Last: "꾸밈:ㄴ", Marks: "꾸밈:ㄴ"}, false},
		{rule.Near{Stem: "는", Kind: "열쇠:체언"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
