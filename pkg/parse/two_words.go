//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what twoWords — are a stem and the suffix read after it two words of the word table written together: a nominal of two syllables or more, and then the rest of the stem with the suffix as a word (기동+전사 for 기동전+사); a nominal of one syllable before the word does not count (letters that long meet by chance)
package parse

func (ix *index) twoWords(stem, form string) bool {
	s := []rune(stem)
	for k := 2; k < len(s); k++ {
		_, tail := ix.lex[string(s[k:])+form]
		if tail && ix.lex[string(s[:k])].kinds&kindNominal != 0 {
			return true
		}
	}
	return false
}
