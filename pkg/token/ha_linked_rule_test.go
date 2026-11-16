//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestHaLinkedRule — the rule is named as the trace prints it, adds bonusShortTaken, does not wait, and speaks of a candidate that reads 하 as the verb with a closing (not the noun 하, not 하 with no closing, not another verb, not 하 as an auxiliary of the closed set)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHaLinkedRule(t *testing.T) {
	r := haLinkedRule()
	if r.name != ruleHaLinked || r.bonus != bonusShortTaken || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "하", Kind: "열쇠:동사", Last: "맺음:나", Marks: "맺음:나", Pred: true}, true},
		{rule.Near{Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Marks: "꾸밈:ㄴ", Pred: true}, true},
		{rule.Near{Stem: "하", Kind: "열쇠:체언", Last: "보조사:나", Marks: "보조사:나"}, false},
		{rule.Near{Stem: "하", Kind: "열쇠:동사"}, false},
		{rule.Near{Stem: "되", Kind: "열쇠:동사", Last: "맺음:나", Marks: "맺음:나", Pred: true}, false},
		{rule.Near{Stem: "하", Kind: "보조 용언:하", Last: "맺음:나", Marks: "맺음:나", Pred: true}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
