//ff:func feature=token type=rule control=sequence
//ff:what joinedGateOf — the graph the word before a shortened particle must pass: the claim, the exception that the word is no stem of the word table of the kinds asked (only when kinds are asked), and the exception that the stem and the piece right on it together are a word of the word table (the whole word stands first)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A shortened particle (서 for 에서, 선 for 에서는, 론 for 로는) is one
// syllable that ends many words as well: 요구서, 표선, 나이론. The
// exception keeps the particle off only the reading it makes itself; the
// whole word is read as it was, and nothing that stood before is taken
// away.
func joinedGateOf(name string, kinds int) *entry {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
	if kinds != 0 {
		except(g, rule.StemKind, whyKind, &rule.Bits{Of: "kinds", Mask: kinds}, not).Attacks(claim)
	}
	except(g, rule.PieceIs, whyJoinedParticle, fact(rule.PieceWhole)).Attacks(claim)
	return &entry{name: name, g: g}
}
