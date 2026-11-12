//ff:func feature=token type=reader control=sequence
//ff:what TestMovedOne — the words of a reading are carried over to the letters after the sign: the first without the sign, the others with their ends moved back; a reading of the sign alone is nil
package token

import "testing"

func TestMovedOne(t *testing.T) {
	sc := scanOf(")를 가")
	sc.atoms = map[int]atom{0: {1, pairName, false}}
	words := []step{
		{end: 2, word: Eojeol{Parts: []Part{{Text: ")", Stem: true, Names: []Name{{Kind: pairName}}}, {Text: "를"}}, End: 2}, verdict: 1, pref: -6},
		{end: 4, word: Eojeol{Parts: []Part{{Text: "가"}}, End: 4}, verdict: 0.5, pref: -2, sig: "가"},
	}
	got := sc.movedOne(words, 1)
	w := wordsOf(got)
	if len(w) != 2 || w[0].end != 1 || !w[0].word.Attach || w[1].end != 3 || w[1].word.End != 3 || got.verdict != 0.5 {
		t.Fatal(w)
	}
	if words[1].end != 4 || words[1].word.End != 4 {
		t.Fatal("changed", words)
	}
	alone := []step{{end: 1, word: Eojeol{Parts: []Part{{Text: ")"}}, End: 1}, verdict: 1}}
	if sc.movedOne(alone, 1) != nil {
		t.Fatal("nothing after the sign")
	}
}
