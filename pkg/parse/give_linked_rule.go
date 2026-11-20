//ff:func feature=parse type=rule control=sequence
//ff:what giveLinkedRule — the rule of a sentence "달다 after the 어/아 form is the auxiliary verb of asking": a candidate that reads 달 as a verb of the word table (giveStem; 달라, 달라고, 달라는) is preferred, by bonusOverCount, when the word before it ends in the linking closing 어/아/여, whole or fused (connNames), and not otherwise
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 달다 as an auxiliary verb in -어 달라 and -어
// 다오, asking the hearer to do what the verb before it says (도와 달라고
// 했다). Its form 달라 has the letters of 다르다 with 아 fused into it,
// which the word table counts more often. The word before is read as its
// first candidate reads it, whether a space stands between the two or
// not.
func giveLinkedRule() sentenceRule {
	g, claim := sentenceClaim(ruleGiveLinked, rule.Listed, listed("달다", rule.ListSelfStem, []string{giveStem}))
	except(g, rule.Listed, whyNotConn, listed("어/아 꼴", rule.ListPrevLast, connNames), not).Attacks(claim)
	return sentenceRule{name: ruleGiveLinked, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool {
		return n.Stem == giveStem && n.Kind == keyKind+":"+kindLabel[kindVerb]
	}}
}
