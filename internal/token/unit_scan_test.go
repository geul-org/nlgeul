//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestUnitScan — the scan of a unit holds the letters that are read and the unit's head, and is told closed only when a sign that ends a sentence stands right after those letters
package token

import "testing"

func TestUnitScan(t *testing.T) {
	ix, _ := newIndex(Options{})
	cases := []struct {
		u      unit
		core   string
		closed bool
	}{
		{unit{text: "문제야."}, "문제야", true},
		{unit{text: "문제야?"}, "문제야", true},
		{unit{text: "문제야!"}, "문제야", true},
		{unit{text: "문제야"}, "문제야", false},
		{unit{text: "문제야..."}, "문제야", false},
		{unit{text: "문제야,"}, "문제야", false},
		{unit{text: ")야.", head: 1, name: "닫는 묶음표"}, ")야", true},
		{unit{text: "?"}, "", false},
	}
	for _, c := range cases {
		sc := unitScan(ix, c.u)
		if sc.ix != ix || string(sc.s) != c.core || sc.closed != c.closed || sc.head != c.u.head || sc.headName != c.u.name {
			t.Errorf("unitScan(%q): core %q closed %v head %d %q", c.u.text, string(sc.s), sc.closed, sc.head, sc.headName)
		}
	}
}
