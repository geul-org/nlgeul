//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what newGraph — the grammar of kopas parse as a toulmin graph: fifteen claims of a constituent, and the exceptions and rebuttals against them; the claims are returned so that more can be added
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// newGraph judges one reading of a word: its particle chain together
// with one of its tails. The claims are returned by function name.
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
	g := toulmin.NewGraph("kopas:parse")
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
		"rangeAdverbial":   g.Rule(rangeAdverbial).With(toulmin.NewStringListSpec("rangeParticles", rangeParticles)),
		"determinerWord":   g.Rule(determinerWord).With(toulmin.NewStringListSpec("determiners", determiners)),
		"conjunctiveWord":  g.Rule(conjunctiveWord).With(toulmin.NewStringListSpec("conjunctions", conjunctions)),
	}
	last := g.Except(finalWord)
	chainLonger, chainSeen := g.Except(chainReadsMore), g.Counter(chainSeenMore)
	tailLonger, tailSeen := g.Except(tailReadsMore), g.Counter(tailSeenMore)
	cut := g.Except(tailCutShort)
	whole := g.Except(wholeWord).With(toulmin.NewStringListSpec("wholeWords", append(append([]string{}, determiners...), conjunctions...)))
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
	noStem := g.Except(boundStem)
	noStem.Attacks(c["adnominalClosing"])
	noStem.Attacks(tailSeen)
	g.Counter(complementShare).Attacks(c["subjectCase"])
	return g, c
}
