//ff:func feature=token type=rule control=sequence
//ff:what newIndex — build the token index: the closed-set words, then a graph for every reading of every particle, closing and slot of the tables, and the graph of a new word
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A token sequence that no table lists is a new word: the claim that
// it is a word stands, rebutted at half force by "it is in no key
// table", which always holds while there is no key table. Its verdict
// (1/3) is a low positive one, as stream-graph-algorithm.md section 12
// asks; the force is a placeholder, not a measured value.
func newIndex() *index {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}}
	ix.addWords()
	ix.addParticles()
	ix.addClosings()
	ix.addSlots()
	g := toulmin.NewGraph(newWord)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: newWord, Concludes: "열쇠"})
	g.Counter(rule.PieceRead).With(&rule.Claim{Name: noKey, Concludes: "반박"}).Attacks(claim, 0.5)
	ix.stem, ix.graphs, ix.edges = g, ix.graphs+1, ix.edges+1
	ix.unknown, _ = verdictOf(g, rule.Piece{}, "")
	return ix
}
