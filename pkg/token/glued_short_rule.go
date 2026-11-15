//ff:func feature=token type=rule control=sequence
//ff:what gluedShortRule — the rule of a sentence "a noun of one syllable with its particle, written onto a nominal, is that noun and the particle": a candidate that is a noun of one syllable with a particle or the copula written right after it (about: shortNominal; 패소자의 as 자+의, not 자의) is preferred when no space stands before its word and the word right before it ends in a nominal (a noun, a number, an order or a unit after a number)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A noun of one syllable is rebutted as "one syllable long" because its
// letters are easily met by chance in a longer run (kindRebuttals); a
// word of two syllables that ends in the letters of a particle (자의, 상의,
// 기와) is not, and wins the same letters. Written onto a nominal, as the
// Sino-Korean nouns of one syllable that name a person, a standpoint or a
// period are (者 패소자, 上 실제상, 期 제1기), the front end of the noun is
// set by the nominal before it and its back end by the particle, so the
// rule gives back what that rebuttal took (bonusGluedShort) and the
// preferences of the readings decide. The claim is that the word is
// written onto the word before it.
func gluedShortRule() sentenceRule {
	g := toulmin.NewGraph(ruleGluedShort)
	claim := g.Rule(rule.AroundIs).With(&rule.Bits{Of: "around", Mask: rule.AroundGlued}).With(&rule.Claim{Name: ruleGluedShort, Concludes: preferLabel})
	except(g, rule.AroundIs, whyNotAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}, not).Attacks(claim)
	return sentenceRule{name: ruleGluedShort, graph: g, bonus: bonusGluedShort, about: shortNominal}
}
