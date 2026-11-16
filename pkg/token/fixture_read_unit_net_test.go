//ff:func feature=token type=helper control=iteration dimension=1
//ff:what readUnitNet — one unit read as the stream reads it (newUnitReader: the signs before its letters, the letters by the net, the signs after, and its trace), all the stretches at once instead of one by one
package token

func readUnitNet(ix *index, u unit, trace bool) ([]unitStretch, []string) {
	r := newUnitReader(ix, u, trace, &netStats{})
	var out []unitStretch
	for st, ok := r.next(); ok; st, ok = r.next() {
		out = append(out, st)
	}
	return out, r.take()
}
