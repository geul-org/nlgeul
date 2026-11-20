//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestHasMark — true when one of the candidate's markers (in any place of its names) is one of the names given; false for another name, a part of a name, no marker at all, or no names given
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHasMark(t *testing.T) {
	cases := []struct {
		marks string
		names []string
		want  bool
	}{
		{"부사격:에서", adverbialFirst, true},
		{"부사격:서", adverbialFirst, true},
		{"부사격:에서+보조사:는", adverbialFirst, true},
		{"보조사:는+부사격:서", adverbialFirst, true},
		{"주격:에서", adverbialFirst, false},
		{"주격:서+보조사:는", adverbialFirst, false},
		{"부사격:에", adverbialFirst, false},
		{"", adverbialFirst, false},
		{"부사격:서", nil, false},
	}
	for _, c := range cases {
		if got := hasMark(rule.Near{Stem: "학교", Marks: c.marks}, c.names); got != c.want {
			t.Errorf("hasMark(%q, %q) = %v; want %v", c.marks, c.names, got, c.want)
		}
	}
}
