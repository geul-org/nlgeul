//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestNounUnmodifiedRule — the rule is named as the trace prints it, adds bonusSameLetters, and speaks of and waits for a candidate that reads 수 or 줄 as a noun of the word table alone (not as the dependent noun, not another noun, not 줄다)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounUnmodifiedRule(t *testing.T) {
	r := nounUnmodifiedRule(adnBound)
	if r.name != ruleNounUnmodified || r.bonus != bonusSameLetters || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "수", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는"}, true},
		{rule.Near{Stem: "줄", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "수", Kind: "의존명사:수"}, false},
		{rule.Near{Stem: "시", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "줄", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true, Marks: "꾸밈:ㄴ"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
