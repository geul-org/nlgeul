//ff:func feature=token type=reader control=iteration dimension=2
//ff:what chains — every particle chain that ends at j, as long as the kinds of its particles follow one another (rule.Links: no length is set), read right to left through the index; each reading of a particle is followed back by its own kind; an auxiliary particle is read right after a linking closing too, and the emphatic 가 only after another particle (not after an honorific, a number's unit, or a closing); the particle right on a stem is read as onStem gives it (asked again where a sign that ends a sentence closes the unit), and each run asks the word before it to pass the graph its readings share (sharedGate)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// next is the kind of the particle to the right (-1 for none), depth
// the number of particles already read there.
func (sc *scan) chains(j, depth, next int) []tail {
	if j < 1 {
		return nil
	}
	pieces, from, starts := sc.readGroup(classParticle, j, next, false)
	forms, merged, kinds := byForm(pieces, from)
	ends, endKinds := sc.onStem(j, next, merged, kinds)
	var out []tail
	for _, f := range forms {
		pc, start := merged[f], starts[entryGroup(kinds[f][0])]
		if end, ok := ends[f]; ok && kinds[f][0].slot != rule.LinkEmph {
			out = append(out, withGate([]tail{{start: start, pieces: []piece{end}}}, sharedGate(endKinds[f]))...)
		}
		var here []tail
		for _, slot := range slotsOf(kinds[f]) {
			here = append(here, joined(sc.chains(start, depth+1, slot), pc)...)
			here = append(here, joined(sc.fixed(start, slot), pc)...)
		}
		out = append(out, withGate(here, sharedGate(kinds[f]))...)
		if kinds[f][0].slot == rule.LinkEmph {
			continue
		}
		out = append(out, joined(sc.loneHonors(start, false), pc)...)
		out = append(out, joined(sc.units(start, 0, false), pc)...)
		out = append(out, sc.afterClosings(start, depth, pc, kinds[f])...)
	}
	return out
}
