//ff:func feature=hypo type=helper control=sequence
//ff:what order — the sorting formula of nlgeul hypo, all of it: score makes a hypothesis's score from the readings picked for it, and before says which of two scores comes first
package hypo

import (
	"math"

	"github.com/geul-org/nlgeul/internal/token"
)

// To sort the hypotheses another way, change this function and nothing
// else: List sorts with what it gives, and the lines print the Score.
//
// score: the verdict of a hypothesis is the least verdict of its
// readings, and its preference the sum of their preferences. A bundle
// with one reading counts like any other, with the verdict and the
// preference the token layer gave that reading (signs and closing dots
// come with verdict 1 and preference 0, so they change nothing).
//
// before: the greater verdict first; among equal verdicts the greater
// preference first. Two scores that are equal are left in the order of
// their numbers (List sorts stably), so the output is always the same.
func order() (score func(picked []token.Branch) Score, before func(a, b Score) bool) {
	score = func(picked []token.Branch) Score {
		s := Score{Verdict: 1}
		for _, br := range picked {
			s.Verdict, s.Pref = math.Min(s.Verdict, br.Verdict), s.Pref+br.Pref
		}
		return s
	}
	before = func(a, b Score) bool {
		if a.Verdict != b.Verdict {
			return a.Verdict > b.Verdict
		}
		return a.Pref > b.Pref
	}
	return score, before
}
