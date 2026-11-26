//ff:func feature=parse type=rule control=sequence
//ff:what droppedAfterRunRule — the rule of a sentence "a closing after a run of Latin letters and digits stands on the copula's dropped stem": a candidate one of whose markers is the dropped stem (droppedNames) is preferred when its stem is a run of letters and digits (runName), over the same closing read right on the run
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// A run of Latin letters and digits is no predicate's stem (it is a
// name, a word of another script, a code), so a closing written right
// after it is the copula's with the stem left out (Herder라는, value면),
// as it is after a noun of Hangul. With no rule the closing read right
// on the run came first, for the dropped stem's piece costs the
// preference of a short piece.
func droppedAfterRunRule() sentenceRule {
	g, claim := sentenceClaim(ruleDropRun, rule.SelfHas, listed("서술격 이(빠짐)", rule.ListSelfMarks, droppedNames))
	except(g, rule.Listed, whyNotRun, listed("영숫자 열", rule.ListSelfKind, []string{runName}), not).Attacks(claim)
	return sentenceRule{name: ruleDropRun, graph: g, bonus: bonusShortTaken, about: func(n rule.Near) bool {
		return n.Kind == runName && hasMark(n, droppedNames)
	}}
}
