//ff:func feature=token type=helper control=sequence
//ff:what TestLatticeReached — the states on a way that reads the whole unit: the states that end it and every one before them, found back by preds; a state that no end comes from and the root are not
package token

import "testing"

func TestLatticeReached(t *testing.T) {
	sc := &scan{}
	root := sc.registry().root
	a := addState(sc, []*path{root}, wordStep(2, "a", 1, -1))
	c := addState(sc, []*path{root}, wordStep(2, "c", 1, -1))
	b := addState(sc, []*path{a}, wordStep(3, "b", 1, -1))
	dead := addState(sc, []*path{c}, wordStep(3, "d", 1, 0))
	end := addState(sc, []*path{a, c}, wordStep(3, "e", 1, -1))
	lat := &lattice{n: 3, ends: []*path{end, b}, all: sc.registry().all}
	got := lat.reached()
	if len(got) != 4 || !got[a] || !got[b] || !got[c] || !got[end] || got[dead] || got[root] {
		t.Fatal(got)
	}
}
