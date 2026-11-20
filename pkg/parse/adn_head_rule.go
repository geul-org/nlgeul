//ff:func feature=parse type=rule control=sequence
//ff:what adnHeadRule — the rule of rebuttal of a sentence "an adnominal stands right before a word it modifies", judged on the adnominal: a candidate that is an adnominal (adnominalNear: a modifying form, a determiner standing alone) does not stand (forceUngrammatical) when no adnominal may stand right before the word after it (AroundBeforeUnmodifiable: each of its known candidates is a predicate ending otherwise, an adverb or an interjection, or it is a mark that ends a sentence); not when that word's stem is an auxiliary of a dependent noun and 하다 (boundAux: 올 듯하다) or an adverb of the closed set 꾸밈 사이 부사 (adnAdverbs: 이른바, 가장), nor when it ends in -고 or -며 (coordLinks: 그 맥없고 가련한), nor when the word after it ends in a modifying form (a clause of adnominals begins with it); it waits for the words after it
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// School grammar: an adnominal (관형어: a determiner, a predicate's
// modifying form) modifies the nominal after it. Another adnominal may
// stand in between (새로운 그 길), and so may a clause of adnominals that
// begins with a predicate's -고 or -며 form, an adverb or a linking form
// (있는 가장 큰 섬, 다른 남아 있는 것); the auxiliaries made of a dependent
// noun and 하다 or 싶다 follow a modifying form (올 듯하다). A predicate
// that ends otherwise, an adverb, an interjection and the end of a
// sentence are none of these: an adnominal reading right before such a
// word is not one the grammar lets stand. When the word after it has a
// reading an adnominal may modify, this rule does not judge the
// adnominal: the rule of the word after it does (adnHeadAfterRule).
func adnHeadRule() sentenceRule {
	g, claim := sentenceRebuttal(ruleAdnHead, rule.AroundIs, &rule.Bits{Of: "around", Mask: rule.AroundBeforeUnmodifiable})
	except(g, rule.Listed, whyBoundAuxNext, listed("의존명사 보조 용언", rule.ListNextStem, boundAux)).Attacks(claim)
	except(g, rule.Listed, whyCoordNext, listed("-고·-며", rule.ListNextLast, coordLinks)).Attacks(claim)
	except(g, rule.Listed, whyAdnAdverbNext, listed("꾸밈 사이 부사", rule.ListNextStem, adnAdverbs)).Attacks(claim)
	except(g, rule.KindListed, whyAdnAfterNext, listed("꾸밈 맺음", rule.ListNext2Last, []string{closingLabel(rule.KindAdn)})).Attacks(claim)
	return sentenceRule{name: ruleAdnHead, graph: g, force: forceUngrammatical, about: adnominalNear, waits: adnominalNear}
}
