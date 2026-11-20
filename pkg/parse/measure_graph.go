//ff:func feature=parse type=rule control=sequence
//ff:what measureGraph — the judgment graph of a quotation mark right after a digit read as a unit of measure (3", 5'): the claim "this sign is a unit", and the rebuttal that a sign of its kind stands open before it (then the sign closes that pair); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The orthography's appendix on punctuation writes quotation marks in
// pairs: a quotation mark of the kind left open before closes it. The
// feet, the inches, the minutes and the seconds have their own letters
// (′ ″); a straight quotation mark right after a digit is read as one
// of them only as a candidate, which is lowered, not taken away, where
// it would close a pair (the sign of a pair's own graph judges that).
func measureGraph() (*toulmin.Graph, int) {
	g := toulmin.NewGraph(measureName)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: measureName, Concludes: "기호"})
	g.Counter(rule.MarkIs).With(&rule.Bits{Of: "mark", Mask: rule.MarkOpened}).Attacks(claim, forceMeasureOpen)
	return g, 1
}
