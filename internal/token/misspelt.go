//ff:func feature=token type=rule control=iteration dimension=1
//ff:what misspelt — the right form of a misspelt stem as a known word with what follows it: its name and its verdict, the best among the kinds both the right form and the misspelt form have (0 when no kind takes that tail); each kind's graph is told that the stem is a misspelt form
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// It is known's second layer for a misspelt stem: in the first layer
// every kind counts as standing. The right form's own count is the
// stem's (the count of a misspelt line is 0, nothing was counted).
func (ix *index) misspelt(right string, kinds, tail int, sole bool) (string, float64) {
	w := ix.lex[right]
	var names []string
	best := 0.0
	for i, k := range kindOrder {
		if w.kinds&kinds&k == 0 {
			continue
		}
		v := 1.0
		if ix.layers >= layerKind {
			v = ix.stemVerdict(k, rule.Piece{Form: right, Tail: tail, Count: countOf(w, k, i), Sole: sole, Misspelt: true})
		}
		if v > 0 {
			names, best = append(names, kindLabel[k]), max(best, v)
		}
	}
	if len(names) == 0 {
		return "", 0
	}
	return "열쇠:" + strings.Join(names, "|"), best
}
