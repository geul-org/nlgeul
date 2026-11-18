//ff:func feature=token type=rule control=sequence
//ff:what predAfterObjectRule — the rule of a sentence "a predicate stands right after an object": a candidate that reads its stem as a verb or an adjective of the word table (predStemKinds) and has a closing (Near.Pred) is preferred, by bonusOverCount, when the word right before it ends in the object case particle (objectKinds), and not otherwise, nor when its last piece is a modifying closing (argumentAdnRule judges those after an argument); a word whose letters are a negative adverb of the closed set (negAdverbs: 안, 못, 아니) is not asked
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// An object (목적어: a nominal with 을/를) is what the predicate after
// it acts on, and Korean puts the predicate after its arguments (school
// grammar: 목적어는 서술어 앞에 온다), so the word right after an object
// is most often its predicate (책을 읽었다, 살림을 보니). Read alone, the
// same letters as a noun (보니, 모아) or a noun with a particle (대서+는)
// come first by the word table's counts. A modifying form after an
// argument (뜰, 품은, 한) is the business of argumentAdnRule (and 한 of
// haOneRule), which also look at the word after it: this rule does not
// add to it again.
// The negative adverbs 안, 못 and 아니 (of which 안 is the short form)
// stand between an object and the predicate they deny (밥을 안 먹었다,
// 밥을 아니 먹었다), so a word of their letters is not asked: 아니 is
// also 알다 with -니. An auxiliary verb stands after a linking form, not after an
// object, and is not asked either.
func predAfterObjectRule() sentenceRule {
	g, claim := sentenceClaim(rulePredAfterObject, rule.Listed, listed("용언", rule.ListSelfKind, predStemKinds))
	except(g, rule.KindListed, whyNoObjectBefore, listed("목적격", rule.ListPrevLast, objectKinds), not).Attacks(claim)
	except(g, rule.KindListed, whyAdnSelf, listed("꾸밈 맺음", rule.ListSelfLast, []string{closingLabel(rule.KindAdn)})).Attacks(claim)
	return sentenceRule{name: rulePredAfterObject, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool {
		return n.Pred && slices.Contains(predStemKinds, n.Kind) && !slices.Contains(negAdverbs, n.Text)
	}}
}
