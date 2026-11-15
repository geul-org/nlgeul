//ff:func feature=token type=reader control=sequence
//ff:what TestCloseAtom — a run of digits is named a number and its end kept; a number with a sign at its front begins at the sign, and so does one with 제 right before it at a word's head (named 차례 too); 제 after a syllable is left out; a run that is no number keeps no name
package token

import "testing"

func TestCloseAtom(t *testing.T) {
	sc := scanOf("-30 ab")
	if a := sc.atoms[0]; a.end != 3 || atomNames(a) != "부호,수" || !a.num || !sc.numEnds[3] || !sc.inside[1] || !sc.inside[2] {
		t.Fatal(sc.atoms)
	}
	if a := sc.atoms[4]; a.end != 6 || atomNames(a) != "" || a.num || sc.numEnds[6] {
		t.Fatal(sc.atoms)
	}
	sc = scanOf("제1항")
	if a := sc.atoms[0]; a.end != 2 || atomNames(a) != "차례,수" || !a.num || !sc.numEnds[2] || !sc.inside[1] {
		t.Fatal(sc.atoms)
	}
	sc = scanOf("문제1")
	if a, ok := sc.atoms[2]; !ok || a.end != 3 || atomNames(a) != numName {
		t.Fatal(sc.atoms)
	}
	sc = scanOf("1,000")
	if a := sc.atoms[0]; a.end != 5 || atomNames(a) != numName || !a.num {
		t.Fatal(sc.atoms)
	}
}
