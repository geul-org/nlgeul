//ff:func feature=token type=reader control=sequence
//ff:what netStart — begin the net of a unit: the scan is made ready (prepare) and builds the net, and the root is all that is kept at place 0; the places 1 to the unit's end are then worked out one after another (netStep), and the net is closed (netEnd)
package token

func (sc *scan) netStart() [][]*path {
	sc.netted = true
	sc.prepare()
	best := make([][]*path, len(sc.s)+1)
	best[0] = []*path{sc.registry().root}
	return best
}
