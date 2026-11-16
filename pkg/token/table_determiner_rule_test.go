//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestTableDeterminerRule — the rule is named as the trace prints it, adds bonusSameLetters, and speaks of and waits for a candidate that reads letters the word table gives as a nominal and as a determiner as the determiner of the table with no particle (written apart or not); not the nominal of those letters, not one with a particle, not the determiner of letters the table gives as a determiner only, nor of letters of a closed set (a determiner or a pronoun of the closed set)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTableDeterminerRule(t *testing.T) {
	both := word{kinds: kindNominal | kindDeterminer}
	lex := map[string]word{"총": both, "한": both, "내": both, "각": {kinds: kindDeterminer}}
	closed := map[string]int{"한": kindDeterminer, "내": kindNominal}
	r := tableDeterminerRule(lex, closed)
	if r.name != ruleTableDeterminer || r.bonus != bonusSameLetters || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "총", Kind: "열쇠:관형사"}, true},
		{rule.Near{Stem: "총", Kind: "열쇠:관형사", Glued: true}, true},
		{rule.Near{Stem: "총", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "총", Kind: "열쇠:관형사", Last: "주격:이", Marks: "주격:이"}, false},
		{rule.Near{Stem: "각", Kind: "열쇠:관형사"}, false},
		{rule.Near{Stem: "한", Kind: "열쇠:관형사"}, false},
		{rule.Near{Stem: "내", Kind: "열쇠:관형사"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
