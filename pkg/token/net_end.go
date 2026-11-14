//ff:func feature=token type=reader control=sequence
//ff:what netEnd — close the net of a unit worked out up to its end: the states that end the unit (with the markers written apart from the unit before, attached), every state made, and the words found that are one print of other values (twins)
package token

func (sc *scan) netEnd(best [][]*path) *lattice {
	return sc.netClose(best, 0)
}
