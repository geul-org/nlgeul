//ff:func feature=token type=rule control=sequence
//ff:what newIndex — build the token index: the closed-set words, then a graph for every reading of every particle, closing and slot of the tables, the graph of a new word (with its exceptions: letters that read as markers, a stem that ends in ㅆ before markers), and with a word table its keys and a graph for each kind of stem
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A token sequence that no table lists is a new word: the claim that
// it is a word stands, rebutted at half force by "it is in no key
// table". Its verdict (1/3) is a low positive one, as
// token-stream-algorithm.md section 12 asks; the force is a
// placeholder, not a measured value. A whole new word after a boundary
// inside a unit does not stand when its letters read as markers there
// (the exception PieceMarked). A new word with markers after it does
// not stand when it ends in the final consonant ㅆ (the exception
// StemCoda): such a stem is 있 or a word made with it, which a word
// table lists, and any other ㅆ is the slot 았/었/겠.
func newIndex(o Options) (*index, error) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}, lexed: o.Words != ""}
	ix.addWords()
	ix.addParticles()
	ix.addClosings()
	ix.addSlots()
	if o.Words != "" {
		ix.addWide()
	}
	ix.addCopula()
	ix.addFused()
	ix.addDots()
	ix.addNumbers()
	g := toulmin.NewGraph(newWord)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: newWord, Concludes: "열쇠"})
	g.Counter(rule.PieceRead).With(&rule.Claim{Name: noKey, Concludes: "반박"}).Attacks(claim, 0.5)
	g.Except(rule.PieceMarked).Attacks(claim)
	g.Except(rule.StemCoda).With(&rule.Bits{Of: "jong", Mask: jongSS}).Attacks(claim)
	ix.stem, ix.graphs, ix.edges = g, ix.graphs+1, ix.edges+3
	ix.unknown, _ = verdictOf(g, rule.Piece{}, "")
	if o.Words == "" {
		return ix, nil
	}
	lex, total, err := loadWords(o.Words)
	if err != nil {
		return nil, err
	}
	ix.lex, ix.lexTotal, ix.layers = lex, total, layerCost
	ix.addAlts()
	if o.Layers >= layerKnown && o.Layers <= layerCost {
		ix.layers = o.Layers
	}
	ix.addKinds()
	return ix, nil
}
