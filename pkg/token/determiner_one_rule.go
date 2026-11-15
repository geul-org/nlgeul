//ff:func feature=token type=rule control=sequence
//ff:what determinerOneRule — the rule of a sentence "한 standing alone is the determiner": the candidate that reads 한 as the determiner of the closed set (관형사:한), written apart from the word before it, is preferred, but not when the word before it ends in a form that makes it a complement of 하다 (haTakes: the object particle, 로/으로, -게, -도록, -다고 …), where 한 is 하다 with the modifying ㄴ
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 한 as a determiner (한 사람, 한 번) and 한 as
// 하다 with ㄴ; the word table's counts put the second first. 하다
// takes an object (일을 한), "…으로" (전제로 한) and the clauses of
// -게/-도록/-어야/-려고/-고자 하다 and of quoting (-다고 한); a
// determiner takes none of them. A determiner is written apart from
// the word before it (clause 2 of the orthography): 한 written onto a
// word (무형한, 불쾌한) is not asked about.
func determinerOneRule() sentenceRule {
	g, claim := sentenceClaim(ruleDeterminerOne, rule.Listed, listed("관형사 한", rule.ListSelfKind, []string{oneDeterminer}))
	except(g, rule.Listed, whyHaTakes, listed("하다가 받는 꼴", rule.ListPrevLast, haTakes)).Attacks(claim)
	return sentenceRule{name: ruleDeterminerOne, graph: g, bonus: bonusOne, about: func(n rule.Near) bool { return n.Kind == oneDeterminer && !n.Glued }}
}
