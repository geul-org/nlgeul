//ff:func feature=token type=rule control=sequence
//ff:what shutTopicRule — the rule of a sentence "이란 and 란 written right after a closing sign are the topic particle": a candidate one of whose markers is the particle 이란 or 란 (topicNames) is preferred when the word right before it is a sign that closes a pair (shutName), and not otherwise; it adds bonusOverCount
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 이란 as the particle that picks out what stands
// before it as the topic ("소득"이란 …을 말한다). Letters written right
// onto a closing sign are what attaches to the words in the pair, a
// particle, not a noun written on with no space (the noun 이란).
func shutTopicRule() sentenceRule {
	g, claim := sentenceClaim(ruleShutTopic, rule.SelfHas, listed("보조사 이란·란", rule.ListSelfMarks, topicNames))
	except(g, rule.Listed, whyNotShut, listed("닫는 기호", rule.ListPrevKind, []string{shutName}), not).Attacks(claim)
	return sentenceRule{name: ruleShutTopic, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool { return hasMark(n, topicNames) }}
}
