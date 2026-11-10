//ff:func feature=token type=reader control=iteration dimension=1
//ff:what chains — every particle chain that ends at j (three particles at most), read right to left through the index
package token

import "sort"

// next is the place of the particle to the right (-1 for none), depth
// the number of particles already read there.
func (sc *scan) chains(j, depth, next int) []tail {
	if j < 1 || depth == 3 {
		return nil
	}
	pieces, from, starts := sc.readGroup(classParticle, j, next, false)
	keys := make([]string, 0, len(pieces))
	for k := range pieces {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var out []tail
	for _, k := range keys {
		pc, start := pieces[k], starts[k]
		out = append(out, tail{start: start, pieces: []piece{pc}})
		out = append(out, joined(sc.chains(start, depth+1, from[k].slot), pc)...)
	}
	return out
}
