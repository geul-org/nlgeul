//ff:func feature=token type=reader control=sequence
//ff:what TestUnitsDigit — with no word table, a unit sign with a digit running on after it is read (5'10: feet), as it is with letters running on only when a word table is given; a unit sign with a letter running on is not read
package token

import "testing"

func TestUnitsDigit(t *testing.T) {
	got := scanOf("5'10").units(2, 0, true)
	if len(got) == 0 || got[0].start != 1 || got[0].pieces[0].names != "단위:피트" {
		t.Fatal(got)
	}
	if got = scanOf("5'a").units(2, 0, true); len(got) != 0 {
		t.Fatal("a letter runs on", got)
	}
	if got = scanOf("5'").units(2, 0, true); len(got) != 2 || got[0].pieces[0].names != "단위:피트" || got[1].pieces[0].names != "단위:분" {
		t.Fatal("the lines of the table", got)
	}
}
