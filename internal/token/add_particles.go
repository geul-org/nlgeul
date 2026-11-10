//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addParticles — one entry, and so one graph, for each case reading of each particle of the table, with the observed count this layer keeps for its form, and for the suffixes only this layer reads
package token

import (
	"github.com/geul-org/koparse/internal/parse"
	"github.com/geul-org/koparse/pkg/rule"
)

// innerSuffixes are read here and not by kopas parse. 이 comes after a
// name that ends in a consonant (영숙이가; morphemes.md section 1) and
// stands where 들, 씩, 쯤 and 끼리 stand, before the particles. It is
// read only with a particle after it: alone it looks the same as the
// subject particle 이. The emphatic 야 after a vowel (에야, 이제야;
// particles.md section 3 gives 이야 / 야) is in the particle table only
// as 이야. The tables give neither a count.
// comparing, together, directional and shortTopics tell kinds of particle apart for
// chains (linkOf): the particles of comparison, those of company, and
// the shortened forms that end in 는.
var comparing = []string{"보다", "처럼", "같이", "만큼"}
var together = []string{"과", "와", "이랑", "랑"}
var directional = []string{"으로", "로", "으로부터", "로부터"}
var shortTopics = []string{"엔", "에선", "에겐", "한텐", "까진", "보단", "선", "론"}

var innerSuffixes = []entry{
	{name: "접미:이", form: "이", class: classParticle, after: afterBatchim, slot: rule.LinkSuffix, inner: true},
	{name: "보조사:야", form: "야", class: classParticle, after: afterVowel, slot: rule.LinkTopic},
}

func (ix *index) addParticles() {
	for _, m := range parse.ParticleMarkers() {
		for _, label := range caseNames(m.Kase, m.Slot) {
			ix.add(&entry{name: label + ":" + m.Form, form: m.Form, class: classParticle, after: m.After, slot: linkOf(label, m.Form, m.Slot), kase: m.Kase, count: particleCount(m)})
		}
	}
	for i := range innerSuffixes {
		e := innerSuffixes[i]
		ix.add(&e)
	}
}
