//ff:func feature=token type=rule control=sequence
//ff:what mainUnlinkedRule — the rule of a sentence "a stem an auxiliary verb shares is the main predicate where no linking form stands before it": a candidate whose stem is one of the auxiliary verbs' (the closed set given) read under the word table's kind of its main predicate (mainKindOf: 있 the adjective, the others the verb) is preferred, but not when the word before it ends in a linking closing an auxiliary follows (linkNames) or in a question closing 보다 follows as an auxiliary (askNames)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func mainUnlinkedRule(aux []string) sentenceRule {
	g, claim := sentenceClaim(ruleMainUnlinked, rule.Listed, listed("보조 용언의 줄기", rule.ListSelfStem, aux))
	except(g, rule.Listed, whyLinked, listed("연결 꼴", rule.ListPrevLast, linkNames)).Attacks(claim)
	except(g, rule.Listed, whyAsked, listed("물음 꼴", rule.ListPrevLast, askNames)).Attacks(claim)
	return sentenceRule{name: ruleMainUnlinked, graph: g, bonus: bonusSentence, about: func(n rule.Near) bool {
		return slices.Contains(aux, n.Stem) && n.Kind == mainKindOf(n.Stem)
	}}
}
