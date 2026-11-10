//ff:func feature=token type=reader control=sequence
//ff:what grow — a reading with one more word: the least verdict so far (section 12: a reading holds as much as its weakest piece) and the summed preference
package token

import "math"

func grow(p *path, e eojeol) *path {
	sig := e.marked
	if p.prev != nil || p.sig != "" {
		sig = p.sig + " " + e.marked
	}
	return &path{prev: p, last: e, sig: sig, verdict: math.Min(p.verdict, e.verdict), pref: p.pref + e.pref}
}
