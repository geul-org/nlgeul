//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what wordOver — is some run of the unit's letters that begins before from and ends at end a word of the word table: a word that holds letters of the stem, the closing after it and the particle read right after that closing (살아나 over 살+아+나)
package parse

func (sc *scan) wordOver(from, end int) bool {
	for b := 0; b < from; b++ {
		if _, ok := sc.ix.lex[string(sc.s[b:end])]; ok {
			return true
		}
	}
	return false
}
