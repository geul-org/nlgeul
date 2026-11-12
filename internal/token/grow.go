//ff:func feature=token type=reader control=sequence
//ff:what grow — a reading with one more word: the least verdict so far (section 12: a reading holds as much as its weakest piece), the summed preference and whether any word of it is read from letters written wrong; an auxiliary verb written right after the 어/아 form of its main verb has the cost of the boundary between them given back
package token

import "math"

func grow(p *path, e step) *path {
	sig := e.sig
	if p.prev != nil || p.sig != "" {
		sig = p.sig + " " + e.sig
	}
	pref := p.pref + e.pref
	if p.last.conn {
		pref += e.back
	}
	return &path{prev: p, last: e, sig: sig, verdict: math.Min(p.verdict, e.verdict), pref: pref, typo: p.typo || e.typo}
}
