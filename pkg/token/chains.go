//ff:func feature=token type=reader control=iteration dimension=1
//ff:what chains — every particle chain that ends at j, as long as the kinds of its particles follow one another (rule.Links: no length is set), read right to left through the index; each candidate of a particle (one for each name) is followed back by its own kind, and the word before it passes its own graph; an auxiliary particle is read right after a linking closing too, and the emphatic 가 only after another particle (not after an honorific, a number's unit, or a closing); the particle right on a stem is read as onStem gives it (asked again where a sign that ends a sentence closes the unit)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// next is the kind of the particle to the right (-1 for none), depth
// the number of particles already read there.
func (sc *scan) chains(j, depth, next int) []tail {
	if j < 1 {
		return nil
	}
	pieces, from, starts := sc.readGroup(classParticle, j, next, false)
	ends, endFrom := sc.onStem(j, next, pieces, from)
	var out []tail
	for _, k := range sc.groupKeys(pieces, from, j) {
		pc, e, start := pieces[k], from[k], starts[k]
		if end, ok := ends[k]; ok && e.slot != rule.LinkEmph {
			out = append(out, withGate([]tail{{start: start, pieces: []piece{end}}}, endFrom[k].gate)...)
		}
		here := joined(sc.chains(start, depth+1, e.slot), pc)
		here = append(here, joined(sc.fixed(start, e.slot), pc)...)
		out = append(out, withGate(here, e.gate)...)
		if e.slot == rule.LinkEmph {
			continue
		}
		out = append(out, joined(sc.loneHonors(start, false), pc)...)
		out = append(out, joined(sc.units(start, 0, false), pc)...)
		out = append(out, sc.afterClosings(start, depth, pc, e)...)
	}
	return out
}
