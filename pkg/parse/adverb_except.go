//ff:func feature=parse type=rule control=sequence
//ff:what adverbExcept — on the graph of an adverb's stem, against the exception that its kind does not take what follows, what lifts it: the copula follows (adverbLoose); and against that, that the tables give the stem as a nominal too, and that the stem is a run of Latin letters and digits (it is read as the run it is: off이다); the claim is rebutted for such a tail (forceAdverbLoose); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// An adverb takes auxiliary particles; the copula stands after an
// adverb as after a nominal (아직이다, 그래서일까). Such a reading is
// weaker than the adverb alone, so the claim is rebutted, and a stem
// that is a nominal as well takes the copula as the nominal (물론이다).
// A run of Latin letters and digits that the
// word table lists as an adverb is read before such a tail as the run it
// is (off[열:영숫자]+이+다).
func adverbExcept(g *toulmin.Graph, claim, tail *toulmin.Rule) int {
	loose := except(g, rule.StemTail, whyAdverbLoose, &rule.Bits{Of: "tail", Mask: ^adverbLoose})
	loose.Attacks(tail)
	except(g, rule.StemKind, whyAlsoNominal, &rule.Bits{Of: "kinds", Mask: kindNominal}).Attacks(loose)
	except(g, rule.StemRun, whyStemRun).Attacks(loose)
	g.Counter(rule.StemTail).With(&rule.Bits{Of: "tail", Mask: ^adverbLoose}).Attacks(claim, forceAdverbLoose)
	return 4
}
