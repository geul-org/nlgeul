//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestInterjectionAloneRule — the rule is named as the trace prints it, adds bonusOverCount, waits for nothing, and speaks of a candidate that reads letters the word table gives as a nominal and as an interjection as the word of the table that stands alone (열쇠:홀로) with no particle; not the nominal of those letters, not one with a particle, not the interjection of letters the table gives as an interjection only
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestInterjectionAloneRule(t *testing.T) {
	lex := map[string]word{"흠": {kinds: kindNominal | kindAlone}, "아이고": {kinds: kindAlone}}
	r := interjectionAloneRule(lex)
	if r.name != ruleInterjectionAlone || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "흠", Kind: "열쇠:홀로"}, true},
		{rule.Near{Stem: "흠", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "흠", Kind: "열쇠:홀로", Last: "보조사:은", Marks: "보조사:은"}, false},
		{rule.Near{Stem: "아이고", Kind: "열쇠:홀로"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
