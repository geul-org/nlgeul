//ff:func feature=token type=helper control=sequence
//ff:what unspeltAt — is r a syllable after which an 어 left out in writing is read (되도 for 돼도; unspelt): for a fused 어/아 the syllable itself, when more of the tail follows or the word may end there, and for a fused ㅆ the syllable with ㅆ under it (됬다 for 됐다)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func unspeltAt(r rune, lead int, more bool) bool {
	if lead == rule.JongSS {
		base, ok := under(r, lead)
		return ok && slices.Contains(unspelt, base)
	}
	return lead == rule.LeadVowel && more && slices.Contains(unspelt, r)
}
