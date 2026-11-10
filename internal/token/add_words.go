//ff:func feature=token type=rule control=iteration dimension=2
//ff:what addWords — the closed-set words as keys: one graph each, whose claim nothing attacks
package token

import (
	"github.com/geul-org/koparse/internal/parse"
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A word that is in two sets keeps both names.
func (ix *index) addWords() {
	for kind, words := range parse.ClosedWords() {
		for _, w := range words {
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
