//ff:func feature=token type=rule control=iteration dimension=2
//ff:what addWords — the closed-set words as keys: one graph each, with the exception that the word's set does not take what follows it (read when there is no word table); and this layer's own pronouns 네, 내 and 제, which are nominals wherever the kinds of a word are asked (with a word table)
package token

import (
	"sort"

	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A word that is in two sets keeps both names, in the order of the
// sets' names, so that the same input always prints the same, and
// takes what either set takes. The graphs and edges are counted once
// for each word.
func (ix *index) addWords() {
	sets := parse.ClosedWords()
	kinds := make([]string, 0, len(sets))
	for kind := range sets {
		kinds = append(kinds, kind)
	}
	sort.Strings(kinds)
	ix.takes = map[string]int{}
	for _, kind := range kinds {
		for _, w := range sets[kind] {
			name := kind + ":" + w
			if old, ok := ix.words[w]; ok {
				name = old.name + "|" + name
			}
			ix.closed[w] |= closedKinds[kind]
			ix.takes[w] |= closedTails[kind]
			g := toulmin.NewGraph(name)
			claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
			g.Except(rule.StemTail).With(&rule.Bits{Of: "tail", Mask: ix.takes[w]}).Attacks(claim)
			ix.words[w] = &entry{name: name, form: w, class: classKey, g: g}
		}
	}
	ix.graphs, ix.edges = ix.graphs+len(ix.words), ix.edges+len(ix.words)
	for _, w := range pronounWords {
		ix.closed[w] |= kindNominal
	}
}
