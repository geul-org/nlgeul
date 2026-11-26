//ff:func feature=parse type=rule control=sequence
//ff:what quotedAdnRule — the rule of a sentence "는 right after a closing quotation mark that closes a quoted sentence ending in a closing is the modifying -는 of the quoting verb": a candidate that has the modifying 는 (quotedAdnNames) is preferred over the same letters read with the auxiliary particle 는, by bonusSameLetters, when the word right before it is a closing mark whose own word before it ends in a closing (a sign that ends a sentence is skipped), and not otherwise
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives -다는, -냐는, -자는 and -라는 as endings of the
// quoting verb ('-다고 하는' shortened). A quotation mark that closes
// a quoted sentence only splits such an ending (…있다"는 취지로), so the
// 는 written right after it is that ending; after a quoted name or title
// (a nominal before the mark) it stays the particle.
func quotedAdnRule() sentenceRule {
	g, claim := sentenceClaim(ruleQuotedAdn, rule.SelfHas, listed("꾸밈 는", rule.ListSelfMarks, quotedAdnNames))
	except(g, rule.QuotedSentence, whyNotQuoteEnd, listed("문장 끝 부호", rule.ListPrevLast, stopMarks), not).Attacks(claim)
	return sentenceRule{name: ruleQuotedAdn, graph: g, bonus: bonusSameLetters, about: func(n rule.Near) bool { return hasMark(n, quotedAdnNames) }}
}
