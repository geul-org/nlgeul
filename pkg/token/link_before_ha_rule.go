//ff:func feature=token type=rule control=sequence
//ff:what linkBeforeHaRule — the rule of a sentence "the -어야 form right before 하다 or 되다": a candidate that reads its stem as a predicate's (predStem) and one of whose markers is the linking ending -어야 (haLinks: 어야, 아야, 여야, and the merged form) is preferred, by bonusOverCount, when the word right after it is 하다 or 되다 (haDoStems), and not otherwise; it waits for the word after it
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives -어야 with 하다 or 되다 after it as what must be
// done (세금을 내야 한다, 일찍 자야 된다: 표준국어대사전 -어야, '하다',
// '되다' 따위와 함께 쓰여 의무적으로 또는 당위적으로 그리하여야 함). The
// same letters read as a noun (내야, 가야) or as a noun with the
// particle 야 (기대+야, 자+야) come first by the word table's counts where
// nothing is read after them. haLinkedRule reads the same frame from the
// other side (하다 right after -어야).
func linkBeforeHaRule() sentenceRule {
	g, claim := sentenceClaim(ruleLinkBeforeHa, rule.SelfHas, listed("맺음 어야", rule.ListSelfMarks, haLinks))
	except(g, rule.Listed, whyNoHaNext, listed("하다·되다", rule.ListNextStem, haDoStems), not).Attacks(claim)
	is := func(n rule.Near) bool { return predStem(n) && hasMark(n, haLinks) }
	return sentenceRule{name: ruleLinkBeforeHa, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
