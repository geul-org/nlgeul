//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAdverbAloneRule — the rule is named as the trace prints it, adds bonusOverCount, waits for nothing, and speaks of a candidate that reads letters the word table gives as a nominal and as an adverb as the adverb of the table with no particle (written apart or not); not the nominal of those letters, not one with a particle, not the adverb of letters the table gives as an adverb only, nor of letters the table does not hold
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdverbAloneRule(t *testing.T) {
	lex := map[string]word{"모두": {kinds: kindNominal | kindAdverb}, "다시": {kinds: kindAdverb}}
	r := adverbAloneRule(lex)
	if r.name != ruleAdverbAlone || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "모두", Kind: "열쇠:부사"}, true},
		{rule.Near{Stem: "모두", Kind: "열쇠:부사", Glued: true}, true},
		{rule.Near{Stem: "모두", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "모두", Kind: "열쇠:부사", Last: "보조사:는", Marks: "보조사:는"}, false},
		{rule.Near{Stem: "다시", Kind: "열쇠:부사"}, false},
		{rule.Near{Stem: "직접", Kind: "열쇠:부사"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
