//ff:func feature=token type=helper control=sequence
//ff:what hiddenOf — a marker's text as the readers write it, parted in two: what the letters do not show (written in brackets at its head: (ㄴ)다, (이)) and the letters themselves
package token

import "strings"

func hiddenOf(text string) (string, string) {
	if rest, ok := strings.CutPrefix(text, "("); ok {
		hidden, shown, _ := strings.Cut(rest, ")")
		return hidden, shown
	}
	return "", text
}
