//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLineFacts — the facts of a unit's place its letters do not show: its line holds no letter or digit (MarkBare), a space or the sentence's edge stands on both sides of it (MarkApart), it is one hyphen alone where a sum is written (MarkSum)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLineFacts(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text   string
		i, end int
		want   int
	}{
		{"---", 0, 3, rule.MarkBare | rule.MarkApart},
		{"|---|", 1, 4, rule.MarkBare},
		{"가 - 나", 2, 3, rule.MarkApart},
		{"100 - 20", 4, 5, rule.MarkApart | rule.MarkSum},
		{"가-나", 1, 2, 0},
		{"가 -나", 2, 3, 0},
		{"가- 나", 1, 2, 0},
		{"가 -", 2, 3, rule.MarkApart},
	} {
		if got := newCutter(ix, c.text, false).lineFacts(c.i, c.end); got != c.want {
			t.Errorf("%q %d-%d: %b, want %b", c.text, c.i, c.end, got, c.want)
		}
	}
}
