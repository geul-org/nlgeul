//ff:func feature=token type=helper control=sequence
//ff:what TestLatticeAgreed — a place no live state lies across is one every way agrees on, with 0 first and the unit's end last; a place a word of one letter does not cross, and one a longer word does, are told apart; with no live state there is none
package token

import "testing"

func TestLatticeAgreed(t *testing.T) {
	sc := &scan{}
	root := sc.registry().root
	a := addState(sc, []*path{root}, wordStep(2, "a", 1, -1))
	b := addState(sc, []*path{a}, wordStep(6, "b", 1, -1))
	lat := &lattice{n: 6, ends: []*path{b}, all: sc.registry().all}
	if got := lat.agreed(lat.live()); len(got) != 3 || got[0] != 0 || got[1] != 2 || got[2] != 6 {
		t.Fatal(got)
	}
	c := addState(sc, []*path{root}, wordStep(3, "c", 1, -1))
	d := addState(sc, []*path{c}, wordStep(6, "d", 1, -1))
	lat = &lattice{n: 6, ends: []*path{b, d}, all: sc.registry().all}
	if got := lat.agreed(lat.live()); len(got) != 2 || got[0] != 0 || got[1] != 6 {
		t.Fatal("two ways cross each other's place", got)
	}
	e := addState(sc, []*path{root}, wordStep(1, "e", 1, -1))
	f := addState(sc, []*path{e}, wordStep(2, "f", 1, -1))
	g := addState(sc, []*path{f}, wordStep(6, "g", 1, -1))
	lat = &lattice{n: 6, ends: []*path{g}, all: sc.registry().all}
	if got := lat.agreed(lat.live()); len(got) != 4 || got[1] != 1 || got[2] != 2 {
		t.Fatal("words of one letter", got)
	}
	if got := (&lattice{n: 6}).agreed(nil); got != nil {
		t.Fatal("no way", got)
	}
}
