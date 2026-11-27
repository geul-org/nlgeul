//ff:func feature=data type=helper control=iteration dimension=1
//ff:what TestGroupRowOf — a symbol gives a five-column line with its own value, and an opening bracket, a letter, an unassigned code point and a character the form table holds give none
package main

import "testing"

func TestGroupRowOf(t *testing.T) {
	tbl := loadSample(t)
	jamo := jamoMap(tbl)
	line, ok := groupRowOf(tbl, jamo, 0x2044, "General Punctuation")
	if !ok {
		t.Fatal("Sm")
	}
	assertEq(t, line, "⁄\t기호\t수학\t⁄\t-")
	line, _ = groupRowOf(tbl, jamo, 0x2E, "Basic Latin")
	assertEq(t, line, ".\t기호\t그 밖\t.\t-")
	if _, ok := groupRowOf(tbl, jamo, 0x24B6, "Enclosed Alphanumerics"); ok {
		t.Fatal("a symbol the form table holds")
	}
	for _, r := range []rune{0x28, 0x41, 0x0378, 0x2460, 0xBD} {
		if _, ok := groupRowOf(tbl, jamo, r, "Basic Latin"); ok {
			t.Fatalf("%U", r)
		}
	}
}
