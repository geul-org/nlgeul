//ff:func feature=token type=reader control=iteration dimension=1
//ff:what closingTails — the tails one closing read at one place begins: itself (once for each stem a fused one may come from), and itself after each run of slots
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) closingTails(pc piece, e *entry, start int) []tail {
	ps := partsOf(pc, splits[e.name])
	out := []tail{{start: start, pieces: ps}}
	if e.lead != 0 {
		out = sc.fusedTails(pc, e, start)
	} else {
		out = append(out, joined(sc.slots(start, rule.SlotClosing), ps...)...)
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
