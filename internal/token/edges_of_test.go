//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestEdgesOf — the signs before and after a unit's read letters and where the latter begin: a plain unit is cut by trimmed; a unit headed by a sign of a pair or a span of code has none before, and what stands after its read letters; a unit of signs only is all after, from 0
package token

import "testing"

func TestEdgesOf(t *testing.T) {
	cases := []struct {
		u           unit
		core        string
		lead, trail string
		at          int
	}{
		{unit{text: "#태그?"}, "태그", "#", "?", 3},
		{unit{text: "곰"}, "곰", "", "", 1},
		{unit{text: ")이다.", head: 1, name: pairName}, ")이다", "", ".", 3},
		{unit{text: ")", head: 1, name: pairName}, ")", "", "", 1},
		{unit{text: "`a`를,", head: 3, name: codeName}, "`a`를", "", ",", 4},
		{unit{text: "..."}, "", "", "...", 0},
	}
	for _, c := range cases {
		lead, trail, at := edgesOf(c.u, []rune(c.core))
		if string(lead) != c.lead || string(trail) != c.trail || at != c.at {
			t.Errorf("%q: before %q after %q at %d, want %q %q %d", c.u.text, string(lead), string(trail), at, c.lead, c.trail, c.at)
		}
	}
}
