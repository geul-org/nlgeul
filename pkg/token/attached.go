//ff:func feature=token type=reader control=iteration dimension=1
//ff:what attached — the candidates of a whole unit as markers written apart from their word (로, 에서, 인데 after a space, and 게, 건, 걸, 테니 after a modifying form): each particle chain, and each tail of the copula, that covers the unit from its first letter, read by a scan of its own that takes a stem to stand before the unit (attach), is one word with no stem (a closing alone is not read so); a run that asks what the stem is, or that names a stem of its own, is not read so, but for 거 with a particle under it, which is the whole unit, and 터 with the copula's stem in it (테, 텐), which is the unit's head before a predicate tail
package token

func (sc *scan) attached() []*path {
	n := len(sc.s)
	root := &path{verdict: 1}
	at := *sc
	at.attach = true
	var out []*path
	for _, t := range at.tailsAt(n) {
		switch {
		case t.again != nil && (t.only == string(sc.s) || (t.close != nil && t.only == string(sc.s[:t.start]))):
			out = append(out, at.rootGrown(root, n, t)...)
		case t.start == 0 && t.base == 0 && t.gate == nil && t.as == "" && t.only == "" && t.shown == "" && tailClass(t.pieces)&predTails == 0:
			out = append(out, grow(root, at.loneOf(t, n)))
		}
	}
	return out
}
