//ff:func feature=token type=rule control=sequence
//ff:what adverbAloneRule — the rule of a sentence "a word the word table gives both as a nominal and as an adverb is the adverb where it stands alone": the candidate that reads such letters as the adverb of the word table (열쇠:부사) with no particle after it is preferred (bonusOverCount), but not right after a modifier (a word ending in a modifying closing or in the adnominal case particle, or a determiner of the closed set: a modifier is followed by the nominal it modifies), nor when it is written onto the word before it (unless what stands right before it is an opening sign), nor when a word a modifier may stand before is written onto it (the head of a compound written as one word: 우선출자; a sign or a predicate written onto it does not count)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Many words are given both as nouns and as adverbs (모두, 직접, 서로,
// 계속). A nominal takes particles (모두가, 직접의) and the word table
// counts those uses with the nominal, so it counts the nominal more
// often than the adverb even where the word stands alone. An adverb
// takes no case particle and modifies a predicate, another adverb or
// the sentence (school grammar: 부사는 용언·다른 부사·문장을 꾸민다); a
// nominal standing with no particle mostly stands after a modifier that
// modifies it (관형어 뒤의 체언). An adverb is written apart from the
// word before it (clause 2 of the orthography); an opening quotation
// mark or bracket is written onto the words it encloses. Letters
// with a nominal written onto them with no space are the head of a
// compound noun (우선출자, 보통면허), not an adverb written apart. The
// table's counts put the nominal above the adverb by up to about 0.7
// (잘못, 단지), so the adverb is carried over by bonusOverCount.
func adverbAloneRule(lex map[string]word) sentenceRule {
	adverb := keyKind + ":" + kindLabel[kindAdverb]
	g, claim := sentenceClaim(ruleAdverbAlone, rule.Listed, listed("부사", rule.ListSelfKind, []string{adverb}))
	except(g, rule.KindListed, whyModifierBefore, listed("관형어", rule.ListPrevLast, []string{closingLabel(rule.KindAdn), adnCase})).Attacks(claim)
	except(g, rule.Listed, whyDet, listed("관형사", rule.ListPrevKind, []string{determinerLabel})).Attacks(claim)
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	onto := except(g, rule.AroundIs, whyGluedNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeGlued})
	onto.Attacks(claim)
	except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not).Attacks(onto)
	is := func(n rule.Near) bool {
		return n.Kind == adverb && n.Marks == "" && lex[n.Stem].kinds&kindNominal != 0
	}
	return sentenceRule{name: ruleAdverbAlone, graph: g, bonus: bonusOverCount, about: is}
}
