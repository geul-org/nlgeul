//ff:func feature=data type=helper control=sequence
//ff:what TestBlockRows — the lines of a block come in code order, and private use and surrogate blocks give none
package main

import "testing"

func TestBlockRows(t *testing.T) {
	tbl := loadSample(t)
	jamo := jamoMap(tbl)
	rows := blockRows(tbl, jamo, blockRange{Lo: 0x2460, Hi: 0x2488, Name: "Enclosed Alphanumerics"})
	if len(rows) != 3 || rows[0][:3] != "①" {
		t.Fatal(rows)
	}
	if blockRows(tbl, jamo, blockRange{Lo: 0x2460, Hi: 0x2488, Name: "Private Use Area"}) != nil {
		t.Fatal("private use")
	}
	if blockRows(tbl, jamo, blockRange{Lo: 0x2460, Hi: 0x2488, Name: "High Surrogates"}) != nil {
		t.Fatal("surrogates")
	}
}
