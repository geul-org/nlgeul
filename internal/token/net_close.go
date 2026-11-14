//ff:func feature=token type=reader control=sequence
//ff:what netClose — close the net of a unit worked out up to its end, from the place a window cut it last (0 when it never did): the states that end the unit (the markers written apart from the unit before, attached, only when no cut was made: a cut takes away the words that cover the unit), every state kept, and the words found that are one print of other values (twins)
package token

func (sc *scan) netClose(best [][]*path, from int) *lattice {
	reg := sc.registry()
	ends := best[len(sc.s)]
	if from == 0 {
		ends = append(ends, sc.attached()...)
	}
	return &lattice{n: len(sc.s), ends: sc.settle(ends), all: reg.all, twins: reg.twins, from: from}
}
