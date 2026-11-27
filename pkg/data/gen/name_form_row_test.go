//ff:func feature=data type=helper control=sequence
//ff:what TestNameFormRow — a number gives a number row, a Latin letter gives its capital or small letter, and a name with neither gives none
package main

import "testing"

func TestNameFormRow(t *testing.T) {
	tbl := loadSample(t)
	row, ok := nameFormRow(tbl, 0x2776, "검은 원")
	if !ok || row.Value != "1" || row.Kind != "수" {
		t.Fatal(row, ok)
	}
	row, ok = nameFormRow(tbl, 0x1F150, "검은 원")
	if !ok || row.Value != "A" || row.Kind != "로마자" {
		t.Fatal(row, ok)
	}
	row, ok = nameFormRow(tbl, 0x1F170, "검은 네모")
	if !ok || row.Value != "b" {
		t.Fatal(row, ok)
	}
	if _, ok := nameFormRow(tbl, 0x1F12D, "원"); ok {
		t.Fatal("no letter, no number")
	}
}
