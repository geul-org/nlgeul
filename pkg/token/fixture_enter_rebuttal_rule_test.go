//ff:func feature=token type=rule control=sequence
//ff:what enterRebuttalRule — a test rule of rebuttal that reads no other word: a candidate whose stem is 들어가 does not stand (its verdict is multiplied by 2/(1+forceUngrammatical)-1), enough to put the way that reads 들어가신다 as one word of 들어가 behind 들어 가신다
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func enterRebuttalRule() sentenceRule {
	g, _ := sentenceRebuttal("들어가 반박", rule.Listed, listed("들어가", rule.ListSelfStem, []string{"들어가"}))
	return sentenceRule{name: "들어가 반박", graph: g, force: forceUngrammatical, about: func(n rule.Near) bool { return n.Stem == "들어가" }}
}
