//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addFarKeys — the keys of the tables (the closed-set words, the word table, the changed forms of its stems and its wrong ways to write a word) that are longer than the stem of a new word may be: a stem of that length stands only as a known word (A new word's stem is up to newStemMax long; a known one has no limit)
package token

func (ix *index) addFarKeys() {
	ix.farKeys, ix.farBases, ix.keyPrefixes = map[string]bool{}, map[string]bool{}, map[string]bool{}
	for key := range ix.words {
		ix.addFarKey(key)
	}
	for key := range ix.closed {
		ix.addFarKey(key)
	}
	for key := range ix.lex {
		ix.addFarKey(key)
	}
	for key := range ix.alts {
		ix.addFarKey(key)
	}
	for key := range ix.typos {
		ix.addFarKey(key)
	}
}
