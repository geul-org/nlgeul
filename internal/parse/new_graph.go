//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what newGraph — the grammar of nlgeul parse as a toulmin graph: fifteen claims of a constituent, and the exceptions and rebuttals against them; the claims are returned so that more can be added
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// newGraph judges one reading of a word: its particle chain together
// with one of its tails. The claims are returned by name. The rule
// functions are the rule package's, told apart by their criteria; the
// names below are the nodes'.
//
//	chain claims  subjectCase, topicSubject → 주어; objectCase,
//	              topicObject → 목적어; complementCase → 보어;
//	              adnominalCase → 관형어; adverbialCase,
//	              rangeAdverbial → 부사어; vocativeCase → 독립어
//	tail claims   closingPredicate, lastPredicate → 서술어;
//	              adnominalClosing → 관형어; adverbialClosing → 부사어
//	word claims   determinerWord → 관형어; conjunctiveWord → 독립어
//	exceptions    finalWord → every chain and word claim,
//	              adnominalClosing, adverbialClosing, chainReadsMore,
//	              chainSeenMore
//	              chainReadsMore, tailCutShort → every tail claim
//	              tailReadsMore → every chain claim
//	              wholeWord → every chain and tail claim
//	              beforeBound → every chain claim, chainSeenMore
//	              boundStem → adnominalClosing, tailSeenMore
//	rebuttals     chainSeenMore → every tail claim
//	              tailSeenMore → every chain claim
//	              complementShare → subjectCase
func newGraph() (*toulmin.Graph, map[string]*toulmin.Rule) {
	g := toulmin.NewGraph("nlgeul:parse")
	c := map[string]*toulmin.Rule{}
	list := toulmin.NewStringListSpec
	claim(g, c, "subjectCase", labelSubj, rule.ChainCase, &rule.Bits{Of: "case", Mask: caseSubj})
	claim(g, c, "topicSubject", labelSubj, rule.ChainNoCase)
	claim(g, c, "objectCase", labelObj, rule.ChainCase, &rule.Bits{Of: "case", Mask: caseObj})
	claim(g, c, "topicObject", labelObj, rule.ChainNoCase)
	claim(g, c, "complementCase", labelComp, complementCase)
	claim(g, c, "adnominalCase", labelAdn, rule.ChainCase, &rule.Bits{Of: "case", Mask: caseAdn})
	claim(g, c, "adverbialCase", labelAdv, rule.ChainCase, &rule.Bits{Of: "case", Mask: caseAdv})
	claim(g, c, "vocativeCase", labelVoc, rule.ChainCase, &rule.Bits{Of: "case", Mask: caseVoc})
	claim(g, c, "closingPredicate", labelPred, rule.ClosingKind, &rule.Bits{Of: "kind", Mask: kindFinal | kindLink})
	claim(g, c, "lastPredicate", labelPred, lastPredicate)
	claim(g, c, "adnominalClosing", labelAdn, rule.ClosingKind, &rule.Bits{Of: "kind", Mask: kindAdn})
	claim(g, c, "adverbialClosing", labelAdv, rule.ClosingListed, list("adverbialClosings", adverbialClosings))
	claim(g, c, "rangeAdverbial", labelAdv, rule.ChainListed, list("rangeParticles", rangeParticles))
	claim(g, c, "determinerWord", labelAdn, rule.WordListed, list("determiners", determiners))
	claim(g, c, "conjunctiveWord", labelVoc, rule.WordListed, list("conjunctions", conjunctions))

	last := g.Except(rule.WordIs).With(&rule.Bits{Of: "word", Mask: rule.WordLast})
	chainLonger := g.Except(rule.Longer).With(&rule.Versus{Mine: rule.SideTail, Rival: rule.SideChain})
	tailLonger := g.Except(rule.Longer).With(&rule.Versus{Mine: rule.SideChain, Rival: rule.SideLongest})
	cut := g.Except(rule.Longer).With(&rule.Versus{Mine: rule.SideTail, Rival: rule.SideLongest})
	chainSeen, tailSeen := g.Counter(rule.ChainSeenMore), g.Counter(rule.TailSeenMore)
	whole := g.Except(rule.WordListed).With(list("wholeWords", append(append([]string{}, determiners...), conjunctions...)))
	modifies := g.Except(beforeBound)
	for _, name := range chainClaims {
		last.Attacks(c[name])
		tailLonger.Attacks(c[name])
		tailSeen.Attacks(c[name])
		whole.Attacks(c[name])
		modifies.Attacks(c[name])
	}
	for _, name := range tailClaims {
		chainLonger.Attacks(c[name])
		chainSeen.Attacks(c[name])
		cut.Attacks(c[name])
		whole.Attacks(c[name])
	}
	last.Attacks(c["adnominalClosing"])
	last.Attacks(c["adverbialClosing"])
	last.Attacks(chainLonger)
	last.Attacks(chainSeen)
	last.Attacks(c["determinerWord"])
	last.Attacks(c["conjunctiveWord"])
	modifies.Attacks(chainSeen)
	noStem := g.Except(rule.WordIs).With(&rule.Bits{Of: "word", Mask: rule.WordBound})
	noStem.Attacks(c["adnominalClosing"])
	noStem.Attacks(tailSeen)
	g.Counter(complementShare).Attacks(c["subjectCase"])
	return g, c
}
