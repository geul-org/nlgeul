//ff:func feature=token type=rule control=iteration dimension=1
//ff:what misspelt — the right form of a misspelt stem as a known word with what follows it: its name and its verdict, the best among the kinds both the right form (in the word table or as a closed-set word) and the misspelt form have (0 when no kind takes that tail); each kind's graph is told that the stem is a misspelt form; a closed-set word is named by its set when it stands as that set's kind
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// It is known's second layer for a misspelt stem: in the first layer
// every kind counts as standing. The right form's own count is the
// stem's (the count of a misspelt line is 0, nothing was counted).
// A right form that is a closed-set word has the kind of its set too,
// and is named by the set as known names it.
func (ix *index) misspelt(right string, kinds, tail int, sole bool) (string, float64) {
	w := ix.lex[right]
	have := ix.closed[right] | w.kinds
	var names []string
	best, passed := 0.0, 0
	for i, k := range kindOrder {
		if have&kinds&k == 0 {
			continue
		}
		v := 1.0
		if ix.layers >= layerKind {
			v = ix.stemVerdict(k, rule.Piece{Form: right, Tail: tail, Count: countOf(w, k, i), Sole: sole, Misspelt: true})
		}
		if v > 0 {
			names, best, passed = append(names, kindLabel[k]), max(best, v), passed|k
		}
	}
	if len(names) == 0 {
		return "", 0
	}
	if own := ownName(ix.words[right], passed); own != "" {
		return own, best
	}
	return "열쇠:" + strings.Join(names, "|"), best
}
