//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestTouched — a sign with another sign that is no punctuation mark right next to it is touched (the = of =>, the bar of |-, ★ of ★☆, on either side); one with a punctuation mark of a sentence, a letter, a unit sign read with its number or nothing next to it is not
package parse

import "testing"

func TestTouched(t *testing.T) {
	for _, c := range []struct {
		text        string
		start, stop int
		want        bool
	}{
		{"=>", 0, 1, true}, {"|-", 0, 1, true}, {"★☆", 0, 1, true}, {"★☆", 1, 2, true},
		{"○○·△△", 0, 2, false}, {"○○,", 0, 2, false}, {"(○)", 1, 2, false}, {"가○나", 1, 2, false},
		{"5%↓", 2, 3, false}, {"가%↓", 2, 3, true}, {"○", 0, 1, false},
	} {
		if got := touched([]rune(c.text), c.start, c.stop); got != c.want {
			t.Errorf("%q %d-%d: %v", c.text, c.start, c.stop, got)
		}
	}
}
