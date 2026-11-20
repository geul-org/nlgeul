//ff:func feature=parse type=rule control=sequence
//ff:what kindRebuttals — on the graph of a kind of stem, the two rebuttals of its claim, "rarely seen" (below the count enough) and "one syllable long", and what lifts the second: it is lifted when the stem's word is the whole of its unit, when the stem is an auxiliary verb's written right after the 어/아 form of its main verb with a predicate tail, when the stem is that main verb's, beginning its unit, and it is the dependent noun 것 or 거 written right after a modifying closing; the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A word of one syllable is rebutted because its letters are easily
// met by chance inside a longer run of letters; where the spaces set
// both ends of the word there is no such run, nor where clause 47 of the
// orthography lets an auxiliary verb be written onto its main verb: the
// main verb's 어/아 sets the front end of the auxiliary's stem and the
// predicate tail its back end; nor at the main verb itself when it
// begins its unit: the space sets its front end, and its 어/아 with the
// auxiliary's stem right after it its back end (가봐, 해줘). Nor is there such a run where the
// dependent noun 것 (or its spoken form 거) is written right onto the
// modifying closing it stands on (하는거, 없는게, 할것이): a dependent noun
// stands only after a modifier, so that the closing sets its front end,
// and its own markers or the unit's end set its back end.
func kindRebuttals(g *toulmin.Graph, claim *toulmin.Rule, enough int) int {
	few := g.Counter(rule.SeenFew).With(&rule.Bits{Of: "enough", Mask: enough})
	few.Attacks(claim, forceFew)
	short := g.Counter(rule.ShortKey).With(&rule.Bits{Of: "syllables", Mask: shortKey})
	short.Attacks(claim, forceShort)
	except(g, rule.PieceIs, whySole, fact(rule.PieceSole)).Attacks(short)
	except(g, rule.PieceIs, whyConn, fact(rule.PieceConn)).Attacks(short)
	except(g, rule.PieceIs, whyAdnBound, fact(rule.PieceBound)).Attacks(short)
	except(g, rule.PieceIs, whyMain, fact(rule.PieceMain)).Attacks(short)
	return 6
}
