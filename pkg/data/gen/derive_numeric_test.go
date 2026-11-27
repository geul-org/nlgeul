//ff:func feature=data type=helper control=sequence
//ff:what TestDeriveNumeric — a number character in a number block is a Roman numeral or a number form, and one outside the blocks, with no number, or not No or Nl gives none
package main

import "testing"

func TestDeriveNumeric(t *testing.T) {
	tbl := loadSample(t)
	row, ok := deriveNumeric(tbl, 0x2185, "Number Forms")
	if !ok || row.Form != "로마 숫자" || row.Value != "6" {
		t.Fatal(row, ok)
	}
	row, ok = deriveNumeric(tbl, 0x24EA, "Dingbats")
	if !ok || row.Form != "수 꼴" || row.Value != "0" {
		t.Fatal(row, ok)
	}
	if _, ok := deriveNumeric(tbl, 0x2185, "Basic Latin"); ok {
		t.Fatal("block")
	}
	if _, ok := deriveNumeric(tbl, 0x41, "Number Forms"); ok {
		t.Fatal("no number")
	}
	if _, ok := deriveNumeric(tbl, 0x31, "Number Forms"); ok {
		t.Fatal("Nd")
	}
}
