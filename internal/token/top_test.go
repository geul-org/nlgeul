//ff:func feature=token type=reader control=sequence
//ff:what TestTop — the best readings in order, the same reading once, none left out
package token

import "testing"

func TestTop(t *testing.T) {
	ix := scanOf("x").ix
	a := &path{sig: "a", verdict: 0.3, pref: -1}
	b := &path{sig: "b", verdict: 1, pref: -9}
	c := &path{sig: "c", verdict: 0.3, pref: -2}
	got := ix.top([]*path{a, c, b, a})
	if len(got) != 3 || got[0] != b || got[1] != a || got[2] != c {
		t.Fatal(got)
	}
}
