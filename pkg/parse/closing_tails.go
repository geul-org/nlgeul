//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what closingTails — the tails one closing read at one place begins: itself, written out as its pieces where the index writes it out (splitsOf) (once for each stem a fused one may come from), itself after each run of slots, and, for a closing that takes the link vowel 으 as a piece, itself after 으
package parse

import (
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) closingTails(pc piece, e *entry, start int) []tail {
	ps := partsOf(pc, sc.ix.splitsOf(e.name))
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
		} else if e.lead == 0 && start >= 1 && unicode.IsDigit(sc.s[start-1]) {
			out[i].only = digitOnly
		}
	}
	if e.wide && e.lead == 0 && e.tail == rule.TailVerb && !e.bars {
		out = append(out, notTail(out[0]))
	}
	return out
}
