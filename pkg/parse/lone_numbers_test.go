//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLoneNumbers — a stem whose letters are a number that is no run of digits has one candidate more, named by the way the number is written, with the low verdict of a number nothing in its word tells and the kind that stands behind the words of the word table; a stem with a piece on it and a stem that is no such number have none, and numerals written in Hangul have none with no word table
package parse

import "testing"

func TestLoneNumbers(t *testing.T) {
	ix, bare := wordsScan("", layerCost).ix, scanOf("").ix
	for stem, name := range map[string]string{"사십": numName, "十": numName, "II": romanNumName} {
		got := ix.loneNumbers(stem, tail{})
		if len(got) != 1 || got[0].name != name || got[0].kind != kindBehind || got[0].v != ix.loneNumber {
			t.Errorf("%q: %v", stem, got)
		}
	}
	if ix.loneNumber <= 0 || ix.loneNumber >= 0.5 || ix.loneNumber >= ix.number {
		t.Fatal(ix.loneNumber)
	}
	if got := ix.loneNumbers("사십", tail{pieces: []piece{{}}}); got != nil {
		t.Fatal("a piece on the stem", got)
	}
	if got := bare.loneNumbers("사십", tail{}); got != nil {
		t.Fatal("numerals written in Hangul with no word table", got)
	}
	if got := bare.loneNumbers("十", tail{}); len(got) != 1 {
		t.Fatal("Han numerals with no word table", got)
	}
	if got := ix.loneNumbers("곰", tail{}); got != nil {
		t.Fatal(got)
	}
}
