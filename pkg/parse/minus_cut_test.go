//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestMinusCut — a hyphen between an amount (a digit and a word of the unit table: 원, kg) and a digit is the minus that cuts; not a hyphen between two numbers, after a word that is no unit, after a unit with no digit before it, before a letter, at the sentence's end, in a span of code or inside a path, nor any other sign
package parse

import "testing"

func TestMinusCut(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text string
		at   int
		want bool
	}{
		{"3,000원-1,000원", 6, true},
		{"5kg-3kg", 3, true},
		{"10-20", 2, false},
		{"3가-1", 2, false},
		{"원-1", 1, false},
		{"3원-가", 2, false},
		{"3원-", 2, false},
		{"`3원-1`", 3, false},
		{"/a/3원-1", 5, false},
		{"3원+1", 2, false},
	} {
		c0 := newCutter(ix, c.text, false)
		if got := c0.minusCut(c.at); got != c.want {
			t.Errorf("%q at %d: %v, want %v", c.text, c.at, got, c.want)
		}
	}
}
