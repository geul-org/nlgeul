//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestArgumentAdnRule — the rule is named as the trace prints it, adds bonusOverCount, and speaks of and waits for a candidate whose last piece is a modifying closing and whose stem is a verb's, an adjective's or an auxiliary verb's (not a nominal with the copula, not a new word, not a candidate ending otherwise, nor a word whose letters are a pronoun of the closed sets with a particle: 나는, 이를)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestArgumentAdnRule(t *testing.T) {
	r := argumentAdnRule()
	if r.name != ruleArgumentAdn || r.bonus != bonusOverCount || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "팔", Kind: "열쇠:동사", Last: "꾸밈:는"}, true},
		{rule.Near{Stem: "검", Kind: "열쇠:형용사", Last: "꾸밈:은"}, true},
		{rule.Near{Stem: "있", Kind: "보조 용언:있", Last: "꾸밈:는"}, true},
		{rule.Near{Stem: "파", Kind: "열쇠:체언", Last: "보조사:는"}, false},
		{rule.Near{Stem: "팔", Kind: "열쇠:동사", Last: "맺음:다"}, false},
		{rule.Near{Stem: "12", Kind: "수", Last: "꾸밈:ㄹ"}, false},
		{rule.Near{Stem: "허겁", Kind: "", Last: "꾸밈:을"}, false},
		{rule.Near{Stem: "팔", Kind: "열쇠:동사", Last: "꾸밈:는", Text: "파는"}, true},
		{rule.Near{Stem: "나", Kind: "열쇠:동사", Last: "꾸밈:는", Text: "나는"}, false},
		{rule.Near{Stem: "이르", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Text: "이를"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
