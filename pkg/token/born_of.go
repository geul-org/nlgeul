//ff:func feature=token type=helper control=sequence
//ff:what bornOf — the number of birth of a candidate of one word: the one kept with it; noBorn for a branch made without scores or with no word
package token

func bornOf(br Branch) int {
	if len(br.Words) == 0 {
		return noBorn
	}
	return scoresOf(br)[0].born
}
