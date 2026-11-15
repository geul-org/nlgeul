//ff:func feature=token type=rule control=sequence
//ff:what aloneRule — the rule of a sentence "a word alone in a pair of brackets is no bound noun and no determiner": a candidate that reads the letters of a closed-set bound noun or determiner (aloneWords) as a word of another part of speech (aloneFree) is preferred when the word right before it opens a pair (openName) and the word right after it closes one (shutName), so that the bound noun and the determiner of the same letters come after it
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A rule never lowers a candidate, so the other reading of the letters
// is lifted above the bound noun (詩[시]: 시[열쇠:체언] before
// 시[의존명사:시]) and the determiner ((이): 이[열쇠:체언] before
// 이[관형사:이]). The candidate waits for the word after it, the closing
// sign.
func aloneRule() sentenceRule {
	words := aloneWords()
	g, claim := sentenceClaim(ruleAlone, rule.Listed, listed("여는 기호", rule.ListPrevKind, []string{openName}))
	except(g, rule.Listed, whyNotShutAfter, listed("닫는 기호", rule.ListNextKind, []string{shutName}), not).Attacks(claim)
	about := func(n rule.Near) bool { return slices.Contains(words, n.Stem) && aloneFree(n) }
	return sentenceRule{name: ruleAlone, graph: g, bonus: bonusSentence, about: about, waits: about}
}
