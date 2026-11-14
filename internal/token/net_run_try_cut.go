//ff:func feature=token type=reader control=sequence
//ff:what tryCut — a stretch cut off the front of what is worked out, when the window allows it: the window cuts, the places worked out are not the unit's end and lie windowSize tokens or more from the last cut, and the ways of the states at the window's end (front) agree on a first part (agreed) that a stretch of its own reads the same way (cutAt); false when it does not
package token

func (r *netRun) tryCut() (unitStretch, bool) {
	if !r.window || r.j >= r.n || r.sc.tokens(r.c, r.j) < windowSize {
		return unitStretch{}, false
	}
	return r.cutAt(r.agreed())
}
