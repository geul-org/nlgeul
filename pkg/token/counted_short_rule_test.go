//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestCountedShortRule — the rule is named as the trace prints it, adds bonusOverCount, waits for nothing, and speaks of a noun of the word table of one syllable with a particle written after it (대+가, 편+의), not of one whose letters are a dependent noun's (때+에), not of the noun of two syllables (대가), not of a bare noun (대) and not of another name (대[의존명사:대]+가)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCountedShortRule(t *testing.T) {
	r := countedShortRule(numDeterminers)
	if r.name != ruleCountedShort || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "대", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가", Text: "대가"}, true},
		{rule.Near{Stem: "편", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의", Text: "편의"}, true},
		{rule.Near{Stem: "때", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에", Text: "때에"}, false},
		{rule.Near{Stem: "대가", Kind: "열쇠:체언", Text: "대가"}, false},
		{rule.Near{Stem: "대", Kind: "열쇠:체언", Text: "대"}, false},
		{rule.Near{Stem: "대", Kind: "의존명사:대", Last: "주격:가", Marks: "주격:가", Text: "대가"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
