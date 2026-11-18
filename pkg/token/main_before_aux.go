//ff:func feature=token type=reader control=sequence
//ff:what mainBeforeAux — is the stem from b with tail t (which ends at end) a main verb's written with an auxiliary verb onto its 어/아 form (clause 47 of the orthography lets the two be written together): the stem begins the unit, the tail is a predicate tail (not the copula's: a nominal is no main verb), its last piece is the linking closing 어/아, whole or fused (connNames), and the letters from end begin with an auxiliary verb's stem as it stands or with a piece fused into its last syllable (auxNext)
package token

import "slices"

// The fact is read here; what it lifts is the kinds' graphs' to judge.
// It is auxAfterConn seen from the main verb: the space before the unit
// sets the stem's front end, its 어/아 and the auxiliary's stem its back
// end.
func (sc *scan) mainBeforeAux(b, end int, t tail) bool {
	if b != 0 || end >= len(sc.s) || len(t.pieces) == 0 || tailClass(t.pieces)&predTails == 0 {
		return false
	}
	return slices.Contains(connNames, t.pieces[len(t.pieces)-1].names) && sc.auxNext(end)
}
