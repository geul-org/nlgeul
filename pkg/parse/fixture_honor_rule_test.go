//ff:func feature=parse type=rule control=sequence
//ff:what honorRule — a test rule that reads no other word: the candidate 들어가신+다 (a new stem 들어가신 with the closing 다) gains 30, enough to put it before 들어가+시+(ㄴ)다 in its node and to raise the lambda of the way it is in
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func honorRule() sentenceRule {
	return sentenceRuleOf("들어가신", 30, func(n rule.Near) bool { return n.Stem == "들어가신" }, nil,
		listed("다", rule.ListSelfLast, []string{"맺음:다"}), nil, "")
}
