//ff:func feature=parse type=rule control=sequence
//ff:what haOneRule — the rule of a sentence "한 after a complement of 하다 is 하다 with the modifying ㄴ": the candidate that reads the word 한 as the verb 하 of the word table with the modifying closing ㄴ is preferred, by bonusOverCount, when the word before it ends in a form that makes it a complement of 하다 (haTakes: the object particle, 로/으로, 에게/한테/께, -게, -도록, -어야, -려(고), -고자, -다고 …), and not otherwise, unless 한 is written onto the word before it (a determiner is written apart from it), but not right after an opening sign nor after the modifying closing 는 (where it is the noun 한(限), limitOneRule); nor when the word right after it is a unit a numeral determiner counts, of the units that are no dependent noun a modifying form stands before (countedUnitExcept, countedUnits: 한 번); it waits for the word after it
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 한 as a determiner (한 사람, 한 번) and as 하다
// with the modifying ㄴ; the word table's counts put the determiner first.
// 하다 is a transitive verb and an auxiliary one: it takes an object (일을
// 한), "…으로" (전제로 한), the one it is done or said to (그에게 한) and the clauses of -게/-도록/-어야/-려고/-고자
// 하다 and of quoting (-다고 한); a determiner takes none of them. A
// determiner is written apart from the word before it (clause 2 of the
// orthography), so that 한 written onto a word (스무드한, 창백(蒼白)한)
// is 하다 with ㄴ, unless that word is an opening sign (written onto the
// words it encloses) or ends in the modifying 는 (살아있는한: the noun
// 한(限) of -는 한, limitOneRule). Right before a unit 한 is the
// numeral determiner that counts it (countedUnitExcept).
func haOneRule() sentenceRule {
	g, claim := sentenceClaim(ruleHaOne, rule.Listed, listed("꾸밈 ㄴ", rule.ListSelfLast, []string{haAdn}))
	noTakes := except(g, rule.Listed, whyNoHaTakes, listed("하다가 받는 꼴", rule.ListPrevLast, haTakes), not)
	noTakes.Attacks(claim)
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(noTakes)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	except(g, rule.Listed, whyNeunBefore, listed("꾸밈 는", rule.ListPrevLast, []string{neunAdn})).Attacks(glued)
	countedUnitExcept(g, claim, numDeterminers, countedUnits)
	is := func(n rule.Near) bool {
		return n.Text == oneStem && n.Stem == haStem && n.Kind == haName && n.Last == haAdn
	}
	return sentenceRule{name: ruleHaOne, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
