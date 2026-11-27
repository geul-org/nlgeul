//ff:func feature=data type=helper control=sequence
//ff:what TestRowOf — a kept form gives a tab-separated line with its number, and an unassigned or control character, a plain one, an unkept form, a form with no row and a blank value give none
package main

import "testing"

func TestRowOf(t *testing.T) {
	tbl := loadSample(t)
	jamo := jamoMap(tbl)
	line, ok := rowOf(tbl, jamo, 0x2460, "Enclosed Alphanumerics")
	if !ok {
		t.Fatal("kept")
	}
	assertEq(t, line, "①\t원\t수\t1\t1")
	line, _ = rowOf(tbl, jamo, 0xBD, "Latin-1 Supplement")
	assertEq(t, line, "½\t분수\t수\t1/2\t0.5")
	if _, ok := rowOf(tbl, jamo, 0x0378, "X"); ok {
		t.Fatal("Cn")
	}
	if _, ok := rowOf(tbl, jamo, 0x09, "X"); ok {
		t.Fatal("Cc")
	}
	if _, ok := rowOf(tbl, jamo, 0x41, "X"); ok {
		t.Fatal("plain")
	}
	if _, ok := rowOf(tbl, jamo, 0x3131, "Hangul Compatibility Jamo"); ok {
		t.Fatal("form not kept")
	}
	if _, ok := rowOf(tbl, jamo, 0x1F12D, "Enclosed Alphanumeric Supplement"); ok {
		t.Fatal("no row")
	}
	tbl.Chars[0x2460] = charInfo{Category: "No", Decomp: "<circle> 0020"}
	if _, ok := rowOf(tbl, jamo, 0x2460, "X"); ok {
		t.Fatal("blank")
	}
}
