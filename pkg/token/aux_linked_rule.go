//ff:func feature=token type=rule control=sequence
//ff:what auxLinkedRule — the rule of a sentence "an auxiliary verb stands after a linking form": a candidate whose stem is read as an auxiliary verb (보조 용언:있, of the closed set given) is preferred when the word before it ends in a linking closing an auxiliary follows (linkNames: 어/아/여, whole or fused, and 고), and not otherwise
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Clause 47 of the orthography: an auxiliary verb follows the 아/어 form
// (and 고) of its main verb. The word before is read as its first
// candidate reads it, whether a space stands between the two or not.
func auxLinkedRule(aux []string) sentenceRule {
	names := auxNamesOf(aux)
	g, claim := sentenceClaim(ruleAuxLinked, rule.Listed, listed("보조 용언", rule.ListSelfKind, names))
	except(g, rule.Listed, whyNotLinked, listed("연결 꼴", rule.ListPrevLast, linkNames), not).Attacks(claim)
	return sentenceRule{name: ruleAuxLinked, graph: g, bonus: bonusSentence, about: func(n rule.Near) bool { return slices.Contains(names, n.Kind) }}
}
