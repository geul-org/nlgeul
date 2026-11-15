//ff:func feature=token type=reader control=iteration dimension=1
//ff:what modifierWritten — are the letters of the new word from b before tail t (which ends at j) a modifier and a known word it modifies written together: at some place inside them (the stem's or the place where the tail begins, not at j) the letters before read as a modifier and those after as the word it modifies (modifiesAt); a tail that names its own stem or stands for a number is not asked
package token

// The place may be where the tail begins: the new word 없는 before the
// closing 게 is the modifying form 없+는 and 게 (거+이) written together.
func (sc *scan) modifierWritten(b, j int, t tail) bool {
	if t.only != "" || t.as != "" {
		return false
	}
	for k := b + 1; k <= t.start && k < j; k++ {
		if sc.modifiesAt(b, k, j) {
			return true
		}
	}
	return false
}
