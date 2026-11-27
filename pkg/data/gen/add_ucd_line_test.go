//ff:func feature=data type=helper control=sequence
//ff:what TestAddUCDLine — a plain line is stored, First is remembered, Last closes the range, a name in angle brackets is dropped, and bad lines are errors
package main

import "testing"

func TestAddUCDLine(t *testing.T) {
	tbl := &ucdTable{Chars: map[rune]charInfo{}}
	var first rune
	if err := addUCDLine(tbl, "0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;", &first); err != nil {
		t.Fatal(err)
	}
	assertEq(t, tbl.Chars[0x41].Name, "LATIN CAPITAL LETTER A")
	assertEq(t, tbl.Chars[0x41].Category, "Lu")
	if err := addUCDLine(tbl, "AC00;<Hangul Syllable, First>;Lo;0;L;;;;;N;;;;;", &first); err != nil || first != 0xAC00 {
		t.Fatal(first, err)
	}
	if err := addUCDLine(tbl, "D7A3;<Hangul Syllable, Last>;Lo;0;L;;;;;N;;;;;", &first); err != nil {
		t.Fatal(err)
	}
	assertEq(t, tbl.Ranges[0].Category, "Lo")
	if err := addUCDLine(tbl, "0001;<control>;Cc;0;BN;;;;;N;;;;;", &first); err != nil {
		t.Fatal(err)
	}
	assertEq(t, tbl.Chars[1].Name, "")
	if err := addUCDLine(tbl, "XYZ;A;Lu;0;L;;;;;N;;;;;", &first); err == nil {
		t.Fatal("bad code")
	}
	if err := addUCDLine(tbl, "0041;A", &first); err == nil {
		t.Fatal("short")
	}
}
