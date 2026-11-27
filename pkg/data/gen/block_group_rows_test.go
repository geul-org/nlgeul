//ff:func feature=data type=helper control=sequence
//ff:what TestBlockGroupRows — the lines of a block come in code order, and private use and surrogate blocks give none
package main

import "testing"

func TestBlockGroupRows(t *testing.T) {
	tbl := loadSample(t)
	jamo := jamoMap(tbl)
	rows := blockGroupRows(tbl, jamo, blockRange{Lo: 0x20, Hi: 0x7F, Name: "Basic Latin"})
	if len(rows) != 1 || rows[0][:1] != "." {
		t.Fatal(rows)
	}
	if blockGroupRows(tbl, jamo, blockRange{Lo: 0x20, Hi: 0x7F, Name: "Private Use Area"}) != nil {
		t.Fatal("private use")
	}
	if blockGroupRows(tbl, jamo, blockRange{Lo: 0x20, Hi: 0x7F, Name: "Low Surrogates"}) != nil {
		t.Fatal("surrogates")
	}
}
