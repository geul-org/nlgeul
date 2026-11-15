//ff:func feature=token type=rule control=sequence
//ff:what auxAskedRule — the rule of a sentence "보다 after a question form is the auxiliary": the candidate whose stem is read as the auxiliary 보 (보조 용언:보) is preferred when the word before it ends in one of the question closings 보다 follows as an auxiliary adjective (askNames: 나, ㄴ가/은가/는가, ㄹ까/을까), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 보다 as an auxiliary adjective in -은가/는가/나 보다
// (a guess) and -을까 보다 (an intention): 오나 보다, 갈까 보다.
func auxAskedRule() sentenceRule {
	name := auxSet + ":" + askAux
	g, claim := sentenceClaim(ruleAuxAsked, rule.Listed, listed("보조 용언 보다", rule.ListSelfKind, []string{name}))
	except(g, rule.Listed, whyNotAsked, listed("물음 꼴", rule.ListPrevLast, askNames), not).Attacks(claim)
	return sentenceRule{name: ruleAuxAsked, graph: g, bonus: bonusSentence, about: func(n rule.Near) bool { return n.Kind == name }}
}
