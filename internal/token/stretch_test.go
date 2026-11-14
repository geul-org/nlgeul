//ff:func feature=token type=reader control=sequence
//ff:what TestStretch — the words of a candidate within a stretch, with their least verdict and summed preference, and the verdict, preference and number of birth of each word
package token

import "testing"

func TestStretch(t *testing.T) {
	word := func(text string, end int) Eojeol { return Eojeol{Parts: []Part{{Text: text, Stem: true}}, End: end} }
	p := grow(grow(grow(&path{verdict: 1}, step{end: 2, word: word("a", 2), verdict: 1, pref: -1, born: 4}), step{end: 3, word: word("b", 3), verdict: 0.5, pref: -2, born: 9}), step{end: 5, word: word("c", 5), verdict: 0.2, pref: -4})
	if got := stretch(p, 0, 3); got.String() != "a b" || len(got.Words) != 2 || got.Words[1].End != 3 || got.Verdict != 0.5 || got.Pref != -3 {
		t.Fatal(got)
	}
	if got := stretch(p, 0, 3); len(got.scores) != 2 || got.scores[0] != (wordScore{1, -1, 4}) || got.scores[1] != (wordScore{0.5, -2, 9}) {
		t.Fatal(got.scores)
	}
}
