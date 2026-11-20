//ff:func feature=parse type=rule control=sequence
//ff:what tableDeterminerRule — the rule of a sentence "a word the word table gives both as a nominal and as a determiner is the determiner where it stands alone before a word it may modify": the candidate that reads such letters as the determiner of the word table (열쇠:관형사) with no particle after it is preferred (bonusOverCount, the step nounModifiedRule gives the noun of the same letters after a modifier), unless the letters are a closed-set word (a determiner of the closed set, which has a candidate of its own name, or a pronoun of the closed set, which is a nominal whatever the table says), but not when it is written onto the word before it (unless what stands right before it is an opening sign), nor when the word right after it is not one a modifier may stand before (not a nominal, an order, a determiner, nor a modifying or naming form), nor for 전 where priorNounRule stands (after a naming form, a quantity or a bare noun of an action: priorExcept); it waits for the word after it
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// Some words are given both as nouns and as determiners (총, 약, 전,
// 각종, 특정). A nominal takes particles and the word table counts those
// uses with the nominal, so it counts the nominal more often than the
// determiner even where the word stands alone. A determiner takes no
// particle and stands before the nominal it modifies (school grammar:
// 관형사는 체언 앞에서 그 체언을 꾸민다), so the one standing alone with a
// word it may modify after it is the determiner. It is written apart
// from the word before it (clause 2 of the orthography). A modifier may
// stand before the determiner and the nominal it modifies (높은 총
// 비용), where nounModifiedRule carries the noun of the same letters by
// bonusOverCount; the determiner is carried by the same step, so that
// after a modifier the table's counts decide between the two, and
// elsewhere the determiner comes first. 전 is not the determiner where
// priorNounRule carries the noun of its letters by the same step (after a
// naming form, a quantity or a bare noun of an action: priorExcept), for
// there 전 is 전(前) the noun and the two steps would cancel out.
func tableDeterminerRule(lex map[string]word, closed map[string]int) sentenceRule {
	det := keyKind + ":" + determinerLabel
	g, claim := sentenceClaim(ruleTableDeterminer, rule.Listed, listed("관형사", rule.ListSelfKind, []string{det}))
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not).Attacks(claim)
	priorExcept(g, claim, actionNouns(lex))
	is := func(n rule.Near) bool {
		return n.Kind == det && n.Marks == "" && lex[n.Stem].kinds&kindNominal != 0 && closed[n.Stem] == 0
	}
	return sentenceRule{name: ruleTableDeterminer, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
