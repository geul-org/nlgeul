//ff:func feature=parse type=rule control=sequence
//ff:what subjectRebuttalRule — the test rule of rebuttal that waits for the words after a candidate: a candidate whose last piece is the nominative 이 (주격:이) does not stand when the stem of the word right after it is 되 (its verdict is multiplied by 2/(1+forceUngrammatical)-1), and the exception "뒤 어절이 되가 아니다" defeats the rebuttal otherwise
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func subjectRebuttalRule() sentenceRule {
	isSubj := func(n rule.Near) bool { return n.Last == "주격:이" }
	g, claim := sentenceRebuttal("되 앞 주격", rule.Listed, listed("주격", rule.ListSelfLast, []string{"주격:이"}))
	except(g, rule.Listed, "뒤 어절이 되가 아니다", listed("되", rule.ListNextStem, []string{"되"}), not).Attacks(claim)
	return sentenceRule{name: "되 앞 주격", graph: g, force: forceUngrammatical, about: isSubj, waits: isSubj}
}
