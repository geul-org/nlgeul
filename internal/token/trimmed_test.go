//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestTrimmed — where the letters of a unit that are read begin and end (in letters): the signs at its edges left out; a sign at a number's front and a unit sign at its end kept; a path's head and its slash, an option's dashes and $ _ . before Latin letters kept; @ and # left out; nothing to read is 0 and 0; Hangul letters count one each
package token

import "testing"

func TestTrimmed(t *testing.T) {
	cases := []struct {
		in       string
		from, to int
	}{
		{"(곰이)", 1, 3},
		{"-3이다.", 0, 4},
		{"50%,", 0, 3},
		{"~/a/b/", 0, 6},
		{"--help", 0, 6},
		{"$HOME", 0, 5},
		{"@a", 1, 2},
		{"#태그", 1, 3},
		{"?!곰", 2, 3},
		{"ㆍ곰ㆍ", 1, 2},
		{"·곰·", 1, 2},
		{"ㆍ", 0, 0},
		{"...", 0, 0},
		{"", 0, 0},
	}
	for _, c := range cases {
		if from, to := trimmed(c.in); from != c.from || to != c.to {
			t.Errorf("%q: %d..%d, want %d..%d", c.in, from, to, c.from, c.to)
		}
	}
}
