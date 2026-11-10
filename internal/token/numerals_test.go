//ff:func feature=token type=reader control=sequence
//ff:what TestNumerals — every run of Hangul numerals that ends at a place, a Latin letter of the table that is a run by itself, and nothing for a longer run, a run of digits or another word
package token

import "testing"

func TestNumerals(t *testing.T) {
	got := scanOf("삼십개").numerals(2)
	if len(got) != 2 || got[0].text != "십" || got[1].text != "삼십" || got[1].pref >= got[0].pref {
		t.Fatal(got)
	}
	if got = scanOf("N개").numerals(1); len(got) != 1 || got[0].text != "N" || got[0].pref >= 0 {
		t.Fatal(got)
	}
	if len(scanOf("Nx개").numerals(2)) != 0 || len(scanOf("3개").numerals(1)) != 0 || len(scanOf("곰개").numerals(1)) != 0 || len(scanOf("p개").numerals(1)) != 0 {
		t.Fatal("read where it should not be")
	}
}
