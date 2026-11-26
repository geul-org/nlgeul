//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestApartAt — a place has a space or the edge of the letters on both sides: alone, at either edge next to a space, between spaces; not with a letter on a side
package parse

import "testing"

func TestApartAt(t *testing.T) {
	cases := []struct {
		text string
		at   int
		want bool
	}{
		{"<", 0, true},
		{"< 가", 0, true},
		{"가 >", 2, true},
		{"가 < 나", 2, true},
		{"가\t<\n나", 2, true},
		{"<가", 0, false},
		{"가>", 1, false},
		{"가 <나", 2, false},
		{"가< 나", 1, false},
	}
	for _, c := range cases {
		if got := apartAt([]rune(c.text), c.at); got != c.want {
			t.Errorf("%q at %d: %v", c.text, c.at, got)
		}
	}
}
