//ff:func feature=hypo type=helper control=sequence
//ff:what TestOrder — score is the least verdict and the summed preference of the readings picked, verdict 1 and preference 0 for none, and a reading of verdict 1 and preference 0 changes nothing; before puts the greater verdict first, then the greater preference, a lower verdict last whatever its preference, and equal scores neither way
package hypo

import (
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestOrder(t *testing.T) {
	score, before := order()

	picked := []token.Branch{{Verdict: 0.5, Pref: -2}, {Verdict: 0.25, Pref: 1.5}, {Verdict: 0.75, Pref: -0.25}}
	if got, want := score(picked), (Score{Verdict: 0.25, Pref: -0.75}); got != want {
		t.Fatalf("score %+v, want %+v", got, want)
	}
	if got, want := score(nil), (Score{Verdict: 1, Pref: 0}); got != want {
		t.Fatalf("score of no reading %+v, want %+v", got, want)
	}
	withSign := []token.Branch{{Verdict: 0.5, Pref: -2}, {Verdict: 1, Pref: 0}, {Verdict: 0.25, Pref: 1.5}, {Verdict: 0.75, Pref: -0.25}, {Verdict: 1, Pref: 0}}
	if got, want := score(withSign), (Score{Verdict: 0.25, Pref: -0.75}); got != want {
		t.Fatalf("score with signs %+v, want %+v", got, want)
	}
	if got, want := score([]token.Branch{{Verdict: 0.4, Pref: -3}}), (Score{Verdict: 0.4, Pref: -3}); got != want {
		t.Fatalf("score of one reading %+v, want %+v", got, want)
	}

	high, low := Score{Verdict: 0.5, Pref: -10}, Score{Verdict: 0.4, Pref: 5}
	if !before(high, low) || before(low, high) {
		t.Fatalf("the greater verdict is not first: %+v %+v", high, low)
	}
	more, less := Score{Verdict: 0.5, Pref: 2}, Score{Verdict: 0.5, Pref: 1}
	if !before(more, less) || before(less, more) {
		t.Fatalf("among equal verdicts the greater preference is not first: %+v %+v", more, less)
	}
	strong, rich := Score{Verdict: 0.4, Pref: -1000}, Score{Verdict: 0.3, Pref: 1000}
	if before(rich, strong) || !before(strong, rich) {
		t.Fatalf("a lower verdict comes first by its preference: %+v %+v", rich, strong)
	}
	same := Score{Verdict: 0.5, Pref: -1.25}
	if before(same, same) || before(same, Score{Verdict: 0.5, Pref: -1.25}) {
		t.Fatalf("an equal score comes first: %+v", same)
	}
}
