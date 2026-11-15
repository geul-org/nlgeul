//ff:func feature=token type=reader control=sequence
//ff:what auxAfterConn — is the stem from b with tail t an auxiliary verb's stem written right after the 어/아 form of its main verb (clause 47 of the orthography lets the two be written together): the stem does not begin the unit, it is one of the auxiliaries' stems (auxStems; the stem as it was before a piece was fused into it, 주 for 줘), the tail is a predicate tail, and some candidate has a word that ends in the linking 어/아 end at b (connEnds)
package token

import "slices"

// The fact is read here; what it lifts is the kinds' graphs' to judge.
func (sc *scan) auxAfterConn(b int, stem string, t tail) bool {
	return b > 0 && sc.connEnds[b] && tailClass(t.pieces)&predTails != 0 && slices.Contains(sc.ix.auxStems(), stem)
}
