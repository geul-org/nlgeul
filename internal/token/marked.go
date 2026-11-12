//ff:func feature=token type=reader control=sequence
//ff:what marked — do the letters of the stem from b before tail t (which ends at j) read as markers of the index after a boundary inside the unit: a word with no marker after it, as one run of markers or as two runs one after the other (는 and then 게 in 한다 는게); a stem with markers after it, by themselves (에 before 서 in 마당 에+서, 이 before 다 in 것 이+다), or from the link vowel 으 on, with the tail (으 before 며) or after the 으 (으셨 before 다); the 으 may be what is left of 을 when the ㄹ under it is read as a piece (으+(ㄹ)걸); a tail that names its own stem is not asked, nor is a stem a piece is fused into unless it starts with the link vowel (으시 in 으셨다)
package token

// The fact is read here; whether such a new word stands is the new
// word's graph's to judge.
func (sc *scan) marked(b, j int, t tail) bool {
	if b == 0 || t.only != "" {
		return false
	}
	if len(t.pieces) == 0 {
		return sc.marks[j][b] || sc.twoRuns(b, j)
	}
	first := sc.s[b]
	if t.base != 0 && b == t.start-1 {
		first = t.base
	}
	link := first == linkVowel && (sc.marks[j][b] || sc.marks[j][b+1])
	return link || (t.base == 0 && sc.marks[t.start][b])
}
