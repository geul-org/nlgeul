//ff:func feature=token type=reader control=sequence
//ff:what TestTop — the best readings in order, the same reading once, k at most
package token

import "testing"

func TestTop(t *testing.T) {
	ix := scanOf("x").ix
	a := &path{sig: "a", verdict: 0.3, pref: -1}
	b := &path{sig: "b", verdict: 1, pref: -9}
	c := &path{sig: "c", verdict: 0.3, pref: -2}
	got := ix.top([]*path{a, c, b, a}, 2)
	if len(got) != 2 || got[0] != b || got[1] != a {
		t.Fatal(got)
	}
}
