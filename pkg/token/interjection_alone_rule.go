//ff:func feature=token type=rule control=sequence
//ff:what interjectionAloneRule — the rule of a sentence "a word the word table gives both as a nominal and as an interjection is the interjection where it stands alone": the candidate that reads such letters as the word of the table that stands alone (열쇠:홀로: the interjections, and a few Latin words given as a conjunction or a preposition) with no particle after it is preferred (bonusOverCount), but not right after a modifier (a word ending in a modifying closing or in the adnominal case particle, or a determiner of the closed set), nor when it is written onto the word before it (unless what stands right before it is an opening sign), nor when a word a modifier may stand before or an opening sign is written onto it (a sign that closes and a predicate do not count)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Some words are given both as nouns and as interjections (허, 흠, 흥,
// 어). The word table counts the uses of the noun with its particles,
// so it counts the noun more often even where the word stands alone.
// An interjection takes no particle and stands apart from the other
// words of its sentence as an independent word (school grammar:
// 감탄사는 독립어로 쓰인다); a nominal with no particle mostly stands after
// a modifier that modifies it. The table's counts put the noun above the
// interjection by up to about 0.8 (흥), so it is carried over by
// bonusOverCount. The exceptions are those of adverbAloneRule, and an
// opening bracket written onto the word (아하(A-Ha)) opens a gloss of a
// name, so it holds the word back as the words written onto it do.
func interjectionAloneRule(lex map[string]word) sentenceRule {
	alone := keyKind + ":" + kindLabel[kindAlone]
	g, claim := sentenceClaim(ruleInterjectionAlone, rule.Listed, listed("감탄사", rule.ListSelfKind, []string{alone}))
	except(g, rule.KindListed, whyModifierBefore, listed("관형어", rule.ListPrevLast, []string{closingLabel(rule.KindAdn), adnCase})).Attacks(claim)
	except(g, rule.Listed, whyDet, listed("관형사", rule.ListPrevKind, []string{determinerLabel})).Attacks(claim)
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	onto := except(g, rule.AroundIs, whyGluedNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeGlued})
	onto.Attacks(claim)
	sign := except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not)
	sign.Attacks(onto)
	except(g, rule.Listed, whyOpenAfter, listed("여는 기호", rule.ListNextKind, []string{openName})).Attacks(sign)
	is := func(n rule.Near) bool {
		return n.Kind == alone && n.Marks == "" && lex[n.Stem].kinds&kindNominal != 0
	}
	return sentenceRule{name: ruleInterjectionAlone, graph: g, bonus: bonusOverCount, about: is}
}
