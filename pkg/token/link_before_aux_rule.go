//ff:func feature=token type=rule control=sequence
//ff:what linkBeforeAuxRule — the rule of a sentence "the word right before an auxiliary verb is the 어/아 form of its main verb": a candidate that ends in the linking closing 어/아/여, whole or fused (connNames), is preferred when the stem of the word right after it is one of the auxiliary verbs' (the closed set given), whether a space stands between the two or not, and not otherwise
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Clause 47 of the orthography: an auxiliary verb follows the 아/어 form
// of its main verb (남아 있다, 남아있다). auxLinkedRule reads the pair
// from the auxiliary's side; this reads it from the main verb's, where
// the letters of the 어/아 form are also a noun (남아, 서). The word
// after is read as its first candidate reads it, main or auxiliary, so
// only its stem is asked. The candidate waits for the words after it.
func linkBeforeAuxRule(aux []string) sentenceRule {
	g, claim := sentenceClaim(ruleLinkBeforeAux, rule.Listed, listed("어/아 꼴", rule.ListSelfLast, connNames))
	except(g, rule.Listed, whyNoAuxNext, listed("보조 용언의 줄기", rule.ListNextStem, aux), not).Attacks(claim)
	is := func(n rule.Near) bool { return slices.Contains(connNames, n.Last) }
	return sentenceRule{name: ruleLinkBeforeAux, graph: g, bonus: bonusLinkBeforeAux, about: is, waits: is}
}
