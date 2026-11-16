//ff:func feature=token type=rule control=sequence
//ff:what degreeUnitRule — the rule of a sentence "도 right after a number at the end of its word is the unit first": a candidate whose stem is a number and whose last piece is the unit 도 (degreeNames) is preferred over the same letters read with the auxiliary particle 도, by bonusSameLetters
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The standard dictionary has 도 (度) as a dependent noun, the unit of
// an angle and of a temperature (섭씨 10도); written right after a
// number in figures it is that unit (한글 맞춤법 제43항). The particle
// 도 after a bare number (100분의 1도) loses to it; the rule has no
// exception. It reads nothing of the sentence, but the particle's
// preference (by its count) is about 0.25 above the unit's (one of the
// unit table's entries), which bonusSameLetters carries; 0.3 and 1 tried
// in its place on the answer sheet gave the same words.
func degreeUnitRule() sentenceRule {
	g, _ := sentenceClaim(ruleDegreeUnit, rule.Listed, listed("단위 도", rule.ListSelfLast, degreeNames))
	return sentenceRule{name: ruleDegreeUnit, graph: g, bonus: bonusSameLetters, about: func(n rule.Near) bool {
		return n.Kind == numName && slices.Contains(degreeNames, n.Last)
	}}
}
