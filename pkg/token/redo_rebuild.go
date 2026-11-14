//ff:func feature=token type=reader control=iteration dimension=1
//ff:what rebuild — the region of a stretch a window cut off, read again: a copy of the unit's scan (what was prepared is the same; the registry, the candidates kept at signs and the markers are its own, and the trace is none) is put on the seeds it grew from (the root for the first stretch of a unit) at the place the stretch begins and worked out to the place it ends (with the candidates kept there, when a sign asks for them), the stretch is made from the states between (the same stretchOf) and the region from it at the place it has in the sentence; it is the region the cut gave, as far as the words and their order go (the numbers of birth are counted again, in the same order)
package token

func (d *redo) rebuild() *region {
	sc := *d.sc
	sc.trace, sc.floor, sc.marks = nil, d.from, make([]map[int]bool, len(sc.s)+1)
	sc.reg, sc.kept = newRegistry(d.base), map[int][]*path{}
	best := make([][]*path, d.to+1)
	best[d.from] = []*path{sc.reg.root}
	if len(d.seeds)+len(d.keptSeeds) > 0 {
		best[d.from], sc.kept[d.from], sc.reg.seeds = d.seeds, d.keptSeeds, map[*path]bool{}
		for _, sd := range append(append([]*path{}, d.seeds...), d.keptSeeds...) {
			sc.reg.seeds[sd] = true
		}
	}
	for j := d.from + 1; j <= d.to; j++ {
		sc.netStep(best, j, nil)
	}
	if sc.late(d.to) {
		sc.keptAt(best, d.to)
	}
	st := sc.stretchOf(d.u, sc.between(d.from, d.to), d.first)
	return newRegionNet(d.head, d.at, *st.cells)
}
