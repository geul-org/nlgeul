//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addParticles — one entry, and so one graph, for each case reading of each particle of the table
package token

import "github.com/geul-org/koparse/internal/parse"

func (ix *index) addParticles() {
	for _, m := range parse.ParticleMarkers() {
		for _, label := range caseNames(m.Kase, m.Slot) {
			ix.add(&entry{name: label + ":" + m.Form, form: m.Form, class: classParticle, after: m.After, slot: m.Slot, count: m.Count})
		}
	}
}
