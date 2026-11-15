//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestDigitsFrom — the digits in a row from a place: all of a run, those after the place only, none at a letter or past the end
package token

import "testing"

func TestDigitsFrom(t *testing.T) {
	cases := []struct {
		text string
		from int
		n    int
	}{
		{"2004다", 0, 4},
		{"2004다", 2, 2},
		{"2004다", 4, 0},
		{"2004다", 5, 0},
		{"a12", 0, 0},
	}
	for _, c := range cases {
		if got := digitsFrom([]rune(c.text), c.from); got != c.n {
			t.Errorf("%s from %d: %d, want %d", c.text, c.from, got, c.n)
		}
	}
}
