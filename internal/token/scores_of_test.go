//ff:func feature=token type=helper control=sequence
//ff:what TestScoresOf — the scores kept with a branch are given as they are; a branch made without them gets its own verdict and preference for each of its words
package token

import "testing"

func TestScoresOf(t *testing.T) {
	kept := Branch{Words: make([]Eojeol, 2), Verdict: 0.3, Pref: -3, scores: []wordScore{{1, -1, 0}, {0.3, -2, 1}}}
	if got := scoresOf(kept); len(got) != 2 || got[0] != (wordScore{1, -1, 0}) || got[1] != (wordScore{0.3, -2, 1}) {
		t.Fatal(got)
	}
	bare := Branch{Words: make([]Eojeol, 1), Verdict: 0.4, Pref: -5}
	if got := scoresOf(bare); len(got) != 1 || got[0] != (wordScore{0.4, -5, noBorn}) {
		t.Fatal(got)
	}
	if got := scoresOf(Branch{Words: make([]Eojeol, 2), Verdict: 0.4, Pref: -5}); len(got) != 2 || got[1] != (wordScore{0.4, -5, noBorn}) {
		t.Fatal(got)
	}
}
