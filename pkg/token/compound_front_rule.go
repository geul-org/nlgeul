//ff:func feature=token type=rule control=sequence
//ff:what compoundFrontRule — the rule of a sentence "a word the word table gives both as a nominal and as an adverb, with a nominal written onto it, is the nominal": the candidate that reads such letters as the nominal of the word table (열쇠:체언) with no particle after it is preferred (bonusOverCount), but not when the word right after it is not written onto it, nor when that word is not one a modifier may stand before, nor when that word ends in a predicate's piece (a closing, a modifying or a naming form); it waits for the word after it
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Many words are given both as nouns and as adverbs (우선, 직접, 보통,
// 통상, 계속), and the word table counts most of them more often as the
// adverb. An adverb modifies a predicate, another adverb or the sentence,
// not a nominal (school grammar: 부사는 용언·다른 부사·문장을 꾸민다),
// and it is written apart from the word after it (clause 2 of the
// orthography). Such letters with a nominal written onto them with no
// space (우선출자, 직접사인, 보통면허) are the front of a compound noun,
// the nominal; adverbAloneRule leaves them by the same sign. A predicate
// written onto them (안하는, 가득찬) is the adverb's, written against the
// rule of spacing, and is left alone.
func compoundFrontRule(lex map[string]word) sentenceRule {
	g, claim := sentenceClaim(ruleCompoundFront, rule.Listed, listed("체언", rule.ListSelfKind, []string{tableNoun}))
	except(g, rule.AroundIs, whyApartNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeGlued}, not).Attacks(claim)
	except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not).Attacks(claim)
	except(g, rule.KindListed, whyPredNext, listed("서술어", rule.ListNextLast, predKinds)).Attacks(claim)
	is := func(n rule.Near) bool {
		return n.Kind == tableNoun && n.Marks == "" && lex[n.Stem].kinds&kindAdverb != 0
	}
	return sentenceRule{name: ruleCompoundFront, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
