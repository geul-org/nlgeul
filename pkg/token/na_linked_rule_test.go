//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNaLinkedRule — the rule is named as the trace prints it, adds bonusShortTaken, does not wait, and speaks of the stem 나 read as the verb with a closing (나서, 난, 나니), not of the verb with no closing, the pronoun 나 with a particle, the noun 나서, nor a stem other than 나
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNaLinkedRule(t *testing.T) {
	r := naLinkedRule()
	if r.name != ruleNaLinked || r.bonus != bonusShortTaken || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "나", Kind: "열쇠:동사", Last: "맺음:어서(합침)", Marks: "맺음:어서(합침)", Pred: true, Text: "나서"}, true},
		{rule.Near{Stem: "나", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Marks: "꾸밈:ㄴ", Pred: true, Text: "난"}, true},
		{rule.Near{Stem: "나", Kind: "열쇠:동사", Last: "맺음:니", Marks: "맺음:니", Pred: true, Text: "나니"}, true},
		{rule.Near{Stem: "나", Kind: "열쇠:동사", Text: "나"}, false},
		{rule.Near{Stem: "나", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는", Text: "나는"}, false},
		{rule.Near{Stem: "나서", Kind: "열쇠:체언", Text: "나서"}, false},
		{rule.Near{Stem: "남", Kind: "열쇠:동사", Last: "꾸밈:은", Marks: "꾸밈:은", Pred: true, Text: "남은"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
