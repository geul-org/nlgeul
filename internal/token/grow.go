//ff:func feature=token type=reader control=sequence
//ff:what grow — a reading with one more word: the least verdict so far (section 12: a reading holds as much as its weakest piece) and the summed preference; an auxiliary verb written right after the 어/아 form of its main verb has the cost of the boundary between them given back
package token

import "math"

func grow(p *path, e eojeol) *path {
	sig := e.marked
	if p.prev != nil || p.sig != "" {
		sig = p.sig + " " + e.marked
	}
	pref := p.pref + e.pref
	if p.last.conn {
		pref += e.back
	}
	return &path{prev: p, last: e, sig: sig, verdict: math.Min(p.verdict, e.verdict), pref: pref}
}
