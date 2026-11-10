//ff:func feature=token type=rule control=sequence
//ff:what markerStem — is the stem from b before tail t a new word, after a boundary inside the unit, whose letters read as markers of the index there (에 before 서 in 마당 에+서, 이 before 다 in 것 이+다): the new word's graph then does not let it stand; a word some table knows, a stem a piece is fused into, and a tail that names its own stem are not asked; only with a word table (without one nothing tells 들 the verb's stem from 들 the suffix)
package token

func (sc *scan) markerStem(b int, t tail) bool {
	if !sc.ix.lexed || b == 0 || t.base != 0 || t.shown != "" || t.as != "" || t.only != "" || !sc.marks[t.start][b] {
		return false
	}
	if _, _, v := sc.stemOf(b, t); v > 0 {
		return false
	}
	return sc.ix.newWord(true) <= 0
}
