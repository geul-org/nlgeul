//ff:func feature=token type=reader control=iteration dimension=1
//ff:what slots — every run of slots (시, 었, 겠, 더, a fused ㅆ) that ends at j, read right to left; a fused one ends the run, 시 is read after the link vowel 으 too, and a slot read only with a word table asks the word before it to pass its graph
package token

import "sort"

// next is the place of what follows (a closing counts as 9).
func (sc *scan) slots(j, next int) []tail {
	if j < 1 {
		return nil
	}
	pieces, from, starts := sc.readGroup(classSlot, j, next, false)
	keys := make([]string, 0, len(pieces))
	for k := range pieces {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var out []tail
	for _, k := range keys {
		pc, e := pieces[k], from[k]
		if e.lead != 0 {
			out = append(out, withGate(sc.fusedTails(pc, e, starts[k]), e.gate)...)
			continue
		}
		here := []tail{{start: starts[k], pieces: []piece{pc}, head: e.head}}
		here = append(here, joined(sc.slots(starts[k], e.slot), pc)...)
		here = append(here, sc.linked(starts[k], []piece{pc}, e.name, false)...)
		out = append(out, withGate(here, e.gate)...)
	}
	return out
}
