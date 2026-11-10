//ff:func feature=token type=reader control=sequence
//ff:what at — the letter at a place of the unit, or 0 outside it
package token

func (sc *scan) at(i int) rune {
	return atOf(sc.s, i)
}
