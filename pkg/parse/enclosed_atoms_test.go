//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestEnclosedAtoms — each circled or parenthesized number is an atom of one letter named 수 with its end kept, also inside a word and next to another; circled Hangul is none
package parse

import "testing"

func TestEnclosedAtoms(t *testing.T) {
	sc := scanOf("사실⑤ ①②㊿ ㉠")
	for _, i := range []int{2, 4, 5, 6} {
		if a := sc.atoms[i]; a.end != i+1 || atomNames(a) != numName || !a.num || !sc.numEnds[i+1] {
			t.Fatal(i, sc.atoms)
		}
	}
	if len(sc.atoms) != 4 || sc.numEnds[9] || sc.numEnds[8] {
		t.Fatal(sc.atoms)
	}
	sc = scanOf("⑴3")
	if sc.atoms[0].end != 1 || sc.atoms[1].end != 2 || atomNames(sc.atoms[1]) != numName {
		t.Fatal(sc.atoms)
	}
}
