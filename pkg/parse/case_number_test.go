//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCaseNumber — a court case number the unit begins with is not read inside (every place within it is inside), and the places after it are free (2008회합9 before 호로); a number with a greater number (17만2000) and a unit of Hangul alone are left as they are
package parse

import "testing"

func TestCaseNumber(t *testing.T) {
	sc := scanOf("2004다25123")
	for p := 1; p < 10; p++ {
		if !sc.inside[p] {
			t.Fatal(p, sc.inside)
		}
	}
	if sc.inside[10] {
		t.Fatal(sc.inside)
	}
	sc = scanOf("2008회합9호로")
	if !sc.inside[4] || !sc.inside[6] || sc.inside[7] || sc.inside[8] {
		t.Fatal(sc.inside)
	}
	sc = scanOf("17만2000")
	if sc.inside[2] || sc.inside[3] {
		t.Fatal(sc.inside)
	}
	sc = scanOf("사건번호")
	if sc.inside[1] || sc.inside[2] {
		t.Fatal(sc.inside)
	}
}
