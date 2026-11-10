//ff:func feature=token type=rule control=iteration dimension=2
//ff:what addWords — the closed-set words as keys: one graph each, whose claim nothing attacks
package token

import (
	"sort"

	"github.com/geul-org/koparse/internal/parse"
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A word that is in two sets keeps both names, in the order of the
// sets' names, so that the same input always prints the same.
func (ix *index) addWords() {
	sets := parse.ClosedWords()
	kinds := make([]string, 0, len(sets))
	for kind := range sets {
		kinds = append(kinds, kind)
	}
	sort.Strings(kinds)
	for _, kind := range kinds {
		for _, w := range sets[kind] {
			name := kind + ":" + w
			if old, ok := ix.words[w]; ok {
				name = old.name + "|" + name
			}
			g := toulmin.NewGraph(name)
			g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
			ix.words[w] = &entry{name: name, form: w, class: classKey, g: g}
			ix.graphs++
		}
	}
}
