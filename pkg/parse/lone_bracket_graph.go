//ff:func feature=parse type=rule control=sequence
//ff:what loneBracketGraph — the judgment graph of a closing bracket that closes a pair with none of its kind open (loneBrackets): the claim "this sign closes a pair", and the exceptions against it: a sign of its kind may stand open in some way (the usual closing reads it then), no letter stands right before it (itself defeated when a mark that ends a sentence stands there: 다.』), a digit stands right before it (a list number), and what the sign's row bars (apart with a space on both sides, or next to a hyphen or an equals sign); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Its name is that of the usual closing (닫기:)), as a lone quotation
// mark's is: the two never stand in the same way.
func loneBracketGraph(sign rune, row pairRow) (*toulmin.Graph, int) {
	full := shutName + ":" + string(sign)
	g := toulmin.NewGraph(full)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: full, Concludes: "기호"})
	except(g, rule.MarkIs, whyOpenedAny, &rule.Bits{Of: "mark", Mask: rule.MarkOpenedAny}).Attacks(claim)
	bare := except(g, rule.MarkIs, whyNotOnWord, &rule.Bits{Of: "mark", Mask: rule.MarkAfterWord}, not)
	bare.Attacks(claim)
	except(g, rule.MarkIs, whyEndsSentence, &rule.Bits{Of: "mark", Mask: rule.MarkAfterStop}).Attacks(bare)
	except(g, rule.MarkIs, whyAfterDigit, &rule.Bits{Of: "mark", Mask: rule.MarkAfterDigit}).Attacks(claim)
	edges := 4
	if row.bars&barApart != 0 {
		except(g, rule.MarkIs, whyApart, &rule.Bits{Of: "mark", Mask: rule.MarkApart}).Attacks(claim)
		except(g, rule.MarkIs, whyArrow, &rule.Bits{Of: "mark", Mask: rule.MarkArrow}).Attacks(claim)
		edges += 2
	}
	return g, edges
}
