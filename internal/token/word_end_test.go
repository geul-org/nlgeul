//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestWordEnd — letters stop at a space, at a sign that is read as a sign of a pair in some way, at a span of code, or at the end; a sign that is no sign of a pair there (it's, x<y spaced apart, a closing sign with nothing open) stays in the letters
package token

import "testing"

func TestWordEnd(t *testing.T) {
	ix, _ := newIndex(Options{})
	cases := []struct {
		text string
		from int
		want int
	}{
		{"곰이(가", 1, 2},
		{"곰이 가", 0, 2},
		{"곰`a`", 1, 1},
		{"곰이", 1, 2},
		{"it's a", 1, 4},
		{"곰이)가", 1, 4},
		{"곰이\"가", 1, 2},
		{"a<b", 0, 1},
		{"a->b", 0, 4},
		{"", 0, 0},
	}
	for _, c := range cases {
		if got := newCutter(ix, c.text, false).wordEnd(c.from); got != c.want {
			t.Errorf("%q from %d: %d, want %d", c.text, c.from, got, c.want)
		}
	}
}
