//ff:func feature=token type=rule control=sequence
//ff:what auxLinkedRule — the rule of a sentence "an auxiliary verb stands after a linking form": a candidate whose stem is read as an auxiliary verb (보조 용언:있, of the closed set given) is preferred (by bonusOverCount) when the word before it ends in a linking closing an auxiliary follows (linkNames: 어/아/여, whole or fused, and 고 — but 고 only before 있 and 보: goAux), and not otherwise, nor after 관하여 (regardStems: the verb after it is a main one)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Clause 47 of the orthography: an auxiliary verb follows the 아/어 form
// (and 고) of its main verb. The dictionary gives -고 only before 있다
// (-고 있다) and 보다 (-고 보니, -고 보면); 주다, 놓다, 두다, 버리다 and
// 지다 follow -어 only (goAux). The word before is read as its first
// candidate reads it, whether a space stands between the two or not.
func auxLinkedRule(aux []string) sentenceRule {
	names := auxNamesOf(aux)
	g, claim := sentenceClaim(ruleAuxLinked, rule.Listed, listed("보조 용언", rule.ListSelfKind, names))
	except(g, rule.Listed, whyNotLinked, listed("연결 꼴", rule.ListPrevLast, linkNames), not).Attacks(claim)
	goOnly := except(g, rule.Listed, whyGoNotAux, listed("-고", rule.ListPrevLast, goNames))
	goOnly.Attacks(claim)
	except(g, rule.Listed, whyGoAux, listed("있다·보다", rule.ListSelfStem, goAux)).Attacks(goOnly)
	except(g, rule.Listed, whyRegard, listed("관하다", rule.ListPrevStem, regardStems)).Attacks(claim)
	return sentenceRule{name: ruleAuxLinked, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool { return slices.Contains(names, n.Kind) }}
}
