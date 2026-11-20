//ff:func feature=parse type=rule control=sequence
//ff:what malLinkedRule — the rule of a sentence "말다 after -지 or -고 is the auxiliary verb": a candidate whose stem is 말 read as the verb (malStem, mainKindOf) is preferred when the word before it ends in the closing 지 or 고 (malLinks), or in such a closing with an auxiliary particle after it (particleEnds), and not otherwise
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 말다 as an auxiliary verb in -지 말다 (가지 마라)
// and -고 말다 (지고 말았다). Its forms that drop the ㄹ (마, 마라,
// 마시고, 마소서) are also other words' letters (마, 마라, 마시다, 마소),
// which come first without the word before. A particle may stand
// between the two (clause 47 of the orthography, 다만: 따내고야 말,
// 하지도 마라): the word before then has 지 or 고 among its markers and
// ends in the particle (particleEnds).
func malLinkedRule() sentenceRule {
	g, claim := sentenceClaim(ruleMalLinked, rule.Listed, listed("말다", rule.ListSelfStem, []string{malStem}))
	unlinked := except(g, rule.Listed, whyNotMalLinked, listed("지·고", rule.ListPrevLast, malLinks), not)
	unlinked.Attacks(claim)
	particled := except(g, rule.PrevHas, whyLinkParticle, listed("지·고", rule.ListPrevLast, malLinks))
	particled.Attacks(unlinked)
	except(g, rule.KindListed, whyNoParticleEnd, listed("보조사", rule.ListPrevLast, particleEnds), not).Attacks(particled)
	return sentenceRule{name: ruleMalLinked, graph: g, bonus: bonusShortTaken, about: func(n rule.Near) bool {
		return n.Stem == malStem && n.Kind == mainKindOf(malStem)
	}}
}
