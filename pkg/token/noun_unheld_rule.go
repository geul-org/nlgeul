//ff:func feature=token type=rule control=sequence
//ff:what nounUnheldRule — the rule of a sentence "the letters of a dependent noun with no modifier and no nominal right before them are a noun": a candidate that reads its stem, one of the words given (unheldBound: 시, 간, 적), as a noun of the word table (열쇠:체언) is preferred when the word right before it ends neither in a modifying closing or the adnominal case particle 의 nor in a nominal, nor in a suffix (피고들 간), the 여 of a rough number (5분여 만) or a closing sign, or is a closing sign itself (20일)간, 선조(宣祖) 때), and is no determiner of the closed set (따라 시와, 단백질은 간에서), and not otherwise; at the head of a sentence and right after an opening sign it is not asked (nounHeadedRule)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A dependent noun is used only with a modifier before it (school
// grammar: 의존 명사는 관형어의 꾸밈을 받아야 쓰인다: a predicate's modifying
// form, a nominal with 의, a determiner), and some of them stand right
// after a nominal as well (bound-nouns.tsv: 서울과 부산 간, 출발 시). After
// a word that ends in another particle, in a closing that does not
// modify, or in an adverb, the same letters are the noun (시 of 詩 or
// 市, 간 of the liver, 적 of an enemy). A suffix (들), the 여 of a rough
// number and a closing sign end a nominal as well: the word before such
// a sign is what the sign closes.
func nounUnheldRule(words []string) sentenceRule {
	noun := keyKind + ":" + kindLabel[kindNominal]
	g, claim := sentenceClaim(ruleNounUnheld, rule.Listed, listed("관형어나 체언 뒤에 서는 의존 명사", rule.ListSelfStem, words))
	except(g, rule.KindListed, whyModifierBefore, listed("관형어", rule.ListPrevLast, []string{closingLabel(rule.KindAdn), adnCase})).Attacks(claim)
	except(g, rule.Listed, whyDet, listed("관형사", rule.ListPrevKind, []string{determinerLabel})).Attacks(claim)
	except(g, rule.AroundIs, whyAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}).Attacks(claim)
	except(g, rule.KindListed, whyHeldBefore, listed("접미·어림·닫는 기호", rule.ListPrevLast, heldLasts)).Attacks(claim)
	except(g, rule.Listed, whyShutBefore, listed("닫는 기호", rule.ListPrevKind, []string{shutName})).Attacks(claim)
	except(g, rule.Headed, whyHeaded, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(claim)
	return sentenceRule{name: ruleNounUnheld, graph: g, bonus: bonusSameLetters, about: func(n rule.Near) bool {
		return n.Kind == noun && slices.Contains(words, n.Stem)
	}}
}
