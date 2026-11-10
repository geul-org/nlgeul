//ff:func feature=token type=reader control=sequence
//ff:what TestRender — one reading on one line; a contest with a head line naming its kind and a line for each branch
package token

import "testing"

func TestRender(t *testing.T) {
	one := render(bundle{"곰이", []branch{{marked: "곰+이"}}})
	if len(one) != 1 || one[0] != "곰+이" {
		t.Fatal(one)
	}
	two := render(bundle{"작은", []branch{{"작+은[꾸밈:은]", "새 낱말 + 꾸밈:은", 0.33, -5.29}, {"작+은[보조사:은]", "새 낱말 + 보조사:은", 0.33, -5.32}}})
	if len(two) != 3 || two[0] != "작은\t경합(새 낱말 + 꾸밈:은 / 새 낱말 + 보조사:은)" || two[1] != "  작+은[꾸밈:은]\t성립 0.33 선호 -5.29" {
		t.Fatal(two)
	}
}
