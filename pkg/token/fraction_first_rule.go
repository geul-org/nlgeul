//ff:func feature=token type=rule control=sequence
//ff:what fractionFirstRule — the rule of a sentence "numbers joined by a slash are a fraction first": a candidate whose stem is named 나눔 (fractionNames, read only of a slash between numbers) is preferred over the same letters read as a date, by bonusOrder (an order among names of the same letters, not evidence of the sentence)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The orthography's appendix on punctuation gives the slash between
// numbers as the way to write a fraction (빗금 (3): 3/4분기, 3/20); it
// gives no date written so. The rule has no exception.
func fractionFirstRule() sentenceRule {
	g, _ := sentenceClaim(ruleFractionFirst, rule.Listed, listed("나눔", rule.ListSelfKind, fractionNames))
	return sentenceRule{name: ruleFractionFirst, graph: g, bonus: bonusOrder, about: func(n rule.Near) bool { return slices.Contains(fractionNames, n.Kind) }}
}
