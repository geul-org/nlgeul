//ff:func feature=token type=rule control=sequence
//ff:what malLinkedRule — the rule of a sentence "말다 after -지 or -고 is the auxiliary verb": a candidate whose stem is 말 read as the verb (malStem, mainKindOf) is preferred when the word before it ends in the closing 지 or 고 (malLinks), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 말다 as an auxiliary verb in -지 말다 (가지 마라)
// and -고 말다 (지고 말았다). Its forms that drop the ㄹ (마, 마라,
// 마시고, 마소서) are also other words' letters (마, 마라, 마시다, 마소),
// which come first without the word before.
func malLinkedRule() sentenceRule {
	g, claim := sentenceClaim(ruleMalLinked, rule.Listed, listed("말다", rule.ListSelfStem, []string{malStem}))
	except(g, rule.Listed, whyNotMalLinked, listed("지·고", rule.ListPrevLast, malLinks), not).Attacks(claim)
	return sentenceRule{name: ruleMalLinked, graph: g, bonus: bonusShortTaken, about: func(n rule.Near) bool {
		return n.Stem == malStem && n.Kind == mainKindOf(malStem)
	}}
}
