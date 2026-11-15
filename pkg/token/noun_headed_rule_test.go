//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNounHeadedRule — the rule is named as the trace prints it, adds bonusSentence, and speaks of and waits for a candidate that reads a dependent noun's letters as a noun of the word table alone (not the dependent noun itself, not another noun)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounHeadedRule(t *testing.T) {
	r := nounHeadedRule(boundNext)
	if r.name != ruleNounHeaded || r.bonus != bonusSentence || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "시", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는"}, true},
		{rule.Near{Stem: "때", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "시", Kind: "의존명사:시"}, false},
		{rule.Near{Stem: "인구", Kind: "열쇠:체언"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
