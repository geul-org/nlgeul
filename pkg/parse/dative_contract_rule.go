//ff:func feature=parse type=rule control=sequence
//ff:what dativeContractRule — the rule of a sentence "내게, 네게 and 제게 are the pronoun with the particle 게": a candidate that reads its stem as one of the pronoun forms the contracts table puts the case particle 게 under (dativeStems: 내, 네, 제, 자네, 자기) with that particle (부사격:에게, dativeNames) is preferred, but not when the word right after it is 하다, 되다 or 만들다 (causeStems), before which 게 is the ending of the verb 내다 (-게 하다); it waits for the word after it
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The dictionary gives 내게, 네게 and 제게 as the pronouns 나, 너 and 저
// with 에게 fused (내 is the form 나 takes before 게). The same letters
// read as the verb 내다 with the ending -게 stand before the verb they
// complete (내게 하다: make one put out, 내게 되다, 내게 만들다); the
// word table counts the verb 내 and the ending 게 more often, so the
// pronoun fell behind wherever it stood.
func dativeContractRule() sentenceRule {
	g, claim := sentenceClaim(ruleDativeContract, rule.SelfHas, listed("부사격 에게", rule.ListSelfMarks, dativeNames))
	except(g, rule.Listed, whyCauseNext, listed("하다·되다·만들다", rule.ListNextStem, causeStems)).Attacks(claim)
	is := func(n rule.Near) bool { return slices.Contains(dativeStems, n.Stem) && hasMark(n, dativeNames) }
	return sentenceRule{name: ruleDativeContract, graph: g, bonus: bonusSameLetters, about: is, waits: is}
}
