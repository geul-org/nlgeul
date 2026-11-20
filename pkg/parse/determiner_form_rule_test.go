//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestDeterminerFormRule — the rule is named as the trace prints it, adds bonusOverUncounted, does not wait, and speaks of a candidate that reads 이런, 그런, 저런, 어떤 or 다른 as a determiner of the word table or of the closed set, written apart or not (not the adjective's form, not the noun, not 이 or 한 of the closed set)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDeterminerFormRule(t *testing.T) {
	r := determinerFormRule()
	if r.name != ruleDeterminerForm || r.bonus != bonusOverUncounted || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "다른", Kind: "열쇠:관형사"}, true},
		{rule.Near{Stem: "그런", Kind: "열쇠:관형사", Glued: true}, true},
		{rule.Near{Stem: "저런", Kind: "열쇠:관형사"}, true},
		{rule.Near{Stem: "이런", Kind: "관형사:이런"}, true},
		{rule.Near{Stem: "어떤", Kind: "관형사:어떤"}, true},
		{rule.Near{Stem: "다르", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Pred: true}, false},
		{rule.Near{Stem: "그런", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "이", Kind: "관형사:이"}, false},
		{rule.Near{Stem: "한", Kind: "관형사:한"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
