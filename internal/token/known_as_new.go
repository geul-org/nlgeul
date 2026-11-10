//ff:func feature=token type=rule control=sequence
//ff:what knownAsNew — is the stem from b before tail t a word the tables know, but only as kinds that take no predicate tail (a nominal, an adverb, a determiner), standing as a new word before one (뭐+냐, 거+네, 코드+라): such a word is no new word's stem there; a word that is a predicate's stem too is left to its kind's graph; only with a word table
package token

func (sc *scan) knownAsNew(b int, t tail) bool {
	if !sc.ix.lexed || t.as != "" || tailClass(t.pieces)&predTails == 0 {
		return false
	}
	stem, _, v := sc.stemOf(b, t)
	kinds := sc.ix.lex[stem].kinds | sc.ix.closed[stem]
	return v <= 0 && kinds != 0 && kinds&(kindVerb|kindAdj) == 0
}
