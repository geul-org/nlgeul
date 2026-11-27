//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestUnitScan — the scan of a unit holds the letters that are read and the unit's head, and is told closed only when a sign that ends a sentence stands right after those letters, and called only when a comma or an exclamation mark does, and bracketed only when the unit stands alone in a round bracket written onto a word with no sign trimmed after its letters
package parse

import "testing"

func TestUnitScan(t *testing.T) {
	ix, _ := newIndex(Options{})
	cases := []struct {
		u      unit
		core   string
		closed bool
		called bool
	}{
		{unit{text: "문제야."}, "문제야", true, false},
		{unit{text: "문제야?"}, "문제야", true, false},
		{unit{text: "문제야!"}, "문제야", true, true},
		{unit{text: "문제야"}, "문제야", false, false},
		{unit{text: "문제야..."}, "문제야", false, false},
		{unit{text: "문제야,"}, "문제야", false, true},
		{unit{text: ")야.", head: 1, name: pairName}, ")야", true, false},
		{unit{text: "?"}, "", false, false},
	}
	for _, c := range []struct {
		u    unit
		want bool
	}{{unit{text: "III", bracketed: true}, true}, {unit{text: "III"}, false}, {unit{text: "III,", bracketed: true}, false}} {
		if sc := unitScan(ix, c.u); sc.bracketed != c.want {
			t.Errorf("unitScan(%q): bracketed %v", c.u.text, sc.bracketed)
		}
	}
	for _, c := range cases {
		sc := unitScan(ix, c.u)
		if sc.ix != ix || string(sc.s) != c.core || sc.closed != c.closed || sc.called != c.called || sc.head != c.u.head || sc.headName != c.u.name {
			t.Errorf("unitScan(%q): core %q closed %v called %v head %d %q", c.u.text, string(sc.s), sc.closed, sc.called, sc.head, sc.headName)
		}
	}
}
