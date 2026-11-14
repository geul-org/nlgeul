//ff:func feature=combo type=helper control=sequence
//ff:what TestOrder — score's place is the sum of each candidate's preference plus Lambda times log10 of its verdict, its verdict the least and its preference the sum; none gives place 0, verdict 1, preference 0, and a candidate of verdict 1 and preference 0 changes nothing; a verdict of 0 or less gives minus infinity, not NaN; before puts the greater place first whatever the verdict and the preference, equal places neither way, and a weaker candidate with more preference comes first only when the preference makes up for Lambda times the gap in log10 verdict
package combo

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/token"
)

func TestOrder(t *testing.T) {
	score, before := order()
	if Lambda <= 0 {
		t.Fatalf("Lambda %v, want more than 0", float64(Lambda))
	}

	picked := []token.Branch{{Verdict: 0.5, Pref: -2}, {Verdict: 0.25, Pref: 1.5}, {Verdict: 0.75, Pref: -0.25}}
	place := -0.75 + Lambda*(math.Log10(0.5)+math.Log10(0.25)+math.Log10(0.75))
	got := score(picked)
	if math.Abs(got.Place-place) > 1e-9 || got.Verdict != 0.25 || got.Pref != -0.75 {
		t.Fatalf("score %+v, want place %v verdict 0.25 preference -0.75", got, place)
	}
	if got, want := score(nil), (Score{Place: 0, Verdict: 1, Pref: 0}); got != want {
		t.Fatalf("score of no candidate %+v, want %+v", got, want)
	}
	withSign := []token.Branch{{Verdict: 1, Pref: 0}, {Verdict: 0.5, Pref: -2}, {Verdict: 0.25, Pref: 1.5}, {Verdict: 0.75, Pref: -0.25}, {Verdict: 1, Pref: 0}}
	if with := score(withSign); math.Abs(with.Place-got.Place) > 1e-9 || with.Verdict != got.Verdict || with.Pref != got.Pref {
		t.Fatalf("score with signs %+v, want %+v", with, got)
	}

	zero := score([]token.Branch{{Verdict: 0.5, Pref: 3}, {Verdict: 0, Pref: 2}})
	if !math.IsInf(zero.Place, -1) || zero.Verdict != 0 || zero.Pref != 5 {
		t.Fatalf("score with a verdict of 0 %+v, want place -Inf verdict 0 preference 5", zero)
	}
	below := score([]token.Branch{{Verdict: -0.5, Pref: 1}})
	if !math.IsInf(below.Place, -1) || below.Verdict != -0.5 || below.Pref != 1 {
		t.Fatalf("score with a verdict below 0 %+v, want place -Inf verdict -0.5 preference 1", below)
	}

	first, second := Score{Place: -3, Verdict: 0.1, Pref: -50}, Score{Place: -4, Verdict: 0.9, Pref: 50}
	if !before(first, second) || before(second, first) {
		t.Fatalf("the greater place is not first: %+v %+v", first, second)
	}
	same := Score{Place: -1.25, Verdict: 0.5, Pref: 2}
	if before(same, same) || before(same, Score{Place: -1.25, Verdict: 0.9, Pref: 7}) {
		t.Fatalf("an equal place comes first: %+v", same)
	}
	if before(zero, below) || before(below, zero) {
		t.Fatalf("minus infinity comes before minus infinity: %+v %+v", zero, below)
	}
	if !before(got, zero) || before(zero, got) {
		t.Fatalf("a place that is a number is not before minus infinity: %+v %+v", got, zero)
	}

	// A weak candidate (verdict 0.33) against a strong one (verdict 0.6,
	// preference 0): the gap in log10 verdict times Lambda is what the
	// weak one's preference must make up.
	gap := Lambda * (math.Log10(0.6) - math.Log10(0.33))
	strong := score([]token.Branch{{Verdict: 0.6, Pref: 0}})
	short := score([]token.Branch{{Verdict: 0.33, Pref: gap / 2}})
	enough := score([]token.Branch{{Verdict: 0.33, Pref: gap * 2}})
	if !before(strong, short) || before(short, strong) {
		t.Fatalf("a weak candidate with too little preference comes first: %+v %+v", short, strong)
	}
	if !before(enough, strong) || before(strong, enough) {
		t.Fatalf("a weak candidate with preference enough does not come first: %+v %+v", enough, strong)
	}
}
