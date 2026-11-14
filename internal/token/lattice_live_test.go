//ff:func feature=token type=helper control=sequence
//ff:what TestLatticeLive — the states on a way that reads the whole unit, in the order they were made
package token

import "testing"

func TestLatticeLive(t *testing.T) {
	sc := &scan{}
	root := sc.registry().root
	a := addState(sc, []*path{root}, wordStep(1, "a", 1, -1))
	dead := addState(sc, []*path{root}, wordStep(2, "d", 1, -1))
	b := addState(sc, []*path{a}, wordStep(3, "b", 1, -1))
	lat := &lattice{n: 3, ends: []*path{b}, all: sc.registry().all}
	got := lat.live()
	if len(got) != 2 || got[0] != a || got[1] != b || dead == nil {
		t.Fatal(got)
	}
}
