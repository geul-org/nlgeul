//ff:func feature=token type=reader control=sequence
//ff:what TestNumeralTails — a unit after a native numeral stands as it is; after a lone Sino-Korean digit with no word table it is not read; with a word table a numeral and a unit that are a word together stand with a low verdict, and others with the unit's own
package token

import "testing"

func TestNumeralTails(t *testing.T) {
	got := scanOf("세명").units(2, 0, true)
	if len(got) != 1 || got[0].pieces[0].verdict != 1 {
		t.Fatal(got)
	}
	if got = scanOf("이건").units(2, 0, true); len(got) != 0 {
		t.Fatal("a lone Sino-Korean digit with no word table", got)
	}
	if got = scanOf("이십건").units(3, 0, true); len(got) != 2 || got[1].only != "이십" {
		t.Fatal(got)
	}
	got = wordsScan("세대", layerCost).units(2, 0, true)
	if len(got) != 1 || got[0].pieces[0].verdict <= 0 || got[0].pieces[0].verdict >= 0.1 {
		t.Fatal(got)
	}
	got = wordsScan("삼대", layerCost).units(2, 0, true)
	if len(got) != 1 || got[0].pieces[0].verdict != 1 {
		t.Fatal(got)
	}
}
