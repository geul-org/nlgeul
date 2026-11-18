//ff:func feature=token type=rule control=sequence
//ff:what priorNounRule — the rule of a sentence "전 after a naming form or a quantity is the noun": a candidate that reads 전 (priorStem) as the noun of the word table (열쇠:체언) is preferred, by bonusOverCount, when the word right before it ends in a naming form (-기, -음: priorKinds) or in a quantity (a unit or an approximation after a number), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The standard dictionary gives 전(前) as a noun used after a -기 form
// (밥을 먹기 전) and after a noun of time (3년 전, 10분 전): the time
// before what is named. 전 the determiner (전 장관, 전 세계) modifies
// the noun after it and is not bound to what stands before it. The word
// table counts the determiner more often, and the rule of a word of two
// kinds standing alone puts it first; a naming form or a number with its
// unit right before 전 shows the noun.
func priorNounRule() sentenceRule {
	g, claim := sentenceClaim(rulePriorNoun, rule.Listed, listed("전", rule.ListSelfStem, []string{priorStem}))
	except(g, rule.KindListed, whyNoPrior, listed("이름꼴·수량", rule.ListPrevLast, priorKinds), not).Attacks(claim)
	return sentenceRule{name: rulePriorNoun, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool {
		return n.Stem == priorStem && n.Kind == tableNoun
	}}
}
