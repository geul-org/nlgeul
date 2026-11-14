//ff:func feature=token type=helper control=sequence
//ff:what TestLatticePortions — the live states cut at the agreed places: each in the portion that holds where it begins, in the order they were made; a lattice no way reads has none
package token

import "testing"

func TestLatticePortions(t *testing.T) {
	sc := &scan{}
	root := sc.registry().root
	a := addState(sc, []*path{root}, wordStep(2, "a", 1, -1))
	b1 := addState(sc, []*path{a}, wordStep(3, "b", 1, -1))
	b2 := addState(sc, []*path{a}, wordStep(4, "c", 1, -1))
	c := addState(sc, []*path{b1}, wordStep(4, "d", 1, -1))
	dead := addState(sc, []*path{root}, wordStep(1, "x", 1, -1))
	lat := &lattice{n: 4, ends: []*path{c, b2}, all: sc.registry().all}
	got := lat.portions()
	if len(got) != 2 || got[0].from != 0 || got[0].to != 2 || len(got[0].states) != 1 || got[0].states[0] != a || dead == nil {
		t.Fatal(got)
	}
	if got[1].from != 2 || got[1].to != 4 || len(got[1].states) != 3 || got[1].states[0] != b1 || got[1].states[1] != b2 || got[1].states[2] != c {
		t.Fatal(got[1])
	}
	if (&lattice{n: 4}).portions() != nil {
		t.Fatal("no way")
	}
}
