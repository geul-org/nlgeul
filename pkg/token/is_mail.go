//ff:func feature=token type=helper control=iteration dimension=1
//ff:what isMail — is a text a mail address: one @ with a name's letters on both sides and a dot after it, and no other sign
package token

import "strings"

func isMail(s []rune) bool {
	at := strings.IndexRune(string(s), '@')
	text := string(s)
	if at < 1 || strings.Count(text, "@") != 1 || !strings.Contains(text[at:], ".") || strings.HasSuffix(text, ".") || strings.HasSuffix(text, "@") {
		return false
	}
	for _, r := range s {
		if r != '@' && !nameRune(r) {
			return false
		}
	}
	return true
}
