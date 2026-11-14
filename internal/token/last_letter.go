//ff:func feature=token type=helper control=sequence
//ff:what lastLetter — the last letter of the last part of a word (0 when it has none)
package token

func lastLetter(e Eojeol) rune {
	if len(e.Parts) == 0 {
		return 0
	}
	r := []rune(e.Parts[len(e.Parts)-1].Text)
	if len(r) == 0 {
		return 0
	}
	return r[len(r)-1]
}
