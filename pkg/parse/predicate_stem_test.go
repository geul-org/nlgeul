//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestPredicateStem — a verb's and an adjective's stem of the word table and an auxiliary verb of the closed set are predicate stems; a noun of the word table, a number, a new word and a determiner of the closed set are not
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPredicateStem(t *testing.T) {
	for _, c := range []struct {
		kind string
		want bool
	}{
		{"열쇠:동사", true},
		{"열쇠:형용사", true},
		{"보조 용언:있", true},
		{"열쇠:체언", false},
		{"수", false},
		{"", false},
		{"관형사:한", false},
	} {
		if got := predicateStem(rule.Near{Kind: c.kind}); got != c.want {
			t.Errorf("predicateStem(%q) = %v, want %v", c.kind, got, c.want)
		}
	}
}
