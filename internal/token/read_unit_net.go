//ff:func feature=token type=reader control=iteration dimension=1
//ff:what readUnitNet — one unit read as readUnit reads it (the signs before its letters, the stretches its ways agree on, the signs after, and its trace), but the letters by the net, every stretch read once however many ways there are to cut it (the unit's whole, with a window that cuts it at each word the best ways agree on: netRun); the signs at the edges stay contests, and a sign of a pair with letters written onto it has its letters read by the net beside the unit as a whole; all the stretches at once, as a stream takes them one by one (unitReader)
package token

func readUnitNet(ix *index, u unit, trace bool) ([]unitStretch, []string) {
	r := newUnitReader(ix, u, trace, &netStats{})
	var out []unitStretch
	for st, ok := r.next(); ok; st, ok = r.next() {
		out = append(out, st)
	}
	return out, r.take()
}
