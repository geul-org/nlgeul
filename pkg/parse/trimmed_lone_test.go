//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestTrimmedLone — a unit of signs alone that begins with a path's head that stands with no name has that head read (~/, ./, ../, also before other signs); a slash alone and other signs alone have nothing read
package parse

import "testing"

func TestTrimmedLone(t *testing.T) {
	cases := []struct {
		in       string
		from, to int
	}{
		{"~/", 0, 2}, {"./", 0, 2}, {"../", 0, 3}, {"~/,", 0, 2}, {"/", 0, 0}, {"//", 0, 0}, {"./..", 0, 0}, {"->", 0, 0},
	}
	for _, c := range cases {
		if from, to := trimmed(c.in); from != c.from || to != c.to {
			t.Errorf("%q: %d %d", c.in, from, to)
		}
	}
}
