//ff:func feature=token type=reader control=iteration dimension=2
//ff:what chains — every particle chain that ends at j (chainMax particles at most), read right to left through the index; each reading of a particle is followed back by its own kind; an auxiliary particle is read right after a linking closing too, and the emphatic 가 only after another particle
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// next is the kind of the particle to the right (-1 for none), depth
// the number of particles already read there.
func (sc *scan) chains(j, depth, next int) []tail {
	if j < 1 || depth == chainMax {
		return nil
	}
	pieces, from, starts := sc.readGroup(classParticle, j, next, false)
	forms, merged, kinds := byForm(pieces, from)
	var out []tail
	for _, f := range forms {
		pc, start := merged[f], starts[entryGroup(kinds[f][0])]
		var here []tail
		if kinds[f][0].slot != rule.LinkEmph {
			here = []tail{{start: start, pieces: []piece{pc}}}
		}
		for _, slot := range slotsOf(kinds[f]) {
			here = append(here, joined(sc.chains(start, depth+1, slot), pc)...)
			here = append(here, joined(sc.fixed(start, slot), pc)...)
		}
		out = append(out, withGate(here, kinds[f][0].gate)...)
		out = append(out, joined(sc.loneHonors(start, false), pc)...)
		out = append(out, joined(sc.units(start, 0, false), pc)...)
		out = append(out, sc.afterClosings(start, depth, pc, kinds[f])...)
	}
	return out
}
