//ff:func feature=token type=rule control=sequence
//ff:what addAfter — one piece that stands after a number (여, a unit, 째), kept by its last letter: the claim, and the exceptions that no number stands right before it and that letters run on right after it
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addAfter(name, form string, place int) {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "표지"})
	g.Except(rule.PieceNoNumber).Attacks(claim)
	g.Except(rule.PieceRunsOn).Attacks(claim)
	f := []rune(form)
	key := f[len(f)-1]
	ix.after[key] = append(ix.after[key], &entry{name: name, form: form, slot: place, g: g})
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+2
}
