//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPlateNumber — a registration number the unit begins with is not read inside and is one atom named 차등록 from the unit's head to its end (also with the name of a region before it), and the places after it are free (74모2238 before 호를); where the letters are a court case number too the atom is named 사건 first and 차등록 after it (97도2836); a case number that is no registration number keeps its one name, and a unit with no such number is left as it is
package parse

import "testing"

func TestPlateNumber(t *testing.T) {
	sc := scanOf("306너8412")
	for p := 1; p < 8; p++ {
		if !sc.inside[p] {
			t.Fatal(p, sc.inside)
		}
	}
	if a := sc.atoms[0]; a.end != 8 || len(a.names) != 1 || a.names[0] != plateKind || a.num {
		t.Fatal("the registration number is one atom named 차등록", a)
	}
	sc = scanOf("서울34가5678호를")
	if a := sc.atoms[0]; a.end != 9 || len(a.names) != 1 || a.names[0] != plateKind {
		t.Fatal("the region is in the atom", a)
	}
	if !sc.inside[2] || !sc.inside[8] || sc.inside[9] || sc.inside[10] {
		t.Fatal(sc.inside)
	}
	sc = scanOf("97도2836")
	if a := sc.atoms[0]; a.end != 7 || len(a.names) != 2 || a.names[0] != caseKind || a.names[1] != plateKind {
		t.Fatal("a case number first, a registration number after it", a)
	}
	sc = scanOf("2004다25123")
	if a := sc.atoms[0]; a.end != 10 || len(a.names) != 1 || a.names[0] != caseKind {
		t.Fatal("a case number alone", a)
	}
	sc = scanOf("등록번호")
	if _, ok := sc.atoms[0]; ok || sc.inside[1] {
		t.Fatal(sc.atoms, sc.inside)
	}
}
