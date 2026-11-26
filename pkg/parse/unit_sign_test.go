//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestUnitSign — a sign of the unit table right after a digit is read with its number (the % of 90%, the ℃ of 3℃, the $ of 300$); not one after a letter, not a sign that is no unit, not a unit that is a letter, not a place outside the letters
package parse

import "testing"

func TestUnitSign(t *testing.T) {
	for _, c := range []struct {
		text string
		k    int
		want bool
	}{
		{"90%+", 2, true},
		{"3℃", 1, true},
		{"300$", 3, true},
		{"가%", 1, false},
		{"3+", 1, false},
		{"3개", 1, false},
		{"%", 0, false},
		{"3%", 2, false},
		{"3%", -1, false},
	} {
		if got := unitSign([]rune(c.text), c.k); got != c.want {
			t.Errorf("%q at %d: %v, want %v", c.text, c.k, got, c.want)
		}
	}
}
