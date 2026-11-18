//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNounUnheldRule — the rule is named as the trace prints it, adds bonusSameLetters, waits for nothing, and speaks of a candidate that reads 시, 간 or 적 as a noun of the word table (not as the dependent noun, not 수, not another noun, not a verb's stem); the words it asks are those of the dependent-noun table written with that use
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounUnheldRule(t *testing.T) {
	if !slices.Equal(unheldBound, []string{"간", "적", "시"}) {
		t.Fatal(unheldBound)
	}
	r := nounUnheldRule(unheldBound)
	if r.name != ruleNounUnheld || r.bonus != bonusSameLetters || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "시", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는"}, true},
		{rule.Near{Stem: "간", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "적", Kind: "열쇠:체언", Last: "목적격:을", Marks: "목적격:을"}, true},
		{rule.Near{Stem: "시", Kind: "의존명사:시"}, false},
		{rule.Near{Stem: "수", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "인구", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "적", Kind: "열쇠:동사", Last: "꾸밈:은", Pred: true, Marks: "꾸밈:은"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
