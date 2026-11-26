//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCaseNumberWord — a word whose stem is a court case number is sent up with the stem as three parts (the year, a number and the stem; the mark, named 사건 with the mark as its form; the serial number, a number) and the parts after the stem as they are; a word with another stem, a part named 사건 with a form, a part that is no stem and a word with no part are given back as they are, and the word given is not changed
package parse

import "testing"

func TestCaseNumberWord(t *testing.T) {
	in := Eojeol{Parts: []Part{{Text: "2008회합9", Name: Name{Kind: caseKind}, Stem: true}, {Text: "호", Name: Name{"단위", "호"}}}, End: 9}
	got := caseNumberWord(in)
	if got.String() != "2008[수]+회합[사건:회합]+9[수]+호[단위:호]" || got.End != 9 {
		t.Fatal(got.String(), got.End)
	}
	if !got.Parts[0].Stem || got.Parts[1].Stem || got.Parts[2].Stem || got.Parts[3].Stem {
		t.Fatal("the year is the stem", got.Parts)
	}
	if len(in.Parts) != 2 || in.Parts[0].Text != "2008회합9" {
		t.Fatal("the word given is changed", in.Parts)
	}
	if got := caseNumberWord(Eojeol{Parts: []Part{{Text: "97도2836", Name: Name{Kind: caseKind}, Stem: true}}}); got.String() != "97[수]+도[사건:도]+2836[수]" {
		t.Fatal(got.String())
	}
	for _, w := range []Eojeol{
		{},
		{Parts: []Part{{Text: "2004다25123", Stem: true}}},
		{Parts: []Part{{Text: "사건", Name: Name{"열쇠", "체언"}, Stem: true}}},
		{Parts: []Part{{Text: "다", Name: Name{caseKind, "다"}, Stem: true}}},
		{Parts: []Part{{Text: "2004다25123", Name: Name{Kind: caseKind}}}},
	} {
		if got := caseNumberWord(w); got.String() != w.String() || len(got.Parts) != len(w.Parts) {
			t.Fatal(w.String(), got.String())
		}
	}
}
