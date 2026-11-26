//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestJointAt — a joint to a following number stands where the particle 의 or the prefix 제 has a digit right after it (31조의4, 36조제1항), not where no digit follows (31조의 뜻), where other letters stand (31조3항), or past the unit's end
package parse

import "testing"

func TestJointAt(t *testing.T) {
	cases := []struct {
		text string
		at   int
		want bool
	}{
		{"31조의4", 3, true},
		{"36조제1항", 3, true},
		{"31조의", 3, false},
		{"31조의뜻", 3, false},
		{"31조3항", 3, false},
		{"31조", 3, false},
	}
	for _, c := range cases {
		sc := &scan{ix: scanOf("").ix, s: []rune(c.text)}
		if got := sc.jointAt(c.at); got != c.want {
			t.Errorf("%s at %d: %v, want %v", c.text, c.at, got, c.want)
		}
	}
}
