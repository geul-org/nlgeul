//ff:func feature=token type=helper control=sequence
//ff:what wordStem — the letters of a word's first part as a line of streamLines prints it: what stands before its first name or its first +
package token

import "strings"

func wordStem(word string) string {
	stem, _, _ := strings.Cut(word, "[")
	stem, _, _ = strings.Cut(stem, "+")
	return stem
}
