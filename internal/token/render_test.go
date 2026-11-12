//ff:func feature=token type=reader control=sequence
//ff:what TestRender — a single reading is one line; a contest a head line with what its first two branches are called and a line for each branch, the third and later too
package token

import "testing"

func TestRender(t *testing.T) {
	word := func(stem, marker, kind string) Eojeol {
		return Eojeol{Parts: []Part{{Text: stem, Stem: true}, {Text: marker, Names: []Name{{kind, marker}}}}}
	}
	one := render(Bundle{Text: "곰이", Branches: []Branch{{Words: []Eojeol{word("곰", "이", "주격")}}}}, 0)
	if len(one) != 1 || one[0] != "곰+이[주격:이]" {
		t.Fatal(one)
	}
	two := render(Bundle{Text: "작은", Branches: []Branch{{Words: []Eojeol{word("작", "은", "꾸밈")}, Verdict: 0.33, Pref: -5.29}, {Words: []Eojeol{word("작", "은", "보조사")}, Verdict: 0.33, Pref: -5.32}}}, 0)
	if len(two) != 3 || two[0] != "작은\t경합(새 낱말 + 꾸밈:은 / 새 낱말 + 보조사:은)" || two[1] != "  작+은[꾸밈:은]\t성립 0.33 선호 -5.29" || two[2] != "  작+은[보조사:은]\t성립 0.33 선호 -5.32" {
		t.Fatal(two)
	}
	three := render(Bundle{Text: "곰이", Branches: []Branch{
		{Words: []Eojeol{word("곰", "이", "주격")}, Verdict: 1, Pref: -2},
		{Words: []Eojeol{word("곰", "이", "보격")}, Verdict: 0.5, Pref: -3},
		{Words: []Eojeol{word("곰", "이", "서술격")}, Verdict: 0.1, Pref: -9.5},
	}}, 0)
	if len(three) != 4 || three[0] != "곰이\t경합(새 낱말 + 주격:이 / 새 낱말 + 보격:이)" || three[3] != "  곰+이[서술격:이]\t성립 0.10 선호 -9.50" {
		t.Fatal(three)
	}
}
