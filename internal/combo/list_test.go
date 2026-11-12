//ff:func feature=combo type=helper control=iteration dimension=1
//ff:what TestList — as many combinations as Count, every number once; the greater place first; equal places in the order of their numbers; the combinations that pick a candidate of verdict 0 last, in the order of their numbers; a contest of one candidate among them; no contest gives none
package combo

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestList(t *testing.T) {
	// Verdicts are 1 (log10 0) or 0 (minus infinity), so the order is
	// the same whatever Lambda is. Numbers: the first contest's candidate
	// *3 + the last contest's candidate (the middle contest has one
	// candidate, so its digit is always 0).
	//   0: a0 b c0 → place 0+0+1 = 1
	//   1: a0 b c1 → place minus infinity (c1 has verdict 0)
	//   2: a0 b c2 → place 0+0+1 = 1
	//   3: a1 b c0 → place 2+0+1 = 3
	//   4: a1 b c1 → place minus infinity
	//   5: a1 b c2 → place 2+0+1 = 3
	choices := []token.Contest{
		{Text: "가", Branches: []token.Branch{{Verdict: 1, Pref: 0}, {Verdict: 1, Pref: 2}}},
		{Text: ".", Branches: []token.Branch{{Verdict: 1, Pref: 0}}},
		{Text: "나", Branches: []token.Branch{{Verdict: 1, Pref: 1}, {Verdict: 0, Pref: 9}, {Verdict: 1, Pref: 1}}},
	}
	inf := math.Inf(-1)
	want := []Combination{
		{Number: 3, Score: Score{Place: 3, Verdict: 1, Pref: 3}},
		{Number: 5, Score: Score{Place: 3, Verdict: 1, Pref: 3}},
		{Number: 0, Score: Score{Place: 1, Verdict: 1, Pref: 1}},
		{Number: 2, Score: Score{Place: 1, Verdict: 1, Pref: 1}},
		{Number: 1, Score: Score{Place: inf, Verdict: 0, Pref: 9}},
		{Number: 4, Score: Score{Place: inf, Verdict: 0, Pref: 11}},
	}
	got := List(choices)
	if int64(len(got)) != Count(choices).Int64() || len(got) != len(want) {
		t.Fatalf("%d combinations, want %d (Count %s)", len(got), len(want), Count(choices))
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
		t.Fatalf("no contest gives %+v", got)
	}
}
