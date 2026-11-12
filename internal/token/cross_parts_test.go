//ff:func feature=token type=helper control=sequence
//ff:what TestCrossParts — each reading given is read by each part: the first reading's parts come first
package token

import "testing"

func TestCrossParts(t *testing.T) {
	one := Branch{Words: []Eojeol{{Parts: []Part{{Text: "x"}}}}, Pref: 1}
	two := Branch{Words: []Eojeol{{Parts: []Part{{Text: "y"}}}}, Pref: 2}
	got := crossParts([]Branch{one, two}, 0, []Part{{Text: "a"}, {Text: "b"}})
	if len(got) != 4 || got[0].Pref != 1 || got[1].Pref != 1 || got[2].Pref != 2 || got[1].Words[0].Parts[0].Text != "b" || got[2].Words[0].Parts[0].Text != "a" {
		t.Fatal(got)
	}
}
