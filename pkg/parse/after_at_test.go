//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAfterAt — whether a piece read after a number begins at a place: a unit of one letter (월) or more (개월) and a greater number do, a particle, a digit and the unit's end do not
package parse

import "testing"

func TestAfterAt(t *testing.T) {
	cases := []struct {
		text string
		at   int
		want bool
	}{
		{"3월", 1, true},
		{"3개월", 1, true},
		{"5억", 1, true},
		{"3은", 1, false},
		{"3월", 0, false},
		{"3월", 2, false},
	}
	for _, c := range cases {
		sc := &scan{ix: scanOf("").ix, s: []rune(c.text)}
		if got := sc.afterAt(c.at); got != c.want {
			t.Errorf("%s at %d: %v, want %v", c.text, c.at, got, c.want)
		}
	}
}
