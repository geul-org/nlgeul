//ff:func feature=parse type=rule control=sequence
//ff:what adverbialFirstRule — the rule of a sentence "에서 and 서 are the adverbial particle first": a candidate one of whose markers is the adverbial 에서 or 서 (adverbialFirst) is preferred over the same letters read as the nominative, by bonusSameLetters (an order among names of the same letters, which also carries the candidate over another cut just above it)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 에서 as a nominative particle only after a noun
// that names a body of people (정부에서 발표했다); everywhere else it is
// the adverbial one. The rule has no exception yet: which nouns name
// such a body is in no table.
func adverbialFirstRule() sentenceRule {
	g, _ := sentenceClaim(ruleAdverbialFirst, rule.SelfHas, listed("부사격 에서·서", rule.ListSelfMarks, adverbialFirst))
	return sentenceRule{name: ruleAdverbialFirst, graph: g, bonus: bonusSameLetters, about: func(n rule.Near) bool { return hasMark(n, adverbialFirst) }}
}
