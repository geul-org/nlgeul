//ff:func feature=parse type=reader control=sequence
//ff:what TestCaseNumberBranch — a candidate's words are each sent up as caseNumberWord and plateWord give them (the word of a court case number as its pieces, the word of a registration number as its pieces, the others as they are), its verdict and preference are kept, and the candidate given is not changed
package parse

import "testing"

func TestCaseNumberBranch(t *testing.T) {
	in := Branch{Words: []Eojeol{
		{Parts: []Part{{Text: "대법원", Stem: true}}, End: 3},
		{Parts: []Part{{Text: "2004다25123", Name: Name{Kind: caseKind}, Stem: true}}, End: 13},
	}, Verdict: 0.5, Pref: -3}
	got := caseNumberBranch(in)
	if len(got.Words) != 2 || got.Words[0].String() != "대법원" || got.Words[1].String() != "2004[수:십진]+다[사건:다]+25123[수:십진]" {
		t.Fatal(got.Words)
	}
	if got.Verdict != 0.5 || got.Pref != -3 || got.Words[1].End != 13 {
		t.Fatal(got)
	}
	if len(in.Words[1].Parts) != 1 {
		t.Fatal("the candidate given is changed", in.Words[1].Parts)
	}
	plate := caseNumberBranch(Branch{Words: []Eojeol{{Parts: []Part{{Text: "306너8412", Name: Name{Kind: plateKind}, Stem: true}}, End: 8}}})
	if len(plate.Words) != 1 || plate.Words[0].String() != "306[수:십진]+너[차등록:너]+8412[수:십진]" || plate.Words[0].End != 8 {
		t.Fatal(plate.Words)
	}
	if got := caseNumberBranch(Branch{}); len(got.Words) != 0 {
		t.Fatal(got)
	}
}
