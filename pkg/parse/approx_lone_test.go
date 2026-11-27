//ff:func feature=parse type=reader control=sequence
//ff:what TestApproxLone — an approximate number written in Hangul with no unit written onto it is a tail of one piece, the second number, on the first number as the stem, both named 수:한글십진, with the verdict of a number read with nothing that tells it and each letter valued as a new word's; a place where no such number ends has none, and with no word table none is read
package parse

import (
	"math"
	"testing"
)

func TestApproxLone(t *testing.T) {
	if got := scanOf("육칠").approxLone(2); got != nil {
		t.Fatal("no word table", got)
	}
	sc := wordsScan("육칠", layerCost)
	got := sc.approxLone(2)
	cost := math.Log10(1.0 / syllables)
	if len(got) != 1 || got[0].start != 1 || got[0].only != "육" || got[0].as != numName || len(got[0].pieces) != 1 || got[0].pieces[0].text != "칠" || got[0].pieces[0].names != numName {
		t.Fatal(got)
	}
	if got[0].pieces[0].verdict != sc.ix.loneNumber || got[0].pieces[0].verdict >= 0.5 || got[0].weight != cost || got[0].pieces[0].pref != cost {
		t.Fatal(got[0].pieces[0], got[0].weight)
	}
	if got = wordsScan("이삼십", layerCost).approxLone(3); len(got) != 1 || got[0].only != "이" || got[0].pieces[0].text != "삼십" || got[0].pieces[0].pref != 2*cost {
		t.Fatal(got)
	}
	if got = wordsScan("삼십", layerCost).approxLone(2); got != nil {
		t.Fatal(got)
	}
}
