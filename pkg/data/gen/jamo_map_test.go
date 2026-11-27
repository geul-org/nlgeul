//ff:func feature=data type=helper control=sequence
//ff:what TestJamoMap — a compatibility jamo is found by the jamo it decomposes to; jamo with no one-character decomposition are left out
package main

import "testing"

func TestJamoMap(t *testing.T) {
	tbl := loadSample(t)
	tbl.Chars[0x3132] = charInfo{Decomp: "<compat> 1100 1100"}
	tbl.Chars[0x3133] = charInfo{Decomp: "11A8"}
	m := jamoMap(tbl)
	if m[0x1100] != 0x3131 || m[0x1161] != 0x314F || len(m) != 2 {
		t.Fatal(m)
	}
}
