//ff:func feature=token type=reader control=sequence
//ff:what keysAt — the index keys of one token: the token itself, and as a fused token the final consonant under it or the vowel 어/아 fuses into
package token

func keysAt(r rune) []rune {
	keys := []rune{r}
	if k, ok := jamoKeys[jongOf(r)]; ok {
		keys = append(keys, k)
	}
	if jongOf(r) == 0 {
		keys = append(keys, keyVowel)
	}
	return keys
}
