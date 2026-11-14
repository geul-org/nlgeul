//ff:func feature=token type=reader control=sequence
//ff:what TestCloseAtom — a run of digits is named a number and its end kept; a number with a sign at its front begins at the sign; a run that is no number keeps no name
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
	sc = scanOf("1,000")
	if a := sc.atoms[0]; a.end != 5 || atomNames(a) != numName || !a.num {
		t.Fatal(sc.atoms)
	}
}
