//ff:func feature=token type=rule control=sequence
//ff:what outsideNounRule — the rule of a sentence "밖에 with no negative after it is the noun 밖 with the adverbial 에": a candidate that reads 밖 as a noun of the word table with the adverbial particle 에 as its first marker is preferred, by bonusOverCount, over the particle 밖에, but not when the word right after it or the one after that is a word of negation (negWords: 없다, 않다, 못하다, 모르다, 아니다 and the negative adverbs 안, 못); it waits for the words after it
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives the particle 밖에 ("그것 말고는") as always
// followed by a word of negation (공부밖에 모른다, 하나밖에 없다); where
// none follows, the same letters are the noun 밖 with the adverbial
// particle 에 (집 밖에 나갔다). The word of negation is looked for in the
// two words after it (하나밖에 남지 않았다).
func outsideNounRule() sentenceRule {
	g, claim := sentenceClaim(ruleOutsideNoun, rule.Listed, listed("밖", rule.ListSelfStem, []string{outsideStem}))
	except(g, rule.Listed, whyNegNext, listed("부정을 나타내는 말", rule.ListNextStem, negWords)).Attacks(claim)
	except(g, rule.Listed, whyNegNext2, listed("부정을 나타내는 말", rule.ListNext2Stem, negWords)).Attacks(claim)
	is := func(n rule.Near) bool {
		marks := rule.MarksOf(n)
		return n.Stem == outsideStem && n.Kind == tableNoun && len(marks) > 0 && marks[0] == outsideCase
	}
	return sentenceRule{name: ruleOutsideNoun, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
