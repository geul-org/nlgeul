//ff:func feature=token type=reader control=iteration dimension=1
//ff:what attached — the readings of a whole unit as markers written apart from their word (로, 에서, 인데 after a space, and 게, 건, 걸 after a modifying form): each particle chain, and each tail of the copula, that covers the unit from its first letter, read as if a stem stood before it, is one word with no stem (a closing alone is not read so); a run that asks what the stem is, or that names a stem of its own, is not read so, but for 거 with a particle under it
package token

func (sc *scan) attached() []*path {
	n := len(sc.s)
	root := &path{verdict: 1}
	sc.attach = true
	var out []*path
	for _, t := range sc.tailsAt(n) {
		switch {
		case t.adn && t.only == string(sc.s):
			out = append(out, grow(root, sc.wordOf(0, n, t)))
		case t.start == 0 && t.base == 0 && t.gate == nil && t.as == "" && t.only == "" && t.shown == "" && tailClass(t.pieces)&predTails == 0:
			out = append(out, grow(root, sc.loneOf(t, n)))
		}
	}
	sc.attach = false
	return out
}
