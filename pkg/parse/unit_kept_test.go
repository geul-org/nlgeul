//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestUnitKept — a sign of the unit table right before a core that begins with 당 is kept (㎡당, %당, after another sign too); a sign that is no unit, a unit sign before other letters, and a core with nothing before it stay as they began
package parse

import "testing"

func TestUnitKept(t *testing.T) {
	cases := []struct {
		in         string
		from, want int
	}{
		{"㎡당", 1, 0},
		{"%당", 1, 0},
		{"#㎡당", 2, 1},
		{"?당", 1, 1},
		{"㎡곰", 1, 1},
		{"당", 0, 0},
		{"곰당", 1, 1},
	}
	for _, c := range cases {
		if got := unitKept([]rune(c.in), c.from); got != c.want {
			t.Errorf("unitKept(%q, %d) = %d, want %d", c.in, c.from, got, c.want)
		}
	}
}
