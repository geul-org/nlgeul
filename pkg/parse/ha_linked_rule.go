//ff:func feature=parse type=rule control=sequence
//ff:what haLinkedRule — the rule of a sentence "하다 after -어야 or after a -기 form is the verb": a candidate that reads its stem as the verb 하 (haStem) with a closing is preferred, by bonusShortTaken, when the word right before it ends in the linking ending -어야 (haLinks) or has the naming ending -기 among its markers (haNamings: -기는, -기도, -기로, -기만), and not otherwise
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 하다 after -어야 (해야 한다, 갖추어야 하나: what
// must be) and after a -기 form with a particle (있기는 하나, 가기로
// 했다, 웃기만 한다: 하다 standing for the predicate named before it) as
// the verb that completes the construction. Read alone, the same
// letters as the numeral 하나 or as the determiner 한 come first by the
// word table's counts, by up to what bonusShortTaken carries.
func haLinkedRule() sentenceRule {
	g, claim := sentenceClaim(ruleHaLinked, rule.Listed, listed("하다", rule.ListSelfStem, []string{haStem}))
	far := except(g, rule.Listed, whyNoHaLink, listed("-어야", rule.ListPrevLast, haLinks), not)
	far.Attacks(claim)
	except(g, rule.PrevHas, whyNamedBefore, listed("이름꼴 -기", rule.ListPrevLast, haNamings)).Attacks(far)
	return sentenceRule{name: ruleHaLinked, graph: g, bonus: bonusShortTaken, about: func(n rule.Near) bool {
		return n.Stem == haStem && n.Kind == haName && n.Pred
	}}
}
