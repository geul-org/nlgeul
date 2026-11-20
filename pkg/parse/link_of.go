//ff:func feature=parse type=reader control=selection
//ff:what linkOf — the kind of a particle candidate for chains, from its case candidate, its form and its place in the particle table
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Candidates of one particle go their own ways: 에서 as a subject marker
// is a case particle, as a place an adverbial one. The subject marker 에서
// (after a word for a group: 정부에서) takes an auxiliary particle after it
// as the honorific subject marker 께서 does (정부에서는, 정부에서도), so it
// goes the way of 께서; the subject markers 이/가 do not (이는 is no
// chain). The shortened forms
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
	case form == "께서" || (form == "에서" && label == "주격"):
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
