//ff:func feature=token type=helper control=sequence
//ff:what TestWithPart — a copy of a candidate has the one part given as its word k, and the candidate itself is not changed
package token

import "testing"

func TestWithPart(t *testing.T) {
	br := Branch{Words: []Eojeol{{Parts: []Part{{Text: "가"}}}, {Parts: []Part{{Text: "x"}, {Text: "y"}}}}, Pref: -2}
	got := withPart(br, 1, Part{Text: "·"})
	if len(got.Words[1].Parts) != 1 || got.Words[1].Parts[0].Text != "·" || got.Words[0].Parts[0].Text != "가" || got.Pref != -2 {
		t.Fatal(got)
	}
	if len(br.Words[1].Parts) != 2 || br.Words[1].Parts[0].Text != "x" {
		t.Fatal("changed", br)
	}
}
