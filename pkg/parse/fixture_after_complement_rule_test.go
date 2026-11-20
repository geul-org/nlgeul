//ff:func feature=parse type=rule control=sequence
//ff:what afterComplementRule — the test rule that reads only the words before a candidate (waits nil): a candidate whose stem is 되 gains 3 when the first candidate of the word right before it ends in the complement case 이 (보격:이)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func afterComplementRule() sentenceRule {
	return sentenceRuleOf("되 앞 보격", 3, func(n rule.Near) bool { return n.Stem == "되" }, nil,
		listed("앞 보격", rule.ListPrevLast, []string{"보격:이"}), nil, "")
}
