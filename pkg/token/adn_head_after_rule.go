//ff:func feature=token type=rule control=sequence
//ff:what adnHeadAfterRule — the rule of rebuttal of a sentence "an adnominal stands right before a word it modifies", judged on the word after it: a candidate no adnominal modifies (takesNoAdnominal: a predicate ending otherwise than in a modifying or naming form, an adverb, an interjection) that does not end in a linking closing (linkClosings) does not stand (forceUngrammatical) right after a word each of whose candidates is an adnominal (AroundAfterAdnominal); not when its stem is an auxiliary of a dependent noun and 하다 (boundAux: 올 듯하다), an adverb of the closed set 꾸밈 사이 부사 (adnAdverbs) or a negative adverb (negAdverbs: 몇 안 되는); it reads nothing after the candidate
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The other side of adnHeadRule: right after an adnominal the word read
// as one no adnominal modifies does not stand, and a reading of it an
// adnominal may modify (a nominal: 스무 채) comes first. The word before
// is read as its first candidate after it was judged; when it could be
// read as no adnominal as well, the grammar does not tell which of the
// two words yields, and the rule asks nothing. A predicate's linking
// form may begin a clause of adnominals (다른 남아 있는 것), and a
// negative adverb modifies the predicate right after it (몇 안 되는
// 사람): neither is judged.
func adnHeadAfterRule() sentenceRule {
	g, claim := sentenceRebuttal(ruleAdnHead, rule.AroundIs, &rule.Bits{Of: "around", Mask: rule.AroundAfterAdnominal})
	except(g, rule.Listed, whyBoundAuxSelf, listed("의존명사 보조 용언", rule.ListSelfStem, boundAux)).Attacks(claim)
	except(g, rule.Listed, whyAdnAdverbSelf, listed("꾸밈 사이 부사", rule.ListSelfStem, adnAdverbs)).Attacks(claim)
	except(g, rule.Listed, whyNegAdverbSelf, listed("부정 부사", rule.ListSelfStem, negAdverbs)).Attacks(claim)
	is := func(n rule.Near) bool { return takesNoAdnominal(n) && !slices.Contains(linkClosings, n.Last) }
	return sentenceRule{name: ruleAdnHead, graph: g, force: forceUngrammatical, about: is}
}
