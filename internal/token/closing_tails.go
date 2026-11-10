//ff:func feature=token type=reader control=iteration dimension=1
//ff:what closingTails — the tails one closing read at one place begins: itself (once for each stem a fused one may come from), itself after each run of slots, and, for a closing that takes the link vowel 으 as a piece, itself after 으
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) closingTails(pc piece, e *entry, start int) []tail {
	ps := partsOf(pc, splits[e.name])
	out := []tail{{start: start, pieces: ps, head: e.head}}
	if e.lead != 0 {
		out = sc.fusedTails(pc, e, start)
	} else {
		out = append(out, joined(sc.slots(start, rule.SlotClosing), ps...)...)
		out = append(out, sc.linked(start, ps, e.name, true)...)
	}
	out = withGate(out, e.gate)
	for i := range out {
		out[i].noYo, out[i].close = e.noYo, e
		if e.needs {
			out[i].only = notOnly
		}
	}
	return out
}
