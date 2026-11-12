//ff:func feature=token type=reader control=sequence
//ff:what TestWholeAtom — a path at a unit's head is one atom named so, with no place inside it for a word to begin; a word is none; the sign of a pair at a unit's head is the atom, named as the unit names it
package token

import "testing"

func TestWholeAtom(t *testing.T) {
	sc := scanOf("~/a/b를")
	if a := sc.atoms[0]; a.end != 5 || a.names != "경로" || !sc.inside[3] || sc.inside[5] {
		t.Fatal(sc.atoms, sc.inside)
	}
	if end := scanOf("고양이").wholeAtom(); end != 0 {
		t.Fatal(end)
	}
	head := &scan{ix: sc.ix, s: []rune(")를"), head: 1, headName: pairName, inside: make([]bool, 3), atoms: map[int]atom{}}
	if a := head.atoms[head.wholeAtom()-1]; a.end != 1 || a.names != pairName {
		t.Fatal("the sign of a pair at a unit's head", head.atoms)
	}
}
