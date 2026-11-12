//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addNumbers — the graphs of numbers: a run of digits, each reading of each sign between runs of digits (the full stop keeps its own and gets two more), the readings of a slash and a hyphen that join the runs of a name, the signs at a number's front, the range sign after a unit, and the pieces after a number (여, the units, 째 and the suffixes that stand where it stands)
package token

import (
	"math"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A number and the words of the unit table are closed sets: their
// graphs have no rebuttal, so they stand with verdict 1. The tables
// give the pieces after a number no counts: each is taken as one of
// the table's entries, as a token no table lists is taken as one of
// the Hangul syllables, and a greater number as one of its own few.
func (ix *index) addNumbers() {
	g := toulmin.NewGraph(numGraph)
	g.Rule(rule.PieceRead).With(&rule.Claim{Name: numGraph, Concludes: "열쇠"})
	ix.number = verdict(g, rule.Piece{})
	ix.graphs++
	ix.signs = map[rune][]*entry{'.': ix.joins}
	for _, sign := range signOrder {
		for _, s := range signShapes[sign] {
			ix.signs[sign] = append(ix.signs[sign], shapeGraph(s))
			ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
		}
	}
	for _, sign := range nameOrder {
		ix.signs[sign] = append(ix.signs[sign], nameGraph(nameSigns[sign]))
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+2
	}
	digits := &rule.Bits{Of: "sides", Mask: rule.SideDigit}
	ix.fronts = map[rune]*entry{}
	for _, sign := range headOrder {
		ix.fronts[sign] = frontGraph(headSigns[sign], digits, true)
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+2
	}
	ix.tilde = frontGraph(tildeName, digits, false)
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	ix.after = map[rune][]*entry{}
	ix.afterPrefs = map[bool]float64{false: math.Log10(1 / float64(len(unitRows)+afterOthers)), true: math.Log10(1 / float64(len([]rune(sinoBigs))))}
	for _, r := range sinoBigs {
		ix.addAfter(bigName+string(r), string(r), placeBig)
	}
	ix.addAfter(overName, "여", placeOver)
	ix.addAfter(orderName, "째", placeOrder)
	for _, s := range unitSuffixes {
		ix.addAfter("접미:"+s, s, placeOrder)
	}
	for _, row := range unitRows {
		ix.addAfter(unitLabel+row.Cells[1], row.Cells[0], placeUnit)
	}
}
