//ff:func feature=token type=rule control=sequence
//ff:what boundAfterRule — the rule of a sentence "a dependent noun stands right after a modifying form": a candidate that reads its stem as a dependent noun (의존명사:수, 의존명사:거) is preferred, by bonusBoundAfter, when the word right before it ends in a modifying closing (할 수도 → 수+도, 하는 건 → 거+는), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The other side of boundBeforeRule: after a modifying form the letters
// that a dependent noun begins are read as that noun with its particle
// first, before a longer noun (수도) or a verb (걸다+ㄴ) of the same
// letters.
func boundAfterRule() sentenceRule {
	g, claim := sentenceClaim(ruleBoundAfter, rule.KindListed, listed("의존 명사", rule.ListSelfKind, []string{boundSet}))
	except(g, rule.KindListed, whyNoAdn, listed("꾸밈 맺음", rule.ListPrevLast, []string{closingLabel(rule.KindAdn)}), not).Attacks(claim)
	return sentenceRule{name: ruleBoundAfter, graph: g, bonus: bonusBoundAfter, about: func(n rule.Near) bool { return nameKind(n.Kind) == boundSet }}
}
