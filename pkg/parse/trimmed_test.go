//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestTrimmed — where the letters of a unit that are read begin and end (in letters): the signs at its edges left out; a sign at a number's front and a unit sign at its end kept (also after greater numbers written onto digits: 17억㎥, not after 억 alone); a path's head and its slash, an option's dashes and $ _ . before Latin letters kept; @ and # left out; nothing to read is 0 and 0; Hangul letters count one each
package parse

import "testing"

func TestTrimmed(t *testing.T) {
	cases := []struct {
		in       string
		from, to int
	}{
		{"(곰이)", 1, 3},
		{"-3이다.", 0, 4},
		{"50%,", 0, 3},
		{"17억㎥.", 0, 4},
		{"억㎥.", 0, 1},
		{"㎡당", 0, 2},
		{"㎡곰", 1, 2},
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
