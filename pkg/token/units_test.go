//ff:func feature=token type=reader control=sequence
//ff:what TestUnits — 개 right after a number is a unit; 번째 is a unit and an order, 여개 more-than and a unit; a longer unit and none for the shorter that does not touch the number; a unit after a run that is no number, after a syllable, out of order, or at the start is not read; Hangul numerals and a letter of the table are numbers before a unit and name the stem; a greater number follows digits only, and 조 after a number with 제 stands lower than the unit 조; a letter of the order (가) is a number before 목 only; a unit with letters running on is not read where nothing is known to follow
package token

import "testing"

func TestUnits(t *testing.T) {
	got := scanOf("3개").units(2, 0, true)
	if len(got) != 1 || got[0].start != 1 || got[0].pieces[0].names != "단위:개" || got[0].pieces[0].verdict != 1 {
		t.Fatal(got)
	}
	got = scanOf("3번째").units(3, 0, true)
	if len(got) != 1 || got[0].start != 1 || len(got[0].pieces) != 2 || got[0].pieces[1].names != orderName {
		t.Fatal(got)
	}
	got = scanOf("10여개").units(4, 0, true)
	if len(got) != 1 || got[0].start != 2 || got[0].pieces[0].names != overName {
		t.Fatal(got)
	}
	got = scanOf("3개월").units(3, 0, true)
	if len(got) != 1 || got[0].pieces[0].names != "단위:개월" {
		t.Fatal(got)
	}
	if got = scanOf("3째").units(2, 0, true); len(got) != 1 {
		t.Fatal("an order right after a number", got)
	}
	if len(scanOf("L2개").units(3, 0, true)) != 0 || len(scanOf("시대").units(2, 0, true)) != 0 || len(scanOf("3째개").units(3, 0, true)) != 0 || len(scanOf("개").units(1, 0, true)) != 0 || len(scanOf("3개").units(2, placeUnit, true)) != 0 {
		t.Fatal("read where it should not be")
	}
	got = scanOf("세명").units(2, 0, true)
	if len(got) != 1 || got[0].only != "세" || got[0].as != numName || got[0].weight >= 0 {
		t.Fatal(got)
	}
	if got = scanOf("k번째").units(3, 0, true); len(got) != 1 || got[0].only != "k" || len(got[0].pieces) != 2 {
		t.Fatal(got)
	}
	got = scanOf("가목").units(2, 0, true)
	if len(got) != 1 || got[0].only != "가" || got[0].as != numName || got[0].pieces[0].names != "단위:목" || got[0].weight >= 0 {
		t.Fatal("a letter of the order before 목 is a number", got)
	}
	if len(scanOf("가개").units(2, 0, true)) != 0 || len(scanOf("다목적").units(2, 0, true)) != 0 {
		t.Fatal("a letter of the order is a number before 목 only, and not with letters running on")
	}
	got = scanOf("제36조").units(4, 0, true)
	if len(got) != 2 || got[0].pieces[0].names != "큰 수:조" || got[0].pieces[0].verdict <= 0 || got[0].pieces[0].verdict >= 0.5 || got[1].pieces[0].names != "단위:조" || got[1].pieces[0].verdict != 1 {
		t.Fatal("after a number with 제 the greater number 조 stands lower than the unit 조", got)
	}
	got = scanOf("36조").units(3, 0, true)
	if len(got) != 2 || got[0].pieces[0].names != "큰 수:조" || got[0].pieces[0].verdict != 1 || got[1].pieces[0].verdict != 1 {
		t.Fatal("with no 제 the greater number 조 stands as the unit does", got)
	}
	got = scanOf("3천만원").units(4, 0, true)
	if len(got) != 3 || got[0].only != "만" || got[1].only != "천만" || len(got[2].pieces) != 3 || got[2].pieces[0].names != "큰 수:천" {
		t.Fatal(got)
	}
	if len(scanOf("삼만").units(2, 0, true)) != 0 {
		t.Fatal("a greater number after a Hangul numeral is part of the numeral")
	}
	if got := wordsScan("3개고양이", layerCost).units(2, 0, true); len(got) != 1 {
		t.Fatal("with a word table a unit after digits ends its word though letters run on", got)
	}
	if got := wordsScan("세개고양이", layerCost).units(2, 0, true); len(got) != 0 {
		t.Fatal("not after a Hangul numeral", got)
	}
	if got := wordsScan("백원짜리", layerCost).units(4, 0, true); len(got) != 1 || len(got[0].pieces) != 2 || got[0].pieces[1].names != "접미:짜리" {
		t.Fatal("a suffix after a unit", got)
	}
	if len(scanOf("두배적").units(2, 0, true)) != 0 || len(scanOf("두배적").units(2, 0, false)) != 1 || len(scanOf("3개.").units(2, 0, true)) != 1 {
		t.Fatal("letters running on")
	}
}
