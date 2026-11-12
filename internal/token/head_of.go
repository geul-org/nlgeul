//ff:func feature=token type=reader control=selection
//ff:what headOf — what a slot or closing begins with, for the changes of an irregular stem (head…): a fused 어/아 or ㅆ; a fused ㄴ, ㄹ, ㅁ or ㅂ (and whether it is the form after a vowel of a tail that takes 으); a syllable 어/아/여; a form with 으; else its first consonant, whether it takes 으 after a consonant, and 네
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A particle begins with nothing an irregular stem changes before, and
// a form read against the tables' attachment rule (loose) stands after
// the stem as it is written.
func headOf(e *entry) int {
	f := []rune(e.form)
	switch {
	case (e.class != classClosing && e.class != classSlot) || e.loose:
		return 0
	case e.lead == leadVowel || e.lead == jongSS:
		return headFused
	case e.lead == jongN && (e.wide || strings.HasPrefix(e.form, "다")):
		return headN
	case e.lead == jongN:
		return headN | headEuOpen
	case e.lead == rule.JongL:
		return headL | headEuOpen
	case e.lead == jongM:
		return headEuOpen
	case e.lead == rule.JongB && e.form == "시다":
		return headB | headEuOpen
	case e.lead == rule.JongB:
		return headB
	case len(f) == 0:
		return 0
	case strings.ContainsRune("어아여었았였", f[0]):
		return headWhole
	case strings.ContainsRune("으은을음읍", f[0]):
		return headEuForm
	}
	h := firstHead(f[0])
	if e.eu || e.name == honorName || slices.Contains(linkTakers, e.name) {
		h |= headEuOpen
	}
	return h
}
