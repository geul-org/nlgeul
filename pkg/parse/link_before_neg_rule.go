//ff:func feature=parse type=rule control=sequence
//ff:what linkBeforeNegRule — the rule of a sentence "the -지 form right before 않다, 못하다 or 말다": a candidate that reads its stem as a predicate's (predStem) and one of whose markers is the linking ending 지 (negLinks: 지, and 진 for 지는) is preferred when the word right after it is 않다, 못하다 or 말다 (negStems), and not otherwise; it waits for the word after it
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives -지 as the ending a predicate takes before 않다,
// 못하다 and 말다, which deny what it says (가지 않는다, 사지 못하여,
// 가지 마라); a particle may stand between (가지도 않는다). The same
// letters read as a noun (가지, 사지) or as an adverb (내지) come first
// by the word table's counts where nothing is read after them. The -지
// form is a verb's, an adjective's or an auxiliary verb's: a new word
// with 지 after it (아버+지) is not asked, nor waited for.
func linkBeforeNegRule() sentenceRule {
	g, claim := sentenceClaim(ruleLinkBeforeNeg, rule.SelfHas, listed("맺음 지", rule.ListSelfMarks, negLinks))
	except(g, rule.Listed, whyNoNegNext, listed("않다·못하다·말다", rule.ListNextStem, negStems), not).Attacks(claim)
	is := func(n rule.Near) bool { return predStem(n) && hasMark(n, negLinks) }
	return sentenceRule{name: ruleLinkBeforeNeg, graph: g, bonus: bonusShortTaken, about: is, waits: is}
}
