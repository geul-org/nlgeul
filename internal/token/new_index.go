//ff:func feature=token type=rule control=sequence
//ff:what newIndex — build the token index: the closed-set words, then a graph for every reading of every particle, closing and slot of the tables, the graph of a new word, and with a word table its keys and a graph for each kind of stem
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A token sequence that no table lists is a new word: the claim that
// it is a word stands, rebutted at half force by "it is in no key
// table". Its verdict (1/3) is a low positive one, as
// stream-graph-algorithm.md section 12 asks; the force is a
// placeholder, not a measured value. Without a word table the index
// reads as it did before there was one.
func newIndex(o Options) (*index, error) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addWords()
	ix.addParticles()
	ix.addClosings()
	ix.addSlots()
	g := toulmin.NewGraph(newWord)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: newWord, Concludes: "열쇠"})
	g.Counter(rule.PieceRead).With(&rule.Claim{Name: noKey, Concludes: "반박"}).Attacks(claim, 0.5)
	ix.stem, ix.graphs, ix.edges = g, ix.graphs+1, ix.edges+1
	ix.unknown, _ = verdictOf(g, rule.Piece{}, "")
	if o.Words == "" {
		return ix, nil
	}
	lex, total, err := loadWords(o.Words)
	if err != nil {
		return nil, err
	}
	ix.lex, ix.lexTotal, ix.layers = lex, total, layerCost
	if o.Layers >= layerKnown && o.Layers <= layerCost {
		ix.layers = o.Layers
	}
	ix.addKinds()
	return ix, nil
}
