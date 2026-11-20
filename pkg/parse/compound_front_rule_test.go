//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestCompoundFrontRule — the rule is named as the trace prints it, adds bonusOverCount, and speaks of and waits for a candidate that reads letters the word table gives as a nominal and as an adverb as the nominal of the table with no particle; not the adverb of those letters, not one with a particle, not a nominal the table gives as a nominal only
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCompoundFrontRule(t *testing.T) {
	lex := map[string]word{"우선": {kinds: kindNominal | kindAdverb}, "학교": {kinds: kindNominal}}
	r := compoundFrontRule(lex)
	if r.name != ruleCompoundFront || r.bonus != bonusOverCount || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "우선", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "우선", Kind: "열쇠:체언", Glued: true}, true},
		{rule.Near{Stem: "우선", Kind: "열쇠:부사"}, false},
		{rule.Near{Stem: "우선", Kind: "열쇠:체언", Last: "주격:이", Marks: "주격:이"}, false},
		{rule.Near{Stem: "학교", Kind: "열쇠:체언"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
