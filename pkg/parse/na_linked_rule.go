//ff:func feature=parse type=rule control=sequence
//ff:what naLinkedRule — the rule of a sentence "나다 after -고 is the auxiliary verb of a finished action": a candidate whose stem is 나 read as the verb (naStem, mainKindOf) with a closing is preferred, by bonusShortTaken, when the word right before it ends in the closing 고 (naLinks), and not otherwise, nor when it ends in the modifying 는 (naTopic: after -고 the pronoun 나 with the topic 는 begins the next clause), nor when it has a naming closing (naNamings: -고 남은 is 남다, 남과 the noun 남)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives 나다 as an auxiliary verb in -고 나다: the action
// before it is finished (듣고 나니, 먹고 나서, 끝나고 난 뒤). Its forms are
// also the letters of other words that the word table counts above it:
// the noun 나서, the pronoun 나 with the topic particle (난), the noun
// 나니. -고 also links two clauses (그는 가고 나는 남았다), and the first
// word of the second is often the pronoun 나 with the topic particle 는,
// so a candidate ending in the modifying 는 is left to the counts. The
// letters of 나다 with the naming -ㅁ are those of 남다 (쓰고 남은 돈:
// -고 남다) and of the noun 남, so a candidate with a naming closing is
// left out too. As with 말다, the value is what the rebuttal of one
// syllable takes (bonusShortTaken): the noun 나서 stands that far above.
func naLinkedRule() sentenceRule {
	g, claim := sentenceClaim(ruleNaLinked, rule.Listed, listed("나다", rule.ListSelfStem, []string{naStem}))
	except(g, rule.Listed, whyNotNaLinked, listed("고", rule.ListPrevLast, naLinks), not).Attacks(claim)
	except(g, rule.Listed, whyNaTopic, listed("꾸밈 는", rule.ListSelfLast, naTopic)).Attacks(claim)
	except(g, rule.SelfHas, whyNaNamed, listed("이름꼴", rule.ListSelfMarks, naNamings)).Attacks(claim)
	return sentenceRule{name: ruleNaLinked, graph: g, bonus: bonusShortTaken, about: func(n rule.Near) bool {
		return n.Stem == naStem && n.Kind == mainKindOf(naStem) && n.Pred
	}}
}
