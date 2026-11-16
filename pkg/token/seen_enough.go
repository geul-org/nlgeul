//ff:func feature=token type=rule control=sequence
//ff:what seenEnough — how many times a stem of the word table must have been seen not to be rebutted as rarely seen: the share seenShare of the sum of every line's count (lexTotal), so that it follows the size of the table, and at least one (with a table too small for the share to reach one count, no row is rarely seen)
package token

func (ix *index) seenEnough() int {
	return max(1, int(seenShare*ix.lexTotal))
}
