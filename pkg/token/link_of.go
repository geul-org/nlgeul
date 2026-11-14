//ff:func feature=token type=reader control=selection
//ff:what linkOf — the kind of a particle candidate for chains, from its case candidate, its form and its place in the particle table
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Candidates of one particle go their own ways: 에서 as a subject marker
// is a case particle, as a place an adverbial one. The shortened forms
// that end in 는 (엔, 에선, 까진 …) close a chain as a topic particle
// does.
func linkOf(label, form string, slot int) int {
	switch {
	case label == "접미":
		return rule.LinkSuffix
	case label == "호격":
		return rule.LinkVoc
	case label == "관형격":
		return rule.LinkAdn
	case form == "께서":
		return rule.LinkHonor
	case label == "주격" || label == "보격" || label == "목적격":
		return rule.LinkCase
	case form == "요":
		return rule.LinkYo
	case slices.Contains(shortTopics, form):
		return rule.LinkTopic
	case slices.Contains(directional, form):
		return rule.LinkDir
	case slices.Contains(comparing, form):
		return rule.LinkComp
	case slices.Contains(together, form):
		return rule.LinkWith
	case slot == 2 && form != "토록" && form != "따라":
		return rule.LinkRange
	case label == "부사격":
		return rule.LinkAdv
	}
	return rule.LinkTopic
}
