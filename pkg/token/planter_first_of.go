//ff:func feature=token type=helper control=sequence
//ff:what firstOf — the first word of the whole unit that has the sign as its first part, as it is carried over to the letters after the sign (movedFirst: the sign taken off, no stem, its preference less the sign's own, its end counted from the letters after the sign); false for any other state and for the sign alone
package token

func (pl *planter) firstOf(q *path) (step, bool) {
	if !pl.first(q) {
		return step{}, false
	}
	return movedFirst(q.last, pl.head, pl.signPref)
}
