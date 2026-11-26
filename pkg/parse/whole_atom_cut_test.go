//ff:func feature=parse type=reader control=sequence
//ff:what TestWholeAtomCut — a unit that begins with a path that ends in a slash before Hangul notes where the path is cut (the atom is the path up to the slash); a path with a particle after a name's letter and a path alone note none; a unit that begins with none of the atoms notes where its letters end as a path with no head and one slash, when they are written so
package parse

import "testing"

func TestWholeAtomCut(t *testing.T) {
	sc := scanOf("~/a/b/회의록")
	if a := sc.atoms[0]; a.end != 6 || sc.cut != 6 || sc.weak != 0 || sc.inside[6] || sc.inside[7] {
		t.Fatal(sc.atoms, sc.cut, sc.weak)
	}
	if sc = scanOf("~/a/b를"); sc.cut != 0 || sc.weak != 0 {
		t.Fatal(sc.cut, sc.weak)
	}
	if sc = scanOf("~/a/b/"); sc.cut != 0 {
		t.Fatal(sc.cut)
	}
	if sc = scanOf("notes/회의-결과"); sc.cut != 0 || sc.weak != 11 || sc.inside[6] || sc.inside[9] {
		t.Fatal(sc.cut, sc.weak, sc.inside)
	}
	if sc = scanOf("고양이가"); sc.cut != 0 || sc.weak != 0 {
		t.Fatal(sc.cut, sc.weak)
	}
}
