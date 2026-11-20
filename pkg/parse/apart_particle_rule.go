//ff:func feature=parse type=rule control=sequence
//ff:what apartParticleRule — the rule of a sentence "a particle written apart is the particle of the word before": a candidate that reads a whole word as markers with no stem whose first is a particle (the letters of a demonstrative of the closed set, 이, are left to determinerAloneRule) is preferred, by bonusNewTaken, but not when the word before it ends in a particle or a closing (apartEnds), nor when it is written onto the word before (unless that word is a run of signs alone with no name, which stands for a name kept back: ○○+가), nor when the word before does not end in a nominal — unless the word stands at the head of the sentence or right after a closing sign, a run of Latin letters and digits, a hyphened name or a code (apartBefore), or right after a dash (apartDash); and not for the vocative particle at the head of a sentence, where nothing stands before it to call
package parse

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
// takes. A word of markers alone written onto the word before is no
// particle written apart, but for a run of signs with no name before it:
// in a text such a run stands for a name kept back (○○, □□□, a nominal)
// and is a word of its own, so the particle written onto it (○○가) is
// read as a word of markers alone. At the head of a sentence nothing
// stands before it in the sentence and the word it belongs to was cut
// off before it; but the vocative particle (아, 야, 이여) stands right
// after the one it calls (the school grammar's 호격 조사: 영희야, 철수
// 아): at the head of a sentence there is nothing before it to call, and
// a word 아 or 야 there is no particle (아, 그렇구나).
func apartParticleRule() sentenceRule {
	g, claim := sentenceClaim(ruleApartParticle, rule.Listed, listed("앞말 없는 표지", rule.ListSelfStem, []string{""}))
	except(g, rule.KindListed, whyEndsMarked, listed("조사·맺음", rule.ListPrevLast, apartEnds)).Attacks(claim)
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.AroundIs, whySignsBefore, &rule.Bits{Of: "around", Mask: rule.AroundAfterSigns}).Attacks(glued)
	bare := except(g, rule.AroundIs, whyNotAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}, not)
	bare.Attacks(claim)
	except(g, rule.Headed, whyHeadOrSign, listed("기호", rule.ListPrevKind, apartBefore)).Attacks(bare)
	except(g, rule.Listed, whyDashBefore, listed("줄표", rule.ListPrevLast, apartDash)).Attacks(bare)
	vocative := except(g, rule.KindListed, whyVocativeHead, listed("호격", rule.ListSelfLast, []string{caseLabel[rule.CaseVoc]}))
	vocative.Attacks(claim)
	except(g, rule.Headed, whyNotHead, listed("첫머리", rule.ListPrevKind, nil), not).Attacks(vocative)
	return sentenceRule{name: ruleApartParticle, graph: g, bonus: bonusNewTaken, about: func(n rule.Near) bool {
		marks := rule.MarksOf(n)
		return n.Stem == "" && n.Kind == "" && len(marks) > 0 && slices.Contains(particleKinds, labelOf(marks[0])) && !slices.Contains(pronouns, n.Text)
	}}
}
