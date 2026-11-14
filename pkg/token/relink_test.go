//ff:func feature=token type=helper control=sequence
//ff:what TestRelink — a state that stays has before it only what stays: the seed of each root for a state that begins at the cut, the states that stay for one that begins after it; the first is the state before it
package token

import "testing"

func TestRelink(t *testing.T) {
	sc := scanOf("가나다라")
	root := sc.registry().root
	x := addState(sc, []*path{root}, wordStep(1, "x", 1, 0))
	y := addState(sc, []*path{root}, wordStep(1, "y", 1, 0))
	seed := &path{last: x.last}
	a := addState(sc, []*path{x, y}, wordStep(2, "a", 1, 0))
	relink(a, 1, map[*path]bool{x: true}, nil, map[*path]*path{x: seed})
	if len(a.preds) != 1 || a.preds[0] != seed || a.prev != seed {
		t.Fatal(a.preds)
	}
	b := addState(sc, []*path{a}, wordStep(3, "b", 1, 0))
	c := addState(sc, []*path{a}, wordStep(3, "c", 1, 0))
	d := addState(sc, []*path{b, c}, wordStep(4, "d", 1, 0))
	relink(d, 1, nil, map[*path]bool{c: true}, nil)
	if len(d.preds) != 1 || d.preds[0] != c || d.prev != c {
		t.Fatal(d.preds)
	}
}
