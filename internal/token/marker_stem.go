//ff:func feature=token type=rule control=sequence
//ff:what markerStem — is the stem from b before tail t (which ends at j) a new word, after a boundary inside the unit, whose letters read as markers of the index there: by themselves (에 before 서 in 마당 에+서, 이 before 다 in 것 이+다), or from the link vowel 으 on, with the tail (으 before 며) or after the 으 (으셨 before 다): the new word's graph then does not let it stand; a word some table knows, a stem a piece is fused into, and a tail that names its own stem are not asked; only with a word table (without one nothing tells 들 the verb's stem from 들 the suffix)
package token

func (sc *scan) markerStem(b, j int, t tail) bool {
	link := sc.s[b] == linkVowel && (sc.marks[j][b] || sc.marks[j][b+1])
	if !sc.ix.lexed || b == 0 || t.base != 0 || t.shown != "" || t.as != "" || t.only != "" || !(sc.marks[t.start][b] || link) {
		return false
	}
	if _, _, v := sc.stemOf(b, t); v > 0 {
		return false
	}
	return sc.ix.newWord(true) <= 0
}
