//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestClosesPair — candidates of a sign of a pair close a pair open before it when one is a closing at a depth of 1 or more; a closing at depth 0 (none of its kind open), openings alone and no candidate do not
package parse

import "testing"

func TestClosesPair(t *testing.T) {
	cases := []struct {
		in   []pairing
		want bool
	}{
		{[]pairing{{name: shutName, depth: 1, verdict: 1}}, true},
		{[]pairing{{name: openName, depth: 2, verdict: 1}, {name: shutName, depth: 2, verdict: 1}}, true},
		{[]pairing{{name: shutName, depth: 0, verdict: 1}}, false},
		{[]pairing{{name: openName, depth: 1, verdict: 1}}, false},
		{nil, false},
	}
	for _, c := range cases {
		if got := closesPair(c.in); got != c.want {
			t.Errorf("closesPair(%v) = %v", c.in, got)
		}
	}
}
