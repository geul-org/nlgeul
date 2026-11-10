//ff:func feature=token type=reader control=sequence
//ff:what TestStretch — the words of a reading within a stretch, with their least verdict and summed preference
package token

import "testing"

func TestStretch(t *testing.T) {
	word := func(text string, end int) Eojeol { return Eojeol{Parts: []Part{{Text: text, Stem: true}}, End: end} }
	p := grow(grow(grow(&path{verdict: 1}, step{end: 2, word: word("a", 2), verdict: 1, pref: -1}), step{end: 3, word: word("b", 3), verdict: 0.5, pref: -2}), step{end: 5, word: word("c", 5), verdict: 0.2, pref: -4})
	if got := stretch(p, 0, 3); got.String() != "a b" || len(got.Words) != 2 || got.Words[1].End != 3 || got.Verdict != 0.5 || got.Pref != -3 || got.Asleep {
		t.Fatal(got)
	}
}
