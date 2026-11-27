//ff:func feature=parse type=reader control=sequence
//ff:what TestCaseNumberParts — every candidate of every contest has its word of a court case number sent up as the number's pieces; the contests' letters, their order and the order of their candidates are kept, the contests given are not changed, and none gives none
package parse

import "testing"

func TestCaseNumberParts(t *testing.T) {
	number := Eojeol{Parts: []Part{{Text: "97도2836", Name: Name{Kind: caseKind}, Stem: true}}, End: 8}
	plain := Eojeol{Parts: []Part{{Text: "판결", Stem: true}}, End: 2}
	in := []Contest{
		{Unit: 1, Text: "97도2836", Branches: []Branch{{Words: []Eojeol{number}, Verdict: 1}, {Words: []Eojeol{plain}, Verdict: 0.4}}},
		{Unit: 2, Glued: true, Text: "판결", Branches: []Branch{{Words: []Eojeol{plain}}}},
	}
	got := caseNumberParts(in)
	if len(got) != 2 || got[0].Unit != 1 || got[0].Text != "97도2836" || !got[1].Glued || got[1].Text != "판결" {
		t.Fatal(got)
	}
	if len(got[0].Branches) != 2 || got[0].Branches[0].Words[0].String() != "97[수:십진]+도[사건:도]+2836[수:십진]" || got[0].Branches[0].Verdict != 1 {
		t.Fatal(got[0].Branches)
	}
	if got[0].Branches[1].Words[0].String() != "판결" || got[1].Branches[0].Words[0].String() != "판결" {
		t.Fatal(got)
	}
	if len(in[0].Branches[0].Words[0].Parts) != 1 {
		t.Fatal("the contests given are changed", in[0].Branches[0].Words[0].Parts)
	}
	if got := caseNumberParts(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
