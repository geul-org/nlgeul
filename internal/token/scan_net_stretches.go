//ff:func feature=token type=reader control=iteration dimension=1
//ff:what netStretches — the letters of a unit read by the net and cut at the places every way agrees on: one stretch for each portion, in order; none for a unit that no way reads
package token

func (sc *scan) netStretches(u unit) []unitStretch {
	lat := sc.net()
	if lat == nil {
		return nil
	}
	var out []unitStretch
	for k, pt := range lat.portions() {
		out = append(out, sc.stretchOf(u, pt, k == 0))
	}
	return out
}
