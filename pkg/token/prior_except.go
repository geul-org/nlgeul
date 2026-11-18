//ff:func feature=token type=rule control=sequence
//ff:what priorExcept — put on a graph of a rule of a sentence the exceptions "전 stands where priorNounRule reads it as the noun": the word right before it ends in a naming form or a quantity (priorKinds: -기, -음, a unit or an approximation after a number), or is a noun of an action standing bare (actions; that one falls when a particle or a closing is on the word before), and both fall when the candidate's stem is not 전 (priorStem), so that they speak only of the letters 전, and when the word right after it is written onto it (전국민: then 전 is the front of a compound, not the noun 전 that ends the phrase)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The standard dictionary gives 전(前) as a noun after a -기 form (밥을
// 먹기 전) and after a noun of time or a quantity (3년 전): the time
// before what is named; a noun of an action standing bare is read so as
// well (발병 전: 발병하기 전). 전(全) the determiner modifies the noun
// after it and is not bound to what stands before it. Where priorNounRule
// carries the noun by bonusOverCount, a rule that carries the determiner
// of the same letters by the same step must not stand, or the two cancel
// out and the word table's counts put the determiner first again.
// The noun 전 ends the phrase it heads and takes the particles itself
// (먹기 전에, 3년 전부터); a noun written onto it with no space (clause 2
// of the orthography: words are written apart) makes it the front of a
// compound (전국민), where the determiner's rule is left as it is.
func priorExcept(g *toulmin.Graph, claim *toulmin.Rule, actions []string) {
	notPrior := except(g, rule.Listed, whyNotPriorStem, listed("전", rule.ListSelfStem, []string{priorStem}), not)
	prior := except(g, rule.KindListed, whyPriorBefore, listed("이름꼴·수량", rule.ListPrevLast, priorKinds))
	prior.Attacks(claim)
	notPrior.Attacks(prior)
	action := except(g, rule.Listed, whyActionBefore, listed("동작성 명사", rule.ListPrevStem, actions))
	action.Attacks(claim)
	notPrior.Attacks(action)
	except(g, rule.Listed, whyMarkedBefore, listed("맺음 없음", rule.ListPrevLast, bareLast), not).Attacks(action)
	onto := except(g, rule.AroundIs, whyGluedNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeGlued})
	onto.Attacks(prior)
	onto.Attacks(action)
}
