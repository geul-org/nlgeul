//ff:func feature=token type=rule control=sequence
//ff:what openMarkerRule — the rule of rebuttal of a sentence "a particle or a closing is written onto the word before it, and no word stands right after an opening sign": a candidate that reads its word as markers alone with no stem whose first is a particle, the copula or a closing (attachedMarks, the kinds of apartEnds) does not stand (forceUngrammatical) when the word right before it is an opening sign (열기: a quotation mark, a bracket); it reads only the word before it
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Clause 41 of the orthography writes a particle onto the word before
// it, and a closing ends the predicate's stem before it: neither stands
// with nothing before it. An opening sign opens what is quoted or put in
// brackets and is written onto the word after it (the spelling rules'
// appendix of punctuation), so a word right after it begins what it
// opens and no particle or closing is written onto it ('이 사건'의 이 is
// the determiner, not the particle 이). The same letters read otherwise
// (a determiner, a noun, a new word) come first.
func openMarkerRule() sentenceRule {
	g, _ := sentenceRebuttal(ruleOpenMarker, rule.KindListed, listed("여는 기호", rule.ListPrevName, []string{openName}))
	return sentenceRule{name: ruleOpenMarker, graph: g, force: forceUngrammatical, about: attachedMarks}
}
