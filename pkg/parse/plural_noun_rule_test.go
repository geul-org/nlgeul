//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPluralNounRule — the rule is named as the trace prints it, adds bonusOverCount, waits for nothing, and speaks of a candidate that reads letters the word table gives as a nominal and as an adverb as the nominal of the table with the suffix 들 among its markers; not the adverb with 들, not the nominal with no 들, not a nominal the table gives as a nominal only
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPluralNounRule(t *testing.T) {
	lex := map[string]word{"모두": {kinds: kindNominal | kindAdverb}, "학교": {kinds: kindNominal}}
	r := pluralNounRule(lex)
	if r.name != rulePluralNoun || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "모두", Kind: "열쇠:체언", Last: "접미:들", Marks: "접미:들"}, true},
		{rule.Near{Stem: "모두", Kind: "열쇠:체언", Last: "보조사:은", Marks: "접미:들+보조사:은"}, true},
		{rule.Near{Stem: "모두", Kind: "열쇠:부사", Last: "접미:들", Marks: "접미:들"}, false},
		{rule.Near{Stem: "모두", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "학교", Kind: "열쇠:체언", Last: "접미:들", Marks: "접미:들"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
