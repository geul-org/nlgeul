//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestPredStem — a verb and an adjective of the word table and an auxiliary verb of the closed set are a predicate's stem; a noun, a dependent noun, a determiner, an adverb and a new word (no name) are not
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPredStem(t *testing.T) {
	for _, c := range []struct {
		kind string
		want bool
	}{
		{"열쇠:동사", true},
		{"열쇠:형용사", true},
		{"보조 용언:있", true},
		{"열쇠:체언", false},
		{"의존명사:것", false},
		{"관형사:이", false},
		{"열쇠:부사", false},
		{"", false},
	} {
		if got := predStem(rule.Near{Stem: "가", Kind: c.kind}); got != c.want {
			t.Errorf("%q: %v, want %v", c.kind, got, c.want)
		}
	}
}
