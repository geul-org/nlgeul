//ff:func feature=parse type=reader control=sequence
//ff:what grow — a candidate with one more word: the least verdict so far (section 12: a candidate holds as much as its weakest piece) and the summed preference
package parse

import "math"

func grow(p *path, e step) *path {
	sig := e.sig
	if p.prev != nil || p.sig != "" {
		sig = p.sig + " " + e.sig
	}
	return &path{prev: p, last: e, sig: sig, verdict: math.Min(p.verdict, e.verdict), pref: p.pref + e.pref}
}
