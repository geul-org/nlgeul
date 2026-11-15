//ff:func feature=token type=rule control=sequence
//ff:what loneShutGraph — the judgment graph of a quotation mark that closes a pair with none of its kind open (loneShuts): the claim "this sign closes a pair", and the exceptions against it: a sign of its kind may stand open in some way (the usual closing reads it then), no letter stands right before it, no Hangul syllable stands right after it, and what the sign's row bars inside a word; the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Its name is that of the usual closing (닫기:"): the two never stand in
// the same way, as one asks that a sign of its kind stands open and the
// other that none may.
func loneShutGraph(sign rune, row pairRow) (*toulmin.Graph, int) {
	full := shutName + ":" + string(sign)
	g := toulmin.NewGraph(full)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: full, Concludes: "기호"})
	except(g, rule.MarkIs, whyOpenedAny, &rule.Bits{Of: "mark", Mask: rule.MarkOpenedAny}).Attacks(claim)
	except(g, rule.MarkIs, whyNotOnWord, &rule.Bits{Of: "mark", Mask: rule.MarkAfterWord}, not).Attacks(claim)
	except(g, rule.MarkRight, whyNoHangulOn, &rule.Bits{Of: "sides", Mask: rule.SideHangul}).Attacks(claim)
	edges := 3
	if row.bars&barWord != 0 {
		except(g, rule.MarkBoth, whyInWord, &rule.Bits{Of: "sides", Mask: rule.SideAlnum}).Attacks(claim)
		edges++
	}
	return g, edges
}
