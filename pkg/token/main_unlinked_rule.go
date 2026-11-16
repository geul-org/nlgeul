//ff:func feature=token type=rule control=sequence
//ff:what mainUnlinkedRule — the rule of a sentence "a stem an auxiliary verb shares is the main predicate where no linking form stands before it": a candidate whose stem is one of the auxiliary verbs' (the closed set given) read under the word table's kind of its main predicate (mainKindOf: 있 the adjective, the others the verb) is preferred, but not when the word before it ends in a linking closing an auxiliary follows (linkNames) or in a question closing 보다 follows as an auxiliary (askNames); after 관하여 (regardStems) the linking form does not count; nor when the word before ends in 다 or 다가 (whileLinks) and the candidate is 보 with 니, 니까 or 면 (whileClosings: -다 보니, -다 보면)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func mainUnlinkedRule(aux []string) sentenceRule {
	g, claim := sentenceClaim(ruleMainUnlinked, rule.Listed, listed("보조 용언의 줄기", rule.ListSelfStem, aux))
	linked := except(g, rule.Listed, whyLinked, listed("연결 꼴", rule.ListPrevLast, linkNames))
	linked.Attacks(claim)
	except(g, rule.Listed, whyRegard, listed("관하다", rule.ListPrevStem, regardStems)).Attacks(linked)
	except(g, rule.Listed, whyAsked, listed("물음 꼴", rule.ListPrevLast, askNames)).Attacks(claim)
	while := except(g, rule.Listed, whyWhile, listed("다·다가", rule.ListPrevLast, whileLinks))
	while.Attacks(claim)
	except(g, rule.SelfHas, whyNoWhileClosing, listed("보니·보면", rule.ListSelfMarks, whileClosings), not).Attacks(while)
	except(g, rule.Listed, whyNotSee, listed("보다", rule.ListSelfStem, []string{seeAux}), not).Attacks(while)
	return sentenceRule{name: ruleMainUnlinked, graph: g, bonus: bonusSameLetters, about: func(n rule.Near) bool {
		return slices.Contains(aux, n.Stem) && n.Kind == mainKindOf(n.Stem)
	}}
}
