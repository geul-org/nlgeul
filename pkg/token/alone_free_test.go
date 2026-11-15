//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAloneFree — a candidate whose stem is a bound noun or a determiner, by the name of a closed-set word (의존명사:시, 관형사:이) or of the word table (열쇠:관형사), is not free; a noun, a number, a verb, a new word (no name) and a candidate with no stem are
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAloneFree(t *testing.T) {
	for _, c := range []struct {
		kind string
		want bool
	}{
		{"의존명사:시", false},
		{"관형사:이", false},
		{"열쇠:관형사", false},
		{"열쇠:체언", true},
		{"열쇠:동사", true},
		{"수", true},
		{"", true},
		{"보조 용언:있", true},
	} {
		if got := aloneFree(rule.Near{Stem: "시", Kind: c.kind}); got != c.want {
			t.Errorf("aloneFree(%q) = %v", c.kind, got)
		}
	}
}
