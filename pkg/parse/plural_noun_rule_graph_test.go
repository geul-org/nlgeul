//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPluralNounRuleGraph — the graph of pluralNounRule evaluated by hand on rule.Around: the nominal 모두 stands with the suffix 들 after it, alone or before a particle; it falls with a particle and no 들
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPluralNounRuleGraph(t *testing.T) {
	r := pluralNounRule(map[string]word{"모두": {kinds: kindNominal | kindAdverb}})
	for _, c := range []struct {
		name string
		self rule.Near
		want bool
	}{
		{"들", rule.Near{Stem: "모두", Kind: "열쇠:체언", Last: "접미:들", Marks: "접미:들"}, true},
		{"들 and a particle", rule.Near{Stem: "모두", Kind: "열쇠:체언", Last: "보조사:은", Marks: "접미:들+보조사:은"}, true},
		{"a particle", rule.Near{Stem: "모두", Kind: "열쇠:체언", Last: "보조사:은", Marks: "보조사:은"}, false},
	} {
		if got := graphHolds(t, r, rule.Around{Self: c.self}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
