//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what addMisfit — with a word table, one more entry for each case candidate of each particle of the table that the final consonant of the word before it chooses (misfitAfter), read where the other form of its pair is the right one (민주당는, 골드은), and for the particles of misfitOwn: the name and the count are the particle's own, the graph rebuts the candidate (markerGraph), and the word before it must pass a graph of its own that asks for a nominal of the word table (misfitGate)
package parse

func (ix *index) addMisfit() {
	gate := gateOf(misfitGate, kindNominal)
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	for _, m := range ParticleMarkers() {
		wrong, ok := misfitAfter[m.After]
		if !ok {
			continue
		}
		for _, label := range caseNames(m.Kase, m.Slot) {
			ix.add(&entry{name: label + ":" + m.Form, form: m.Form, class: classParticle, after: wrong, slot: linkOf(label, m.Form, m.Slot), kase: m.Kase, count: particleCount(m), misfit: true, gate: gate})
		}
	}
	for i := range misfitOwn {
		e := misfitOwn[i]
		e.gate = gate
		ix.add(&e)
	}
}
