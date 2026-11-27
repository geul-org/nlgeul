//ff:func feature=parse type=reader control=sequence
//ff:what TestNumerals — every run of Hangul numerals that ends at a place, a Latin letter of the table that is a run by itself, the run of Han numerals and a Roman numeral of Latin capitals that is a run by itself, each named by the way it is written, and nothing for a longer run, a run of digits or another word
package parse

import "testing"

func TestNumerals(t *testing.T) {
	got := scanOf("삼십개").numerals(2)
	if len(got) != 2 || got[0].text != "십" || got[1].text != "삼십" || got[1].pref >= got[0].pref {
		t.Fatal(got)
	}
	if got = scanOf("N개").numerals(1); len(got) != 1 || got[0].text != "N" || got[0].pref >= 0 {
		t.Fatal(got)
	}
	if got = scanOf("二三일").numerals(2); len(got) != 1 || got[0].text != "二三" || got[0].name != numName {
		t.Fatal("a run of Han numerals", got)
	}
	if got = scanOf("XIV세").numerals(3); len(got) != 1 || got[0].text != "XIV" || got[0].name != romanNumName {
		t.Fatal("a Roman numeral of Latin capitals", got)
	}
	if got = scanOf("삼십개").numerals(2); got[0].name != numName || got[1].name != numName {
		t.Fatal(got)
	}
	if len(scanOf("Nx개").numerals(2)) != 0 || len(scanOf("3개").numerals(1)) != 0 || len(scanOf("곰개").numerals(1)) != 0 || len(scanOf("p개").numerals(1)) != 0 {
		t.Fatal("read where it should not be")
	}
}
