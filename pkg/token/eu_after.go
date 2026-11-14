//ff:func feature=token type=helper control=sequence
//ff:what euAfter — what must stand before a whole-syllable closing: as the table says, but a form with the link vowel 으 in it (은, 을, 음, 읍시다 and their like), which the table lets attach after any consonant, attaches after a consonant other than ㄹ (a stem that ends in ㄹ takes ㄴ, ㄹ, ㅁ instead: 물 것, not 물을 것)
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func euAfter(form string, after int) int {
	if after == rule.AfterBatchim && strings.ContainsRune(euLetters, []rune(form)[0]) {
		return rule.AfterHard
	}
	return after
}
