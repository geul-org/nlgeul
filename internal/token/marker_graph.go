//ff:func feature=token type=rule control=sequence
//ff:what markerGraph — the judgment graph of one marker reading, stamped from its table row and its exception list: the claim "this token sequence is this marker" and the structural exceptions against it
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Every marker needs a stem before it. A whole-syllable form checks
// how the token before it ends; a particle whether its kind takes the particle after it; a suffix that needs a particle after it that one follows; a
// slot its order; a lone final consonant the syllable it sits under
// (with a word table the word before it is asked instead, where the
// stem is known: lexed); a piece that begins with 어/아/여 is rebutted
// when it does not fit the vowel before it; and a form that ends a
// closed-set word yields to that word. The number of edges is returned.
func markerGraph(e *entry, whole []string, lexed bool) int {
	g := toulmin.NewGraph(e.name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: e.name, Concludes: "표지"})
	g.Except(rule.PieceNoStem).Attacks(claim)
	edges := 1
	if e.lead == 0 && e.after != afterAny {
		g.Except(rule.PieceAfter).With(&rule.Bits{Of: "after", Mask: e.after}).Attacks(claim)
		edges++
	}
	if e.class == classParticle {
		g.Except(rule.ChainLinks).Attacks(claim)
		edges++
	}
	if e.inner {
		g.Except(rule.PieceLast).Attacks(claim)
		edges++
	}
	if e.class == classSlot {
		g.Except(rule.SlotOrder).Attacks(claim)
		edges++
	}
	if e.class == classClosing && e.form == "" && e.lead > 0 && !lexed {
		g.Except(rule.PieceBase).With(toulmin.NewStringListSpec("bases", rule.BareBases(e.lead))).Attacks(claim)
		edges++
	}
	if e.wide && e.class == classClosing {
		g.Except(rule.PieceParticle).Attacks(claim)
		edges++
	}
	if e.head&headWhole != 0 {
		g.Counter(rule.PieceClash).Attacks(claim, forceClash)
		edges++
	}
	if len(whole) > 0 {
		g.Except(rule.WordListed).With(toulmin.NewStringListSpec("wholeWords", whole)).Attacks(claim)
		edges++
	}
	e.g = g
	return edges
}
