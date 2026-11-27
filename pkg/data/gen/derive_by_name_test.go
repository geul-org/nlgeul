//ff:func feature=data type=helper control=sequence
//ff:what TestDeriveByName — the longest form key wins, a key outside the enclosed blocks is not found, a name with no key is not found, and found without a row is not ok
package main

import "testing"

func TestDeriveByName(t *testing.T) {
	tbl := loadSample(t)
	row, found, ok := deriveByName(tbl, 0x2776, "Dingbats")
	if !found || !ok || row.Form != "검은 원" {
		t.Fatal(row, found, ok)
	}
	if _, found, _ := deriveByName(tbl, 0x2776, "Number Forms"); found {
		t.Fatal("block")
	}
	if _, found, _ := deriveByName(tbl, 0x41, "Dingbats"); found {
		t.Fatal("no key")
	}
	if _, found, ok := deriveByName(tbl, 0x1F12D, "Enclosed Alphanumeric Supplement"); !found || ok {
		t.Fatal(found, ok)
	}
}
