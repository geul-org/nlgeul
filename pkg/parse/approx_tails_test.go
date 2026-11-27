//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestApproxTails — a unit after an approximate number written in Hangul is read with the first number as the stem and the second as a piece before the unit, both named 수:한글십진 (칠팔년: 칠, 팔, 년); with a word table the two numbers and a unit that are a word together stand with a low verdict; letters running on after the unit leave no such tail, and a greater number takes none
package parse

import "testing"

func TestApproxTails(t *testing.T) {
	var got []tail
	for _, x := range scanOf("칠팔년").units(3, 0, true) {
		if len(x.pieces) == 2 {
			got = append(got, x)
		}
	}
	if len(got) != 1 || got[0].start != 1 || got[0].only != "칠" || got[0].as != numName || got[0].pieces[0].text != "팔" || got[0].pieces[0].names != numName || got[0].pieces[1].names != "단위:년" || got[0].pieces[1].verdict != 1 || got[0].weight >= 0 {
		t.Fatal(got)
	}
	sc := wordsScan("칠팔년", layerCost)
	sc.ix.lex["칠팔년"] = word{kinds: kindNominal, count: 1}
	for _, x := range sc.units(3, 0, true) {
		if len(x.pieces) == 2 && x.pieces[1].verdict >= 0.1 {
			t.Fatal("two numbers and a unit that are a word together", x)
		}
	}
	for _, x := range wordsScan("칠팔년곰", layerCost).units(3, 0, true) {
		if len(x.pieces) == 2 {
			t.Fatal("letters running on after the unit", x)
		}
	}
	for _, x := range scanOf("이삼만").units(3, 0, true) {
		if len(x.pieces) == 2 {
			t.Fatal("a greater number after two numbers", x)
		}
	}
}
