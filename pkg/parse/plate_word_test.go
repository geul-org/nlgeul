//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPlateWord — a word whose stem is a registration number is sent up with the stem as its pieces (the kind, a number and the stem; the letter, named 차등록 with the letter as its form; the serial number, a number) and the parts after the stem as they are; with the name of a region before the kind the region is a nominal and the stem; a word with another stem, a part named 차등록 with a form, a part that is no stem and a word with no part are given back as they are, and the word given is not changed
package parse

import "testing"

func TestPlateWord(t *testing.T) {
	in := Eojeol{Parts: []Part{{Text: "74모2238", Name: Name{Kind: plateKind}, Stem: true}, {Text: "호", Name: Name{"단위", "호"}}}, End: 8}
	got := plateWord(in)
	if got.String() != "74[수]+모[차등록:모]+2238[수]+호[단위:호]" || got.End != 8 {
		t.Fatal(got.String(), got.End)
	}
	if !got.Parts[0].Stem || got.Parts[1].Stem || got.Parts[2].Stem || got.Parts[3].Stem {
		t.Fatal("the kind is the stem", got.Parts)
	}
	if len(in.Parts) != 2 || in.Parts[0].Text != "74모2238" {
		t.Fatal("the word given is changed", in.Parts)
	}
	got = plateWord(Eojeol{Parts: []Part{{Text: "서울34가5678", Name: Name{Kind: plateKind}, Stem: true}}})
	if got.String() != "서울[열쇠:체언]+34[수]+가[차등록:가]+5678[수]" {
		t.Fatal(got.String())
	}
	if !got.Parts[0].Stem || got.Parts[1].Stem || got.Parts[2].Stem || got.Parts[3].Stem {
		t.Fatal("the region is the stem", got.Parts)
	}
	for _, w := range []Eojeol{
		{},
		{Parts: []Part{{Text: "306너8412", Stem: true}}},
		{Parts: []Part{{Text: "차량", Name: Name{"열쇠", "체언"}, Stem: true}}},
		{Parts: []Part{{Text: "너", Name: Name{plateKind, "너"}, Stem: true}}},
		{Parts: []Part{{Text: "306너8412", Name: Name{Kind: plateKind}}}},
	} {
		if got := plateWord(w); got.String() != w.String() || len(got.Parts) != len(w.Parts) {
			t.Fatal(w.String(), got.String())
		}
	}
}
