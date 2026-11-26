//ff:func feature=parse type=rule control=sequence
//ff:what namedGate — the graph that asks the word before the suffix 이 of a called name (영숙이, 영숙이!) to be one that may be a name: the claim, the exception that it is a dependent noun of the closed set, and the exception that it stands as a word of the word table that is no nominal (an adverb, a determiner, a predicate, an interjection)
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The suffix is written after a person's name. A name is a nominal of
// the word table or a new word (no kind); a dependent noun names no one.
func namedGate() *entry {
	g := toulmin.NewGraph(nameGate)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: nameGate, Concludes: "열쇠"})
	except(g, rule.Listed, whyBound, listed("boundNouns", rule.ListForm, ClosedWords()[boundSet])).Attacks(claim)
	except(g, rule.StemKind, whyNoName, &rule.Bits{Of: "kinds", Mask: kindAny &^ kindNominal}).Attacks(claim)
	return &entry{name: nameGate, g: g}
}
