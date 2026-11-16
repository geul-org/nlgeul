//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestHaOneRule — the rule is named as the trace prints it, adds bonusOverCount, does not wait, and speaks of a candidate that reads the word 한 as the verb 하 of the word table with the modifying ㄴ, written apart or not (not the determiner 한, not the noun 한, not 하 with ㄴ in a word of other letters, not 하 as a noun)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHaOneRule(t *testing.T) {
	r := haOneRule()
	if r.name != ruleHaOne || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Text: "한", Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true}, true},
		{rule.Near{Text: "한", Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true, Glued: true}, true},
		{rule.Near{Text: "한", Stem: "한", Kind: "관형사:한"}, false},
		{rule.Near{Text: "한", Stem: "한", Kind: "열쇠:체언"}, false},
		{rule.Near{Text: "했던", Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:던", Pred: true}, false},
		{rule.Near{Text: "한", Stem: "하", Kind: "열쇠:체언", Last: "꾸밈:ㄴ", Pred: true}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
