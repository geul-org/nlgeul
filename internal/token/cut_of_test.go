//ff:func feature=token type=helper control=sequence
//ff:what TestCutOf — a branch of two words becomes spans from where the stretch begins in the sentence (the ends are token offsets of the unit: the last one stands for the stretch's end), each word a branch of its own with its own values; a branch of one word is the whole stretch
package token

import "testing"

func TestCutOf(t *testing.T) {
	a := Eojeol{Parts: []Part{{Text: "ab", Stem: true}}, End: 8}
	b := Eojeol{Parts: []Part{{Text: "cde", Stem: true}}, End: 11}
	br := Branch{Words: []Eojeol{a, b}, Verdict: 0.5, Pref: -3, scores: []wordScore{{1, -1, 4}, {0.5, -2, 5}}}
	c := cutOf(br, 6, 5)
	if len(c.spans) != 2 || c.spans[0] != (Span{6, 8}) || c.spans[1] != (Span{8, 11}) {
		t.Fatal(c.spans)
	}
	if len(c.words) != 2 || c.words[0].String() != "ab" || c.words[0].Pref != -1 || c.words[1].String() != "cde" || c.words[1].Verdict != 0.5 || c.words[1].Pref != -2 || len(c.words[1].Words) != 1 {
		t.Fatal(c.words)
	}
	if c.words[0].scores[0].born != 4 || c.words[1].scores[0].born != 5 {
		t.Fatal("the number of birth goes with the word", c.words[0].scores, c.words[1].scores)
	}
	one := cutOf(Branch{Words: []Eojeol{b}, Verdict: 0.7, Pref: -4}, 6, 5)
	if len(one.spans) != 1 || one.spans[0] != (Span{6, 11}) || one.words[0].Verdict != 0.7 || one.words[0].Pref != -4 || one.words[0].scores[0].born != noBorn {
		t.Fatal(one)
	}
}
