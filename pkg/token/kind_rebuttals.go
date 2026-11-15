//ff:func feature=token type=rule control=sequence
//ff:what kindRebuttals — on the graph of a kind of stem, the two rebuttals of its claim and what lifts each: "rarely seen", lifted for a spoken form of the closed table (spoken: 거), and "one syllable long", lifted when the stem's word is the whole of its unit and when the stem is an auxiliary verb's written right after the 어/아 form of its main verb with a predicate tail; the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A word of one syllable is rebutted because its letters are easily
// met by chance inside a longer run of letters; where the spaces set
// both ends of the word there is no such run, nor where clause 47 of the
// orthography lets an auxiliary verb be written onto its main verb: the
// main verb's 어/아 sets the front end of the auxiliary's stem and the
// predicate tail its back end.
func kindRebuttals(g *toulmin.Graph, claim *toulmin.Rule) int {
	few := g.Counter(rule.SeenFew).With(&rule.Bits{Of: "enough", Mask: seenEnough})
	few.Attacks(claim, forceFew)
	except(g, rule.Listed, whySpoken, listed("spoken", rule.ListForm, spoken)).Attacks(few)
	short := g.Counter(rule.ShortKey).With(&rule.Bits{Of: "syllables", Mask: shortKey})
	short.Attacks(claim, forceShort)
	except(g, rule.PieceIs, whySole, fact(rule.PieceSole)).Attacks(short)
	except(g, rule.PieceIs, whyConn, fact(rule.PieceConn)).Attacks(short)
	return 5
}
