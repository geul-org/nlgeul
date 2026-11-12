//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestOntoEnd — what is written onto a closing sign or a span of code ends at the next space, sign of a pair or span of code; at the end, at a space, at a sign of punctuation or at a symbol nothing is written onto it
package token

import "testing"

func TestOntoEnd(t *testing.T) {
	ix, _ := newIndex(Options{})
	cases := []struct {
		text string
		from int
		want int
	}{
		{")를 가", 1, 2},
		{")를(나)", 1, 2},
		{")를`a`", 1, 2},
		{")", 1, 1},
		{") 가", 1, 1},
		{").", 1, 1},
		{")~", 1, 1},
		{")+가", 1, 1},
		{")가나다", 1, 4},
	}
	for _, c := range cases {
		if got := newCutter(ix, c.text, false).ontoEnd(c.from); got != c.want {
			t.Errorf("%q from %d: %d, want %d", c.text, c.from, got, c.want)
		}
	}
}
