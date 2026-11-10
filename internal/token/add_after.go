//ff:func feature=token type=rule control=sequence
//ff:what addAfter — one piece that stands after a number (여, a unit, 째), kept by its last letter: the claim, the exceptions that no number stands right before it and that letters run on right after it, and for all but a greater number the exception that the number is a lone Sino-Korean digit with no word table and the rebuttal that the number and the piece together are a word
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
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+2
	if place != placeBig {
		g.Except(rule.PieceLone).Attacks(claim)
		g.Counter(rule.PieceWhole).Attacks(claim, forceWhole)
		ix.edges += 2
	}
	f := []rune(form)
	key := f[len(f)-1]
	ix.after[key] = append(ix.after[key], &entry{name: name, form: form, slot: place, g: g})
}
