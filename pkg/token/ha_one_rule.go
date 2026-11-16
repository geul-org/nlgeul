//ff:func feature=token type=rule control=sequence
//ff:what haOneRule — the rule of a sentence "한 after a complement of 하다 is 하다 with the modifying ㄴ": the candidate that reads the word 한 as the verb 하 of the word table with the modifying closing ㄴ is preferred, by bonusOverCount, when the word before it ends in a form that makes it a complement of 하다 (haTakes: the object particle, 로/으로, -게, -도록, -어야, -려(고), -고자, -다고 …), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 한 as a determiner (한 사람, 한 번) and as 하다
// with the modifying ㄴ; the word table's counts put the determiner first.
// 하다 is a transitive verb and an auxiliary one: it takes an object (일을
// 한), "…으로" (전제로 한) and the clauses of -게/-도록/-어야/-려고/-고자
// 하다 and of quoting (-다고 한); a determiner takes none of them.
func haOneRule() sentenceRule {
	g, claim := sentenceClaim(ruleHaOne, rule.Listed, listed("꾸밈 ㄴ", rule.ListSelfLast, []string{haAdn}))
	except(g, rule.Listed, whyNoHaTakes, listed("하다가 받는 꼴", rule.ListPrevLast, haTakes), not).Attacks(claim)
	is := func(n rule.Near) bool {
		return n.Text == oneStem && n.Stem == haStem && n.Kind == haName && n.Last == haAdn
	}
	return sentenceRule{name: ruleHaOne, graph: g, bonus: bonusOverCount, about: is}
}
