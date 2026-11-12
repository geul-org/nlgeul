//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAfterSign — a sign of a pair that only opens ends its unit right after it; one with a closing candidate takes in what is written onto it (the particle up to the next space or sign); a closing sign with a space or another sign after it ends there
package token

import "testing"

func TestAfterSign(t *testing.T) {
	ix, _ := newIndex(Options{})
	opens, shuts := []pairing{{openName, 1}}, []pairing{{shutName, 1}}
	cases := []struct {
		text       string
		from       int
		candidates []pairing
		want       int
	}{
		{"(가", 1, opens, 1},
		{"(가", 1, append(append([]pairing{}, opens...), shuts...), 2},
		{"(가)를 나", 3, shuts, 4},
		{"(가) 나", 3, shuts, 3},
		{"(가).", 3, shuts, 3},
		{"(가)", 3, shuts, 3},
		{"(가)를(나)", 3, shuts, 4},
	}
	for _, c := range cases {
		if got := newCutter(ix, c.text, false).afterSign(c.from, c.candidates); got != c.want {
			t.Errorf("%q from %d: %d, want %d", c.text, c.from, got, c.want)
		}
	}
}
