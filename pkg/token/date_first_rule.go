//ff:func feature=token type=rule control=sequence
//ff:what dateFirstRule — the rule of a sentence "numbers joined by full stops in the shape of a date are a date first": a candidate whose stem is named 날짜 (dateNames) and is written with full stops is preferred over the same letters read as a version, by bonusOrder (an order among names of the same letters, not evidence of the sentence)
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The orthography's appendix on punctuation gives the full stop
// between numbers written in Arabic figures alone as the way to write a
// year, a month and a day (마침표 (3): 1919. 3. 1.); the candidate of a
// date only stands for a year of four digits, a month of 1 to 12 and a
// day of 1 to 31, where a version stands for any numbers. The rule has
// no exception.
func dateFirstRule() sentenceRule {
	g, _ := sentenceClaim(ruleDateFirst, rule.Listed, listed("날짜", rule.ListSelfKind, dateNames))
	return sentenceRule{name: ruleDateFirst, graph: g, bonus: bonusOrder, about: func(n rule.Near) bool {
		return slices.Contains(dateNames, n.Kind) && strings.Contains(n.Stem, ".")
	}}
}
