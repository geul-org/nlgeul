//ff:func feature=parse type=rule control=sequence
//ff:what numberQuoteRule — the rule of a sentence "라고 and 이라고 after a number are the particle of quotation": a candidate one of whose markers is the particle 라고 or 이라고 (quoteNames) is preferred when its stem is read as a number (numberKinds), over 라고 read as an ending with the copula dropped after the digits
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A number or a symbol written in digits (②, 85%) is cited as it is
// written, as a name is (이하 ②라고 한다), and after a noun of Hangul
// 이라고 reads as the particle before the copula with -라고. After the
// digits the copula's dropped 이 costs nothing (a closing is read right
// on them), so with no rule the ending came first there alone.
func numberQuoteRule() sentenceRule {
	g, _ := sentenceClaim(ruleNumberQuote, rule.SelfHas, listed("인용 조사 라고", rule.ListSelfMarks, quoteNames))
	return sentenceRule{name: ruleNumberQuote, graph: g, bonus: bonusSameLetters, about: func(n rule.Near) bool {
		return slices.Contains(numberKinds, n.Kind) && hasMark(n, quoteNames)
	}}
}
