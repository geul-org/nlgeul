//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestArgumentAdnRuleForm — the rule speaks of a modifying form whatever its letters: of one whose letters are a demonstrative determiner too (그런, 다른: where determinerFormRule holds the determiner back, this rule carries the adjective's form), and of others (파는, and 한, which haOneRule judges too)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestArgumentAdnRuleForm(t *testing.T) {
	r := argumentAdnRule()
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "그렇", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Text: "그런"}, true},
		{rule.Near{Stem: "다르", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Text: "다른"}, true},
		{rule.Near{Stem: "팔", Kind: "열쇠:동사", Last: "꾸밈:는", Text: "파는"}, true},
		{rule.Near{Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Text: "한"}, true},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
