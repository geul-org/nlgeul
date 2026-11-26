//ff:func feature=parse type=rule control=sequence
//ff:what suffixGateOf — the graph the word before a suffix of nounSuffixes must pass: the claim, the exception that the word is no nominal of the word table, and the exception that the stem and the suffix together are a word of the word table, or two words of it written together (the words of the table stand first)
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// It is named as the nominal gate is. The second exception keeps the
// suffix off only the reading it makes itself: the whole word is read as
// it was (합격률), and nothing that stood before is taken away.
func suffixGateOf() *entry {
	g := toulmin.NewGraph(nominalGate)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: nominalGate, Concludes: "열쇠"})
	except(g, rule.StemKind, whyKind, &rule.Bits{Of: "kinds", Mask: kindNominal}, not).Attacks(claim)
	except(g, rule.PieceIs, whyJoinedWord, fact(rule.PieceWhole)).Attacks(claim)
	return &entry{name: nominalGate, g: g}
}
