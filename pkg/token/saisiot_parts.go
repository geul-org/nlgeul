//ff:func feature=token type=helper control=iteration dimension=1
//ff:what saisiotParts — the two nominals of the word table a stem is when 사이시옷 binds them into one (「한글 맞춤법」 제30항), at the first place from the front that fits: a syllable after the stem's first and before its last with ㅅ alone under it, the letters up to it (with the ㅅ) no word of the word table or of the closed sets (the ㅅ is that word's own: 오얏나무, 이것밖엔), the next syllable one 사이시옷 may be put before (saisiotOnset), the letters before the ㅅ (front, which ends in a vowel) and those after it (back) both nominals of the word table; ok is false when no place fits
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The front word has two syllables or more. 제30항 has fronts of one
// syllable too (귓밥, 콧물); they are left out by measurement, not by the
// article: a ㅅ under a stem's first syllable was more often a prefix
// (헛기침), the determiner 첫, a word's own final consonant or an old way
// to write ㅆ, which the letters do not tell apart.
func (ix *index) saisiotParts(stem string) (front, back string, ok bool) {
	s := []rune(stem)
	for i := 1; i < len(s)-1; i++ {
		own := string(s[:i+1])
		if rule.JongOf(s[i]) != jongS || !saisiotOnset(s[i+1]) || ix.lex[own].kinds != 0 || ix.closed[own] != 0 {
			continue
		}
		front, back = string(s[:i])+string(s[i]-jongS), string(s[i+1:])
		if ix.lex[front].kinds&kindNominal != 0 && ix.lex[back].kinds&kindNominal != 0 {
			return front, back, true
		}
	}
	return "", "", false
}
