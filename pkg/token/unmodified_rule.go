//ff:func feature=token type=rule control=sequence
//ff:what unmodifiedRule — the rule of a sentence "a noun that ends in 인 or 일 is the noun where nothing an adnominal modifies follows it": a candidate that is a noun of the word table ending in the copula's modifying forms (about: adnNoun; 감정인, 시행일) is preferred when the word right after it is not one a word ending in a modifying closing may stand before (not a nominal, an order, a determiner, nor a modifying or naming form: an adverb, a conjunction, a predicate, a sign, or the end of the sentence); it waits for the word after it
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The same letters read as the noun before them with the copula and its
// modifying closing (감정+이+ㄴ, 시행+이+ㄹ) make an adnominal, and an
// adnominal modifies the nominal after it. Where no such word follows
// (감정인 및, 시행일., 방송인 "…"), the reading as the copula cannot stand
// and the noun of the table of the same letters is preferred
// (bonusUnmodified): the rebuttal "rarely seen" that a compound noun
// meets against the copula on a well seen noun does not hold it back.
func unmodifiedRule() sentenceRule {
	g := toulmin.NewGraph(ruleUnmodified)
	g.Rule(rule.AroundIs).With(&rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}).With(not).With(&rule.Claim{Name: ruleUnmodified, Concludes: preferLabel})
	return sentenceRule{name: ruleUnmodified, graph: g, bonus: bonusUnmodified, about: adnNoun, waits: adnNoun}
}
