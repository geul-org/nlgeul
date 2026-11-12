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
// score: the place of a hypothesis is the sum, over its readings, of
// the reading's preference plus Lambda times log10 of its verdict. Its
// verdict is the least verdict of its readings and its preference the
// sum of their preferences: they are printed, not sorted by. A bundle
// with one reading counts like any other, with the verdict and the
// preference the token layer gave that reading (signs and closing dots
// come with verdict 1 and preference 0, so they change nothing). A
// reading whose verdict is 0 or less has no log10: it counts as minus
// infinity, and so does the place of every hypothesis that picks it.
//
// before: the greater place first. Two places that are equal are left
// in the order of their numbers (List sorts stably), so the output is
// always the same; the hypotheses whose place is minus infinity come
// last, in the order of their numbers.
func order() (score func(picked []token.Branch) Score, before func(a, b Score) bool) {
	score = func(picked []token.Branch) Score {
		s := Score{Verdict: 1}
		for _, br := range picked {
			s.Place += br.Pref + float64(Lambda*math.Log10(math.Max(br.Verdict, 0)))
			s.Verdict, s.Pref = math.Min(s.Verdict, br.Verdict), s.Pref+br.Pref
		}
		return s
	}
	before = func(a, b Score) bool {
		return a.Place > b.Place
	}
	return score, before
}
