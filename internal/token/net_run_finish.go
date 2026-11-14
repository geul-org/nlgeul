//ff:func feature=token type=reader control=iteration dimension=1
//ff:what finish — the stretches of what is left of the unit when its end is reached: the live states from the last cut cut at the places every way agrees on, one stretch for each portion, as netStretches makes them (the first of a unit that was not cut touches the head); when nothing reads what is left after a cut, the letters are one stretch with no reading
package token

func (r *netRun) finish() []unitStretch {
	var out []unitStretch
	for k, pt := range r.sc.netClose(r.best, r.c).portions() {
		out = append(out, r.sc.stretchOf(r.u, pt, r.first && r.c == 0 && k == 0))
	}
	if len(out) == 0 && r.c > 0 {
		out = append(out, r.unread(r.c, r.n))
	}
	return out
}
