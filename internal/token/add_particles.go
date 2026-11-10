//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addParticles — one entry, and so one graph, for each case reading of each particle of the table, and for the suffixes only this layer reads
package token

import "github.com/geul-org/koparse/internal/parse"

// innerSuffixes are read here and not by kopas parse. 이 comes after a
// name that ends in a consonant (영숙이가; morphemes.md section 1) and
// stands where 들, 씩, 쯤 and 끼리 stand, before the particles. It is
// read only with a particle after it: alone it looks the same as the
// subject particle 이. The tables give it no count.
var innerSuffixes = []entry{
	{name: "접미:이", form: "이", class: classParticle, after: afterBatchim, slot: 0, inner: true},
}

func (ix *index) addParticles() {
	for _, m := range parse.ParticleMarkers() {
		for _, label := range caseNames(m.Kase, m.Slot) {
			ix.add(&entry{name: label + ":" + m.Form, form: m.Form, class: classParticle, after: m.After, slot: m.Slot, kase: m.Kase, count: m.Count})
		}
	}
	for i := range innerSuffixes {
		e := innerSuffixes[i]
		ix.add(&e)
	}
}
