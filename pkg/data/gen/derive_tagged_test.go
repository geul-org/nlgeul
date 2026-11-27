//ff:func feature=data type=helper control=iteration dimension=1
//ff:what TestDeriveTagged — the form comes from the tag, a compat decomposition is read by its shape, jamo are joined, the fraction slash becomes /, Roman numerals and fractions hold numbers, an unknown tag is its own form
package main

import "testing"

func TestDeriveTagged(t *testing.T) {
	tbl := loadSample(t)
	jamo := jamoMap(tbl)
	cases := []struct {
		r    rune
		want string
	}{
		{0x2460, "원|수|1"}, {0x2474, "괄호|수|1"}, {0x2488, "점 붙음|수|1"},
		{0x3200, "괄호|한글|ㄱ"}, {0x320E, "괄호|한글|가"}, {0xBD, "분수|수|1/2"},
		{0x2160, "로마 숫자|수|I"}, {0xB2, "위첨자|수|2"}, {0xFF21, "전각|로마자|A"},
	}
	for _, c := range cases {
		row := deriveTagged(tbl, jamo, c.r, "Number Forms")
		assertEq(t, row.Form+"|"+row.Kind+"|"+row.Value, c.want)
	}
	tbl.Chars[0x2460] = charInfo{Decomp: "<odd> 0031"}
	assertEq(t, deriveTagged(tbl, jamo, 0x2460, "").Form, "<odd>")
}
