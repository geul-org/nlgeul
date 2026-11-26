//ff:func feature=parse type=rule control=sequence
//ff:what droppedAfterShutRule — the rule of a sentence "a closing right after a closing sign stands on the copula's dropped stem": a candidate one of whose markers is the dropped stem (droppedNames) is preferred when the word before it is a closing sign of a pair (shutName) written right onto it, unless one of its markers is a closing written right onto the sign with no copula (shutBareClosings: 고, 며, 라는), over the same closing read right on the sign
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// What a closing sign ends stands as a nominal (구본무(具本茂)다), so a
// closing the stem is dropped before is the copula's with the stem
// left out. The closings that quote or link (…)고, ’며, 」라는) are
// written onto the sign with no copula, as the answer sheet writes them.
func droppedAfterShutRule() sentenceRule {
	g, claim := sentenceClaim(ruleDropShut, rule.SelfHas, listed("서술격 이(빠짐)", rule.ListSelfMarks, droppedNames))
	except(g, rule.Listed, whyNotShut, listed("닫는 기호", rule.ListPrevKind, []string{shutName}), not).Attacks(claim)
	except(g, rule.AroundIs, whyApartBefore, &rule.Bits{Of: "around", Mask: rule.AroundGlued}, not).Attacks(claim)
	except(g, rule.SelfHas, whyShutBare, listed("서술격 없이 서는 맺음", rule.ListSelfMarks, shutBareClosings)).Attacks(claim)
	return sentenceRule{name: ruleDropShut, graph: g, bonus: bonusShortTaken, about: func(n rule.Near) bool {
		return hasMark(n, droppedNames)
	}}
}
