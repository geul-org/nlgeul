//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestDeterminerAloneRule — the rule is named as the trace prints it, adds bonusSameLetters, and speaks of and waits for a candidate that reads 이, 그 or 저 as the determiner of the closed set, written apart or not (not the pronoun of the word table, not the determiner of the word table, not another determiner of the closed set)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDeterminerAloneRule(t *testing.T) {
	r := determinerAloneRule()
	if r.name != ruleDeterminerAlone || r.bonus != bonusSameLetters || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "이", Kind: "관형사:이"}, true},
		{rule.Near{Stem: "그", Kind: "관형사:그", Glued: true}, true},
		{rule.Near{Stem: "저", Kind: "관형사:저"}, true},
		{rule.Near{Stem: "이", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "그", Kind: "열쇠:관형사"}, false},
		{rule.Near{Stem: "한", Kind: "관형사:한"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
