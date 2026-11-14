//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestPartTokens — a part made only of Latin letters, digits and underscores is one token; a part with any other letter (Hangul, Greek, a dot, Hangul mixed with Latin) has one for each letter
package token

import "testing"

func TestPartTokens(t *testing.T) {
	cases := []struct {
		part string
		want int
	}{
		{"ab", 1},
		{"abc_12", 1},
		{"2026", 1},
		{"_", 1},
		{"café", 1},
		{"곰", 1},
		{"사자", 2},
		{"고양이", 3},
		{"곰a", 2},
		{"a곰b", 3},
		{"a.b", 3},
		{"a-b", 3},
		{"αβ", 2},
	}
	for _, c := range cases {
		if got := partTokens(c.part); got != c.want {
			t.Fatal(c.part, got, c.want)
		}
	}
}
