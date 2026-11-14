//ff:func feature=token type=rule control=sequence
//ff:what judged — the verdict of a graph for a piece and the letters of its word, with no trace kept; the same question is judged once and remembered, and what is remembered is dropped when it has grown to judgedMax
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The key holds everything verdictOf hands the graph (the piece and the
// word's letters), so that an answer is never given to another question.
// Dropping what is remembered changes no verdict: the graph is asked again.
func (ix *index) judged(g *toulmin.Graph, p rule.Piece, word string) float64 {
	key := judgeKey{g, p, word}
	if v, ok := ix.memo[key]; ok {
		return v
	}
	v, _ := verdictOf(g, p, word, false)
	if ix.memo == nil || len(ix.memo) >= judgedMax {
		ix.memo = map[judgeKey]float64{}
	}
	ix.memo[key] = v
	return v
}
