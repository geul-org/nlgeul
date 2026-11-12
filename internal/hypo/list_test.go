//ff:func feature=hypo type=helper control=iteration dimension=1
//ff:what TestList — as many hypotheses as Count, every number once; the greater verdict first, then the greater preference; equal scores in the order of their numbers; a bundle of one reading among them; no bundle gives none
package hypo

import (
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestList(t *testing.T) {
	// Numbers: the first bundle's reading *3 + the last bundle's reading
	// (the middle bundle has one reading, so its digit is always 0).
	//   0: a0 b c0 → verdict 0.5, preference 0+0+1 = 1
	//   1: a0 b c1 → verdict 1,   preference 0+0-1 = -1
	//   2: a0 b c2 → verdict 0.5, preference 0+0+1 = 1
	//   3: a1 b c0 → verdict 0.5, preference 2+0+1 = 3
	//   4: a1 b c1 → verdict 0.5, preference 2+0-1 = 1
	//   5: a1 b c2 → verdict 0.5, preference 2+0+1 = 3
	choices := []token.Bundle{
		{Text: "가", Branches: []token.Branch{{Verdict: 1, Pref: 0}, {Verdict: 0.5, Pref: 2}}},
		{Text: ".", Branches: []token.Branch{{Verdict: 1, Pref: 0}}},
		{Text: "나", Branches: []token.Branch{{Verdict: 0.5, Pref: 1}, {Verdict: 1, Pref: -1}, {Verdict: 0.5, Pref: 1}}},
	}
	want := []Hypothesis{
		{Number: 1, Score: Score{Verdict: 1, Pref: -1}},
		{Number: 3, Score: Score{Verdict: 0.5, Pref: 3}},
		{Number: 5, Score: Score{Verdict: 0.5, Pref: 3}},
		{Number: 0, Score: Score{Verdict: 0.5, Pref: 1}},
		{Number: 2, Score: Score{Verdict: 0.5, Pref: 1}},
		{Number: 4, Score: Score{Verdict: 0.5, Pref: 1}},
	}
	got := List(choices)
	if int64(len(got)) != Count(choices).Int64() || len(got) != len(want) {
		t.Fatalf("%d hypotheses, want %d (Count %s)", len(got), len(want), Count(choices))
	}
	seen := map[int]bool{}
	for i, h := range got {
		seen[h.Number] = true
		if h != want[i] {
			t.Errorf("rank %d: %+v, want %+v", i+1, h, want[i])
		}
	}
	if len(seen) != 6 {
		t.Fatalf("numbers %v, want 0 to 5 once each", seen)
	}

	if got := List(nil); len(got) != 0 {
		t.Fatalf("no bundle gives %+v", got)
	}
}
