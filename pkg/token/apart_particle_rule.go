//ff:func feature=token type=rule control=sequence
//ff:what apartParticleRule — the rule of a sentence "a particle written apart is the particle of the word before": a candidate that reads a whole word as markers with no stem whose first is a particle (the letters of a demonstrative of the closed set, 이, are left to determinerAloneRule) is preferred, by bonusNewTaken, but not when the word before it ends in a particle or a closing (apartEnds), nor when it is written onto the word before, nor when the word before does not end in a nominal — unless the word stands at the head of the sentence or right after a closing sign, a run of Latin letters and digits, a hyphened name or a code (apartBefore), or right after a dash (apartDash)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Clause 41 of the orthography writes a particle onto the word before
// it; one written apart (나무 는, ABC 가) still belongs to that word. A
// word of markers alone is read with a new word's verdict (loneOf),
// which holds it far behind a verb or a noun of the same letters (늘+ㄴ,
// 가+아, 라고); where the word before is one a particle attaches to (a
// nominal with no particle or closing after it, a closing sign, a dash,
// a run of Latin letters, a code), the rule gives back what that verdict
// takes. At the head of a sentence nothing stands before it in the
// sentence and the word it belongs to was cut off before it.
func apartParticleRule() sentenceRule {
	g, claim := sentenceClaim(ruleApartParticle, rule.Listed, listed("앞말 없는 표지", rule.ListSelfStem, []string{""}))
	except(g, rule.KindListed, whyEndsMarked, listed("조사·맺음", rule.ListPrevLast, apartEnds)).Attacks(claim)
	except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued}).Attacks(claim)
	bare := except(g, rule.AroundIs, whyNotAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}, not)
	bare.Attacks(claim)
	except(g, rule.Headed, whyHeadOrSign, listed("기호", rule.ListPrevKind, apartBefore)).Attacks(bare)
	except(g, rule.Listed, whyDashBefore, listed("줄표", rule.ListPrevLast, apartDash)).Attacks(bare)
	return sentenceRule{name: ruleApartParticle, graph: g, bonus: bonusNewTaken, about: func(n rule.Near) bool {
		marks := rule.MarksOf(n)
		return n.Stem == "" && n.Kind == "" && len(marks) > 0 && slices.Contains(particleKinds, labelOf(marks[0])) && !slices.Contains(pronouns, n.Text)
	}}
}
