//ff:func feature=mini type=rule control=sequence
//ff:what noteSplit — remember, for a claim that stands, how the word divides: at its particle chain, or as the longest tail reading a standing claim rests on
package mini

import "slices"

func noteSplit(r *ruling, name string, t tailFact) {
	if slices.Contains(chainClaims, name) {
		r.chain = true
		return
	}
	if t.size > r.size {
		r.tail, r.size = t.marked, t.size
	}
}
