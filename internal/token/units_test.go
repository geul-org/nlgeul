//ff:func feature=token type=reader control=sequence
//ff:what TestUnits — 개 right after a number is a unit; 번째 is a unit and an order, 여개 more-than and a unit; a longer unit and none for the shorter that does not touch the number; a unit after a run that is no number, after a syllable, out of order, or at the start is not read
package token

import "testing"

func TestUnits(t *testing.T) {
	got := scanOf("3개").units(2, 0)
	if len(got) != 1 || got[0].start != 1 || got[0].pieces[0].names != "단위:개" || got[0].pieces[0].verdict != 1 {
		t.Fatal(got)
	}
	got = scanOf("3번째").units(3, 0)
	if len(got) != 1 || got[0].start != 1 || len(got[0].pieces) != 2 || got[0].pieces[1].names != orderName {
		t.Fatal(got)
	}
	got = scanOf("10여개").units(4, 0)
	if len(got) != 1 || got[0].start != 2 || got[0].pieces[0].names != overName {
		t.Fatal(got)
	}
	got = scanOf("3개월").units(3, 0)
	if len(got) != 1 || got[0].pieces[0].names != "단위:개월" {
		t.Fatal(got)
	}
	if got = scanOf("3째").units(2, 0); len(got) != 1 {
		t.Fatal("an order right after a number", got)
	}
	if len(scanOf("L2개").units(3, 0)) != 0 || len(scanOf("시대").units(2, 0)) != 0 || len(scanOf("3째개").units(3, 0)) != 0 || len(scanOf("개").units(1, 0)) != 0 || len(scanOf("3개").units(2, placeUnit)) != 0 {
		t.Fatal("read where it should not be")
	}
}
