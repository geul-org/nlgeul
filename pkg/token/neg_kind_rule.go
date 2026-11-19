//ff:func feature=token type=rule control=sequence
//ff:what negKindRule — the rule of rebuttal of a sentence "못하다 after the -지 form is of the kind of the predicate before it": a candidate that reads 못하 (negKindStems) as the word table's kind self does not stand (forceUngrammatical) when the word right before it carries the linking ending -지 (negLinks, a particle may follow: 가지도) and the name of its stem is one of prev, those of the other kind (why: the exception when it is not)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The standard dictionary gives 못하다 as an auxiliary verb after the
// -지 form of a verb (가지 못하다) and as an auxiliary adjective after
// that of an adjective (편하지 못하다): the auxiliary is of the kind of
// the predicate it follows. The word table counts the adjective far more
// (it counted both uses as one), so the adjective came first after a
// verb. The rule is put twice, one for each way: the adjective after a
// verb or an auxiliary verb of the closed set, the verb after an
// adjective. The closed set 보조 용언 (주 보 버리 놓 두 있 지) holds
// auxiliary verbs (-어 주다, -어 보다, -어 있다·-고 있다 …); 보다 is an
// auxiliary adjective only in -은가 보다 and -을까 보다, which do not take
// -지 못하다. The word before is read as its first candidate, as every
// rule of a sentence reads it. 않다 and 아니하다 are left out
// (negKindStems).
func negKindRule(self string, prev []string, why string) sentenceRule {
	g, claim := sentenceRebuttal(ruleNegKind, rule.PrevHas, listed("맺음 지", rule.ListPrevLast, negLinks))
	except(g, rule.Listed, why, listed("앞 줄기 갈래", rule.ListPrevName, prev), not).Attacks(claim)
	is := func(n rule.Near) bool { return n.Kind == self && slices.Contains(negKindStems, n.Stem) }
	return sentenceRule{name: ruleNegKind, graph: g, force: forceUngrammatical, about: is}
}
