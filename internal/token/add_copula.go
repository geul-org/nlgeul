//ff:func feature=token type=rule control=sequence
//ff:what addCopula — the graphs of the copula stem: 이 said after a noun, and 이 dropped after a vowel; the claim "this is the stem of 이다 after a noun", and the exceptions that no stem stands before it, that the closing after it is one the copula does not take, that an auxiliary written onto a main predicate follows its 어/아 form (copulaBars), and for the dropped stem that the closing after it is not one the stem is dropped before (droppable, and with a word table droppableKnown), that no vowel stands before it, that the letter before it is no Hangul syllable (lifted before the closings read after such a letter: abc야), and that a naming closing follows (lifted when a particle is written right after that closing: 문제기도), with the rebuttal that a modifying or a naming closing follows
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The copula's count is not added to the total the other markers'
// shares are taken of, so that their preferences stay as they were.
func (ix *index) addCopula() {
	barred := listed("barred", rule.ListAfter, ix.barred)
	g := toulmin.NewGraph(copulaName)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: copulaName, Concludes: "표지"})
	except(g, rule.PieceIs, whyNoStem, fact(rule.PieceStem), not).Attacks(claim)
	except(g, rule.Listed, whyBarred, barred).Attacks(claim)
	except(g, rule.Listed, whyHelper, listed("copulaBars", rule.ListAfter, copulaBars)).Attacks(claim)
	ix.copula = &entry{name: copulaName, form: "이", count: copulaCount, g: g}

	drops := droppable
	if ix.lexed {
		drops = append(append([]string{}, droppable...), droppableKnown...)
	}
	d := toulmin.NewGraph(droppedName)
	claim = d.Rule(rule.PieceRead).With(&rule.Claim{Name: droppedName, Concludes: "표지"})
	except(d, rule.PieceIs, whyNoStem, fact(rule.PieceStem), not).Attacks(claim)
	except(d, rule.Listed, whyBarred, barred).Attacks(claim)
	except(d, rule.Listed, whyHelper, listed("copulaBars", rule.ListAfter, copulaBars)).Attacks(claim)
	except(d, rule.Listed, whyKept, listed("droppable", rule.ListAfter, drops), not).Attacks(claim)
	except(d, rule.PieceAfter, whyAfter, &rule.Bits{Of: "after", Mask: rule.AfterVowel}).Attacks(claim)
	d.Counter(rule.Listed).With(listed("dropAdn", rule.ListAfter, dropAdn)).Attacks(claim, forceDropAdn)
	run := except(d, rule.PieceAfter, whyNoHangul, &rule.Bits{Of: "after", Mask: rule.AfterHangul})
	run.Attacks(claim)
	except(d, rule.Listed, whyRunTakes, listed("droppableRun", rule.ListAfter, droppableRun)).Attacks(run)
	naming := except(d, rule.Listed, whyNaming, listed("dropNaming", rule.ListAfter, dropNaming))
	naming.Attacks(claim)
	except(d, rule.PieceRight, whyParticle, &rule.Bits{Of: "right", Mask: rule.RightParticle}).Attacks(naming)
	ix.dropped = &entry{name: droppedName, count: copulaCount, g: d}
	ix.graphs, ix.edges = ix.graphs+2, ix.edges+13
}
