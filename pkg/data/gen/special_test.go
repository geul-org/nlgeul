//ff:func feature=data type=helper control=iteration dimension=1
//ff:what TestSpecial — punctuation, symbols and number letters are special, plain letters and digits are not, a letter with a compatibility decomposition is, an Arabic presentation form is not
package main

import "testing"

func TestSpecial(t *testing.T) {
	tbl := loadSample(t)
	for _, r := range []rune{0x28, 0x2044, 0x2160, 0xBD, 0xFF21, 0x3131} {
		if !special(tbl, r) {
			t.Fatalf("%U should be special", r)
		}
	}
	for _, r := range []rune{0x41, 0x31, 0x20, 0x3041, 0xFB50, 0x3B1} {
		if special(tbl, r) {
			t.Fatalf("%U should not be special", r)
		}
	}
}
