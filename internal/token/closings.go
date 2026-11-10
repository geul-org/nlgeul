//ff:func feature=token type=reader control=iteration dimension=1
//ff:what closings — every predicate tail that ends at j: a closing, alone or with the slots before it
package token

import (
	"sort"

	"github.com/geul-org/koparse/pkg/rule"
)

// more tells that 요 follows. A closing that sits in the stem's last
// syllable takes no slot before it.
func (sc *scan) closings(j int, more bool) []tail {
	if j < 1 {
		return nil
	}
	pieces, from, starts := sc.readGroup(classClosing, j, -1, more)
	keys := make([]string, 0, len(pieces))
	for k := range pieces {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var out []tail
	for _, k := range keys {
		pc, e := pieces[k], from[k]
		if e.lead != 0 {
			out = append(out, sc.fusedTails(pc, e, starts[k])...)
			continue
		}
		out = append(out, tail{start: starts[k], pieces: []piece{pc}, noYo: e.noYo})
		for _, t := range joined(sc.slots(starts[k], rule.SlotClosing), pc) {
			t.noYo = e.noYo
			out = append(out, t)
		}
	}
	return out
}
