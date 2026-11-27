//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLoneForms — a stem that is one letter of the tables of special letters read as a letter of a run (Ⅲ, ⅳ, Ａ, １) has a candidate named by its shape and what it writes, with the low verdict of a letter read alone by its shape and the kind that stands behind the other readings; a stem with a piece on it, a stem of two such letters, a letter of the tables that is no letter of a run (①, ⓐ) and a letter the tables lack have none
package parse

import "testing"

func TestLoneForms(t *testing.T) {
	ix := wordsScan("", layerCost).ix
	for stem, name := range map[string]string{"Ⅲ": "로마 숫자:수", "ⅳ": "로마 숫자:수", "Ａ": "전각:로마자", "１": "전각:수"} {
		got := ix.loneForms(stem, tail{})
		if len(got) != 1 || got[0].name != name || got[0].kind != kindBehind || got[0].v != ix.loneForm {
			t.Errorf("%q: %v", stem, got)
		}
	}
	if ix.loneForm <= 0 || ix.loneForm >= 0.5 || ix.loneForm >= ix.number {
		t.Fatal(ix.loneForm)
	}
	if got := ix.loneForms("Ⅲ", tail{pieces: []piece{{}}}); got != nil {
		t.Fatal("a piece on the stem", got)
	}
	for _, stem := range []string{"ⅩⅢ", "①", "ⓐ", "A", "가", ""} {
		if got := ix.loneForms(stem, tail{}); got != nil {
			t.Errorf("%q: %v", stem, got)
		}
	}
}
