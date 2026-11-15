//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAdverbialFirstRule — the rule of a sentence is named as the trace prints it, adds bonusOrder (small: an order among names of the same letters) and not bonusSentence, and speaks only of the candidates one of whose markers is the adverbial 에서 or 서 (not the nominative 에서, not a candidate with no marker, not another adverbial particle)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdverbialFirstRule(t *testing.T) {
	r := adverbialFirstRule()
	if r.name != ruleAdverbialFirst || r.graph == nil || r.bonus != bonusOrder || r.bonus >= bonusSentence || r.waits != nil {
		t.Fatal(r.name, r.bonus, r.waits != nil)
	}
	cases := []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "학교", Kind: "열쇠:체언", Last: "부사격:에서", Marks: "부사격:에서"}, true},
		{rule.Near{Stem: "학교에", Last: "부사격:서", Marks: "부사격:서"}, true},
		{rule.Near{Stem: "학교", Marks: "부사격:에서+보조사:는"}, true},
		{rule.Near{Stem: "학교", Marks: "주격:에서"}, false},
		{rule.Near{Stem: "학교", Marks: "부사격:에"}, false},
		{rule.Near{Stem: "학교에서"}, false},
		{rule.Near{}, false},
	}
	for _, c := range cases {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v; want %v", c.near, got, c.want)
		}
	}
}
