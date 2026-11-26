//ff:func feature=parse type=rule control=sequence
//ff:what swapGraph — the judgment graph of a round quotation mark read against its shape, as a text that writes the two shapes the wrong way round means it: the claim "this sign opens a pair" for a mark drawn as a closing one (swapOpens: ”두), with the exception that it does not stand where an opening mark is written (MarkOpenSide); the claim "this sign closes a pair" for a mark drawn as an opening one (swapShuts: 사항‘이라고), with the exception that it does not stand where a closing mark is written (MarkShutSide); the name of the candidate and the number of edges are returned
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The reading by the shape stays a candidate beside it (pairGraph lowers
// it in the same place).
func swapGraph(sign rune) (*toulmin.Graph, string, int) {
	name, side, why := shutName, rule.MarkShutSide, whyNotShutSide
	if strings.ContainsRune(swapOpens, sign) {
		name, side, why = openName, rule.MarkOpenSide, whyNotOpenSide
	}
	full := name + ":" + string(sign)
	g := toulmin.NewGraph(full)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: full, Concludes: "기호"})
	except(g, rule.MarkIs, why, &rule.Bits{Of: "mark", Mask: side}, not).Attacks(claim)
	return g, name, 1
}
