//ff:func feature=token type=helper control=sequence
//ff:what TestMidBranches — a word named 기호 is read by the candidates of its sign; two such words give every pairing; a candidate with none is given back as it is
package token

import "testing"

func TestMidBranches(t *testing.T) {
	sc := scanOf("가·나")
	sc.mid[1] = midSign{reads: []Part{{Text: "·", Name: Name{Kind: "가"}}, {Text: "·", Name: Name{Kind: "나"}}}}
	sc.mid[3] = midSign{reads: []Part{{Text: "ㆍ", Name: Name{Kind: "다"}}}}
	mid := func(end int) Eojeol {
		return Eojeol{Parts: []Part{{Text: "x", Name: Name{Kind: midName}}}, End: end}
	}
	br := Branch{Words: []Eojeol{mid(2), {Parts: []Part{{Text: "z"}}, End: 3}, mid(4)}, Verdict: 1}
	got := sc.midBranches(br)
	if len(got) != 2 || got[0].Words[0].Parts[0].Name.Kind != "가" || got[1].Words[0].Parts[0].Name.Kind != "나" || got[0].Words[2].Parts[0].Text != "ㆍ" || got[0].Words[1].Parts[0].Text != "z" {
		t.Fatal(got)
	}
	if br.Words[0].Parts[0].Text != "x" {
		t.Fatal("changed", br)
	}
	plain := Branch{Words: []Eojeol{{Parts: []Part{{Text: "z"}}, End: 3}}}
	if got = sc.midBranches(plain); len(got) != 1 {
		t.Fatal(got)
	}
}
