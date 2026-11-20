//ff:func feature=parse type=rule control=sequence
//ff:what waitingComplementRule — the test rule that waits for the words after a candidate: a candidate whose last piece is the complement case 이 (보격:이) gains 2 when the stem of the word right after it is 되, and the exception "뒤 어절이 되가 아니다" defeats it otherwise
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func waitingComplementRule() sentenceRule {
	isComp := func(n rule.Near) bool { return n.Last == "보격:이" }
	return sentenceRuleOf("보격 뒤 되", 2, isComp, isComp,
		listed("보격", rule.ListSelfLast, []string{"보격:이"}),
		listed("되", rule.ListNextStem, []string{"되"}), "뒤 어절이 되가 아니다")
}
