//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestRomanBracketed — a Roman numeral of Latin capitals that is the whole of a unit standing alone in a round bracket written onto a word is read as a number named 수:로마십진 with a number's verdict; not in a unit that does not stand so, as a part of the unit, with a piece on it, or when its letters leave it open (one letter, XX)
package parse

import "testing"

func TestRomanBracketed(t *testing.T) {
	sc := scanOf("III")
	if got := sc.romanBracketed(0, "III", tail{start: 3}); got != nil {
		t.Fatal("a unit that stands in no bracket", got)
	}
	sc.bracketed = true
	got := sc.romanBracketed(0, "III", tail{start: 3})
	if len(got) != 1 || got[0].name != romanNumName || got[0].v != sc.ix.number || got[0].kind != 0 {
		t.Fatal(got)
	}
	if got = sc.romanBracketed(1, "II", tail{start: 3}); got != nil {
		t.Fatal("a part of the unit", got)
	}
	if got = sc.romanBracketed(0, "II", tail{start: 2}); got != nil {
		t.Fatal("a part of the unit", got)
	}
	if got = sc.romanBracketed(0, "III", tail{start: 3, pieces: []piece{{}}}); got != nil {
		t.Fatal("a piece on the numeral", got)
	}
	for _, text := range []string{"I", "XX", "CD", "곰"} {
		sc = scanOf(text)
		sc.bracketed = true
		if got = sc.romanBracketed(0, text, tail{start: len([]rune(text))}); got != nil {
			t.Errorf("%q: %v", text, got)
		}
	}
}
