//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestRomanForms — a stem that is one Roman numeral letter of Unicode with nothing on it has a candidate named by the letter's shape with the verdict of a number and no kind (Ⅲ, ⅳ); a stem with a piece on it, letters in a row, a Latin capital, a full-width letter and a circled number have none
package parse

import "testing"

func TestRomanForms(t *testing.T) {
	ix := wordsScan("", layerCost).ix
	for _, stem := range []string{"Ⅲ", "ⅳ"} {
		got := ix.romanForms(stem, tail{})
		if len(got) != 1 || got[0].name != "로마 숫자:수" || got[0].kind != 0 || got[0].v != ix.number {
			t.Errorf("%q: %v", stem, got)
		}
	}
	if got := ix.romanForms("Ⅲ", tail{pieces: []piece{{}}}); got != nil {
		t.Fatal("a piece on the stem", got)
	}
	for _, stem := range []string{"ⅩⅢ", "I", "Ａ", "①", ""} {
		if got := ix.romanForms(stem, tail{}); got != nil {
			t.Errorf("%q: %v", stem, got)
		}
	}
}
