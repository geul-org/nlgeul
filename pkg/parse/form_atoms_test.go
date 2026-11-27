//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestFormAtoms — each letter that writes a number in a shape of its own is an atom of one letter named 수 with its end kept, also inside a word and next to another (⑤, ①, ②, ㊿, ½); a letter of the tables that is no sign and writes no number is an atom with the table's name and no number (㊀: 원:한자); a sign of the tables (㉠) and a Roman numeral letter alone make none here
package parse

import "testing"

func TestFormAtoms(t *testing.T) {
	sc := scanOf("사실⑤ ①②㊿ ㉠")
	if half := scanOf("½"); !half.atoms[0].num || !half.numEnds[1] {
		t.Fatal(half.atoms)
	}
	if one := scanOf("㊀"); one.atoms[0].end != 1 || one.atoms[0].num || atomNames(one.atoms[0]) != "원:한자" || one.numEnds[1] {
		t.Fatal(one.atoms)
	}
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
