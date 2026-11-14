//ff:func feature=token type=reader control=sequence
//ff:what TestKeepAlive — the states of a list that stay, in the order given, in a list of their own
package token

import "testing"

func TestKeepAlive(t *testing.T) {
	a, b, c := &path{}, &path{}, &path{}
	in := []*path{a, b, c}
	got := keepAlive(in, map[*path]bool{a: true, c: true})
	if len(got) != 2 || got[0] != a || got[1] != c || len(in) != 3 {
		t.Fatal(got)
	}
}
