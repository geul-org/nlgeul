//ff:func feature=parse type=rule control=sequence
//ff:what gluedNominalRule — the rule of a sentence "a word the word table gives both as a nominal and as an adverb or a determiner, written onto a nominal before it, is the nominal": the candidate that reads such letters as the nominal of the word table (열쇠:체언) with no particle after it is preferred (bonusOverCount), unless the letters are a closed-set word, but not when it is not written onto the word before it, nor when that word does not end in a nominal, nor when a word ending in a predicate's piece (a closing, a modifying or a naming form) is written onto it; it waits for the word after it
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// A determiner stands before the nominal it modifies and an adverb
// before what it modifies (school grammar), and both are written apart
// from the word before them (clause 2 of the orthography). Letters the
// word table gives both as a nominal and as a determiner or an adverb
// (전, 조, 약, 계속), written onto a nominal with no space (발병전,
// 저장조, 상비약, 사건계속법원), are the back of a compound noun, the
// nominal; the table counts several of them more often as the
// determiner or the adverb. A predicate written onto them after (도입안하고)
// shows an adverb written against the rule of spacing, and is left
// alone. The letters of a closed-set word (이, 그, 네) are left to the
// rules of their own.
func gluedNominalRule(lex map[string]word, closed map[string]int) sentenceRule {
	g, claim := sentenceClaim(ruleGluedNominal, rule.Listed, listed("체언", rule.ListSelfKind, []string{tableNoun}))
	except(g, rule.AroundIs, whyApartBefore, &rule.Bits{Of: "around", Mask: rule.AroundGlued}, not).Attacks(claim)
	except(g, rule.AroundIs, whyNotAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}, not).Attacks(claim)
	pred := except(g, rule.KindListed, whyPredNext, listed("서술어", rule.ListNextLast, predKinds))
	pred.Attacks(claim)
	except(g, rule.AroundIs, whyApartNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeGlued}, not).Attacks(pred)
	is := func(n rule.Near) bool {
		return n.Kind == tableNoun && n.Marks == "" && lex[n.Stem].kinds&(kindAdverb|kindDeterminer) != 0 && closed[n.Stem] == 0
	}
	return sentenceRule{name: ruleGluedNominal, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
