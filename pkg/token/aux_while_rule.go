//ff:func feature=token type=rule control=sequence
//ff:what auxWhileRule — the rule of a sentence "보다 after -다(가) in 보니, 보면 is the auxiliary": the candidate whose stem is read as the auxiliary 보 (보조 용언:보) and that carries the closing 니, 니까 or 면 (whileClosings) is preferred when the word before it ends in 다 or 다가 (whileLinks), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 보다 as an auxiliary verb in -다(가) 보니 and
// -다(가) 보면: 살다 보니, 걷다가 보면.
func auxWhileRule() sentenceRule {
	name := auxSet + ":" + seeAux
	g, claim := sentenceClaim(ruleAuxWhile, rule.SelfHas, listed("보니·보면", rule.ListSelfMarks, whileClosings))
	except(g, rule.Listed, whyNotWhile, listed("다·다가", rule.ListPrevLast, whileLinks), not).Attacks(claim)
	return sentenceRule{name: ruleAuxWhile, graph: g, bonus: bonusSentence, about: func(n rule.Near) bool { return n.Kind == name }}
}
