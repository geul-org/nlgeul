//ff:func feature=token type=reader control=sequence
//ff:what TestWholeAtom — a path at a unit's head is one atom named so, with no place inside it for a word to begin; a word is none
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
}
