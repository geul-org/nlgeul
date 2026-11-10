//ff:func feature=token type=reader control=sequence
//ff:what TestRender — a single reading is one line; a contest a head line with what its first two branches are called and a line for each branch; readings asleep a head line with how many there are and a line for each
package token

import "testing"

func TestRender(t *testing.T) {
	word := func(stem, marker, kind string) Eojeol {
		return Eojeol{Parts: []Part{{Text: stem, Stem: true}, {Text: marker, Names: []Name{{kind, marker}}}}}
	}
	one := render(Bundle{Text: "곰이", Branches: []Branch{{Words: []Eojeol{word("곰", "이", "주격")}}}})
	if len(one) != 1 || one[0] != "곰+이[주격:이]" {
		t.Fatal(one)
	}
	two := render(Bundle{Text: "작은", Branches: []Branch{{Words: []Eojeol{word("작", "은", "꾸밈")}, Verdict: 0.33, Pref: -5.29}, {Words: []Eojeol{word("작", "은", "보조사")}, Verdict: 0.33, Pref: -5.32}}})
	if len(two) != 3 || two[0] != "작은\t경합(새 낱말 + 꾸밈:은 / 새 낱말 + 보조사:은)" || two[1] != "  작+은[꾸밈:은]\t성립 0.33 선호 -5.29" {
		t.Fatal(two)
	}
	asleep := render(Bundle{Text: "곰이", Branches: []Branch{{Words: []Eojeol{word("곰", "이", "주격")}, Verdict: 0.5, Pref: -3, Asleep: true}}})
	if len(asleep) != 2 || asleep[0] != "곰이\t잠듦(1)" || asleep[1] != "  곰+이[주격:이]\t성립 0.50 선호 -3.00" {
		t.Fatal(asleep)
	}
}
