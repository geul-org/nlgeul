//ff:func feature=token type=reader control=selection
//ff:what headOf — what a slot or closing begins with, for the changes of an irregular stem (head…): a fused 어/아 or ㅆ; a fused ㄴ, ㄹ, ㅁ or ㅂ, and whether it is the form after a vowel of a tail that takes 으 (the closing tables have its form with 으: 은데 for ㄴ데, 읍시다 for ㅂ시다; none for ㄴ다 and ㄴ답니다, whose other form is 는다, nor for ㅂ니다); a syllable 어/아/여; a form with 으; else its first consonant, whether it takes 으 after a consonant, and 네
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// leadHeads are what a final consonant put under the stem's last
// syllable begins a tail with.
var leadHeads = map[int]int{rule.JongN: headN, rule.JongL: headL, rule.JongB: headB}

// A particle begins with nothing an irregular stem changes before, and
// a form read against the tables' attachment rule (loose) stands after
// the stem as it is written. The form with 으 of a fused consonant is
// the link vowel with that consonant under it, and the rest (은, 을, 음,
// 읍).
func (ix *index) headOf(e *entry) int {
	f := []rune(e.form)
	switch {
	case (e.class != classClosing && e.class != classSlot) || e.loose:
		return 0
	case e.lead == rule.LeadVowel || e.lead == rule.JongSS:
		return headFused
	case e.lead > 0 && ix.closingSet[string(linkVowel+rune(e.lead))+e.form]:
		return leadHeads[e.lead] | headEuOpen
	case e.lead > 0:
		return leadHeads[e.lead]
	case len(f) == 0:
		return 0
	case strings.ContainsRune("어아여었았였", f[0]):
		return headWhole
	case strings.ContainsRune(euLetters, f[0]):
		return headEuForm
	}
	h := firstHead(f[0])
	if e.eu || e.name == honorName || slices.Contains(linkTakers, e.name) {
		h |= headEuOpen
	}
	return h
}
