//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestGiveLinkedRule — the rule is named as the trace prints it, adds bonusOverCount, reads nothing after the candidate, and speaks of a candidate that reads 달 as a verb of the word table (not the noun 달, not 다르다, not another verb)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGiveLinkedRule(t *testing.T) {
	r := giveLinkedRule()
	if r.name != ruleGiveLinked || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "달", Kind: "열쇠:동사", Last: "맺음:라"}, true},
		{rule.Near{Stem: "달", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "다르", Kind: "열쇠:형용사", Last: "맺음:어(합침)"}, false},
		{rule.Near{Stem: "주", Kind: "열쇠:동사", Last: "맺음:라"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
