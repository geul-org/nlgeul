//ff:func feature=token type=rule control=sequence
//ff:what verdict — evaluate one graph for one piece and give the claim's verdict alone (for the graphs that are asked without a place in a unit's trace: a stem's kind, a new word, a restored syllable)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func verdict(g *toulmin.Graph, p rule.Piece) float64 {
	v, _ := verdictOf(g, p, "", false)
	return v
}
