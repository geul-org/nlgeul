//ff:func feature=parse type=rule control=sequence
//ff:what namedBeforeHaRule — the rule of a sentence "the -기 form with a particle right before 하다": a candidate that reads its stem as a predicate's (predStem), one of whose markers is the naming ending -기 (haNamings) and whose last piece is not that ending (a particle follows it: 기로, 기도, 기는) is preferred, by bonusOverCount, when the word right after it is 하다 (haStem), and not otherwise; it waits for the word after it
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 하다 after the -기 form with a particle as a frame
// of its own (-기로 하다: what one decides; -기는 하다, -기도 하다: what is
// granted or added). The same letters read as a noun with the particle
// (속기+로, 서기+도) come first by the word table's counts where nothing
// is read after them. haLinkedRule reads the same frame from the other
// side (하다 right after a -기 form with a particle).
func namedBeforeHaRule() sentenceRule {
	g, claim := sentenceClaim(ruleNamedBeforeHa, rule.SelfHas, listed("이름꼴 기", rule.ListSelfMarks, haNamings))
	except(g, rule.Listed, whyNoDoNext, listed("하다", rule.ListNextStem, []string{haStem}), not).Attacks(claim)
	is := func(n rule.Near) bool {
		return predStem(n) && hasMark(n, haNamings) && labelOf(n.Last) != closingLabel(rule.KindNoun)
	}
	return sentenceRule{name: ruleNamedBeforeHa, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
