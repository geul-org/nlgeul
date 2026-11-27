//ff:func feature=data type=helper control=sequence
//ff:what TestDerive — a tag gives a tagged row, a presentation tag falls through, an enclosed name gives its row or none, and a number form gives its row
package main

import "testing"

func TestDerive(t *testing.T) {
	tbl := loadSample(t)
	jamo := jamoMap(tbl)
	if row, ok := derive(tbl, jamo, 0x2460, "Enclosed Alphanumerics"); !ok || row.Form != "원" {
		t.Fatal(row, ok)
	}
	if _, ok := derive(tbl, jamo, 0xFB50, "Arabic Presentation Forms-A"); ok {
		t.Fatal("presentation")
	}
	if row, ok := derive(tbl, jamo, 0x2776, "Dingbats"); !ok || row.Form != "검은 원" {
		t.Fatal(row, ok)
	}
	tbl.Chars[0x1F12D] = charInfo{Name: "CIRCLED CD", Category: "So"}
	if _, ok := derive(tbl, jamo, 0x1F12D, "Enclosed Alphanumeric Supplement"); ok {
		t.Fatal("found but no row")
	}
	if row, ok := derive(tbl, jamo, 0x2185, "Number Forms"); !ok || row.Form != "로마 숫자" {
		t.Fatal(row, ok)
	}
}
