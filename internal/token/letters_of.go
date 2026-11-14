//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lettersOf — the letters of a sentence with its spaces left out, which is what the places of a Span count
package token

import "unicode"

func lettersOf(sentence string) []rune {
	out := make([]rune, 0, len(sentence))
	for _, r := range sentence {
		if !unicode.IsSpace(r) {
			out = append(out, r)
		}
	}
	return out
}
