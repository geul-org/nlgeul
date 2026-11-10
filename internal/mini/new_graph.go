//ff:func feature=mini type=rule control=iteration dimension=1
//ff:what newGraph — the grammar of kopas mini as a toulmin graph: twelve claims of a constituent, and the exceptions and rebuttals against them; the claims are returned so that more can be added
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// newGraph judges one reading of a word: its particle chain together
// with one of its tails. The claims are returned by function name.
//
//	chain claims  subjectCase, topicSubject → 주어; objectCase,
//	              topicObject → 목적어; complementCase → 보어;
//	              adnominalCase → 관형어; adverbialCase → 부사어;
//	              vocativeCase → 독립어
//	tail claims   closingPredicate, lastPredicate → 서술어;
//	              adnominalClosing → 관형어; adverbialClosing → 부사어
//	exceptions    finalWord → every chain claim, adnominalClosing,
//	              adverbialClosing, chainReadsMore, chainSeenMore
//	              chainReadsMore, tailCutShort → every tail claim
//	              tailReadsMore → every chain claim
//	rebuttals     chainSeenMore → every tail claim
//	              tailSeenMore → every chain claim
//	              complementShare → subjectCase
func newGraph() (*toulmin.Graph, map[string]*toulmin.Rule) {
	g := toulmin.NewGraph("kopas:mini")
	c := map[string]*toulmin.Rule{
		"subjectCase":      g.Rule(subjectCase),
		"topicSubject":     g.Rule(topicSubject),
		"objectCase":       g.Rule(objectCase),
		"topicObject":      g.Rule(topicObject),
		"complementCase":   g.Rule(complementCase),
		"adnominalCase":    g.Rule(adnominalCase),
		"adverbialCase":    g.Rule(adverbialCase),
		"vocativeCase":     g.Rule(vocativeCase),
		"closingPredicate": g.Rule(closingPredicate),
		"lastPredicate":    g.Rule(lastPredicate),
		"adnominalClosing": g.Rule(adnominalClosing),
		"adverbialClosing": g.Rule(adverbialClosing).With(toulmin.NewStringListSpec("adverbialClosings", adverbialClosings)),
	}
	last := g.Except(finalWord)
	chainLonger, chainSeen := g.Except(chainReadsMore), g.Counter(chainSeenMore)
	tailLonger, tailSeen := g.Except(tailReadsMore), g.Counter(tailSeenMore)
	cut := g.Except(tailCutShort)
	for _, name := range chainClaims {
		last.Attacks(c[name])
		tailLonger.Attacks(c[name])
		tailSeen.Attacks(c[name])
	}
	for _, name := range tailClaims {
		chainLonger.Attacks(c[name])
		chainSeen.Attacks(c[name])
		cut.Attacks(c[name])
	}
	last.Attacks(c["adnominalClosing"])
	last.Attacks(c["adverbialClosing"])
	last.Attacks(chainLonger)
	last.Attacks(chainSeen)
	g.Counter(complementShare).Attacks(c["subjectCase"])
	return g, c
}
