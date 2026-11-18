//ff:func feature=token type=rule control=sequence
//ff:what countedUnitExcept — put on a graph of a rule of a sentence the exception "the word right after it is a unit a numeral determiner counts": the stem of the word right after the candidate is one of the units given, and the exception itself falls when the candidate's letters are not those of a numeral determiner of the closed set (dets: 한, 두), so that it speaks only of a word whose letters are both a numeral determiner and something else (한 of 하다 with ㄴ), and falls as well when the word is written onto the word before it (a determiner is written apart), unless that word is an opening sign
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A numeral determiner tells how many of the noun after it there are,
// and that noun is mostly a unit (school grammar: 수 관형사; the
// dictionary's 한 「관형사」 한 사람, 책 한 권, 말 한 마리), so that 한
// right before a unit is the determiner rather than 하다 with ㄴ. A
// determiner is written apart from the word before it (clause 2 of the
// orthography), and an opening sign is written onto the words it
// encloses.
func countedUnitExcept(g *toulmin.Graph, claim *toulmin.Rule, dets, units []string) {
	unit := except(g, rule.Listed, whyUnitNext, listed("단위", rule.ListNextStem, units))
	unit.Attacks(claim)
	except(g, rule.Listed, whyNotNumDet, listed("수 관형사", rule.ListSelfText, dets), not).Attacks(unit)
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(unit)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
}
