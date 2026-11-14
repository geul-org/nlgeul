//ff:func feature=token type=helper control=sequence
//ff:what TestRegistryKeep — a state kept again after a cut is noted in the registry: in the order made, by its kind of word, with each state before it, and its word with the number of birth it has, printed the same once
package token

import "testing"

func TestRegistryKeep(t *testing.T) {
	sc := scanOf("가나다")
	root := sc.registry().root
	a := addState(sc, []*path{root}, wordStep(1, "a", 1, 0))
	b := addState(sc, []*path{a}, wordStep(2, "b", 1, 0))
	b.last.born = 7
	reg := newRegistry(0)
	reg.keep(a)
	reg.keep(b)
	if len(reg.all) != 2 || reg.all[1] != b || reg.states[stateKeyOf(b)] != b || !reg.preds[predKey{b, a}] || reg.born[wordKeyOf(b)] != 7 || !reg.lone[wordKeyOf(b)] {
		t.Fatal(reg)
	}
}
