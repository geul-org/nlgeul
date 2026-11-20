//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestGluedNominalRule — the rule is named as the trace prints it, adds bonusOverCount, and speaks of and waits for a candidate that reads letters the word table gives as a nominal and as a determiner or an adverb as the nominal of the table with no particle; not the determiner of those letters, not one with a particle, not a nominal the table gives as a nominal only, nor the letters of a closed-set word
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGluedNominalRule(t *testing.T) {
	lex := map[string]word{"전": {kinds: kindNominal | kindDeterminer}, "계속": {kinds: kindNominal | kindAdverb}, "학교": {kinds: kindNominal}, "이": {kinds: kindNominal | kindDeterminer}}
	r := gluedNominalRule(lex, map[string]int{"이": kindDeterminer})
	if r.name != ruleGluedNominal || r.bonus != bonusOverCount || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "전", Kind: "열쇠:체언", Glued: true}, true},
		{rule.Near{Stem: "계속", Kind: "열쇠:체언", Glued: true}, true},
		{rule.Near{Stem: "전", Kind: "열쇠:관형사", Glued: true}, false},
		{rule.Near{Stem: "전", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에", Glued: true}, false},
		{rule.Near{Stem: "학교", Kind: "열쇠:체언", Glued: true}, false},
		{rule.Near{Stem: "이", Kind: "열쇠:체언", Glued: true}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
