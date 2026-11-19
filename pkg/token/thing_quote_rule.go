//ff:func feature=token type=rule control=sequence
//ff:what thingQuoteRule — the rule of rebuttal of a sentence "이라고 after 것 is not the particle of direct quotation": a candidate that reads the dependent noun 것 (thingStem) with 라고 or 이라고 as the case particle of direct quotation (quoteNames: 것+이라고, 것+이라고+는) does not stand (forceUngrammatical), always (it reads no other word)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The standard dictionary gives 이라고 as the case particle of direct
// quotation, written right after the closing quotation mark of what is
// quoted ("많이 드세요."라고), and -라고 after the copula 이 as the ending of
// indirect quotation (그것이 옳은 것이라고 했다). 것 written onto 이라고
// has no quotation mark between them, so the particle is not there: the
// letters are 것 with the copula and -라고. Other nominals with 이라고 are
// not asked.
func thingQuoteRule() sentenceRule {
	g, _ := sentenceRebuttal(ruleThingQuote, rule.SelfHas, listed("인용 조사 라고", rule.ListSelfMarks, quoteNames))
	is := func(n rule.Near) bool { return n.Stem == thingStem && hasMark(n, quoteNames) }
	return sentenceRule{name: ruleThingQuote, graph: g, force: forceUngrammatical, about: is}
}
