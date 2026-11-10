//ff:func feature=token type=reader control=sequence
//ff:what entryKey — the index key of an entry: its form's last token, or for a piece with no token of its own the key of its lead
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func entryKey(e *entry) rune {
	if f := []rune(e.form); len(f) > 0 {
		return f[len(f)-1]
	}
	if e.lead == leadVowel {
		return keyVowel
	}
	return []rune(rule.LeadText(e.lead))[0]
}
