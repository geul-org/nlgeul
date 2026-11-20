//ff:func feature=parse type=rule control=sequence
//ff:what addSurname — the graph of a family name of the list (one for all of them): a claim with no rebuttal, so that a family name read before a title stands with verdict 1
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The family names are a closed set (spacing-words.tsv), as numbers
// are: the graph has no rebuttal.
func (ix *index) addSurname() {
	g := toulmin.NewGraph(setSurname)
	g.Rule(rule.PieceRead).With(&rule.Claim{Name: setSurname, Concludes: "열쇠"})
	ix.surname = verdict(g, rule.Piece{})
	ix.graphs++
}
