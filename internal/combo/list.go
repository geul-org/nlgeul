//ff:func feature=combo type=helper control=iteration dimension=1
//ff:what List — every combination of candidates of the contests with its score, sorted by the formula of order; equal scores stay in the order of their numbers
package combo

import (
	"sort"

	"github.com/geul-org/nlgeul/internal/token"
)

// All of them are held at once: the caller looks at Count first and
// does not ask for more than Limit.
func List(choices []token.Contest) []Combination {
	score, before := order()
	out := make([]Combination, Count(choices).Int64())
	for n := range out {
		out[n] = Combination{Number: n, Score: score(Picked(choices, n))}
	}
	sort.SliceStable(out, func(i, j int) bool { return before(out[i].Score, out[j].Score) })
	return out
}
