//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSyllablesFrom — the Hangul syllables in a row from a place: up to a digit or the end, none at a digit, a Latin letter or a lone jamo
package token

import "testing"

func TestSyllablesFrom(t *testing.T) {
	cases := []struct {
		text string
		from int
		n    int
	}{
		{"2008회합9", 4, 2},
		{"2008회합9", 5, 1},
		{"2008회합9", 0, 0},
		{"회합", 0, 2},
		{"aㄱ가", 1, 0},
	}
	for _, c := range cases {
		if got := syllablesFrom([]rune(c.text), c.from); got != c.n {
			t.Errorf("%s from %d: %d, want %d", c.text, c.from, got, c.n)
		}
	}
}
