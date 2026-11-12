//ff:func feature=hypo type=helper control=iteration dimension=1
//ff:what List — every combination of candidates of the bundles with its score, sorted by the formula of order; equal scores stay in the order of their numbers
package hypo

import (
	"sort"

	"github.com/geul-org/nlgeul/internal/token"
)

// All of them are held at once: the caller looks at Count first and
// does not ask for more than Limit.
func List(choices []token.Bundle) []Hypothesis {
	score, before := order()
	out := make([]Hypothesis, Count(choices).Int64())
	for n := range out {
		out[n] = Hypothesis{Number: n, Score: score(Picked(choices, n))}
	}
	sort.SliceStable(out, func(i, j int) bool { return before(out[i].Score, out[j].Score) })
	return out
}
