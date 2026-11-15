//ff:func feature=token type=rule control=sequence
//ff:what quotedRule — the rule of a sentence "라고 and 이라고 right after a closing quotation mark are the particle of direct quotation": a candidate one of whose markers is the particle 라고 or 이라고 (quoteNames) is preferred when the word right before it is a sign that closes a pair (shutName), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 라고 and 이라고 as the case particle that marks
// what stands before it as quoted word for word ("많이 드세요."라고).
func quotedRule() sentenceRule {
	g, claim := sentenceClaim(ruleQuoted, rule.SelfHas, listed("인용 조사 라고", rule.ListSelfMarks, quoteNames))
	except(g, rule.Listed, whyNotShut, listed("닫는 기호", rule.ListPrevKind, []string{shutName}), not).Attacks(claim)
	return sentenceRule{name: ruleQuoted, graph: g, bonus: bonusSentence, about: func(n rule.Near) bool { return hasMark(n, quoteNames) }}
}
