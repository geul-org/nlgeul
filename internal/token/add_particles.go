//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addParticles — one entry, and so one graph, for each case reading of each particle of the table, with the observed count this layer keeps for its form, and for the suffixes only this layer reads
package token

import (
	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// ownParticles are the particles and suffixes read here and not by
// nlgeul parse. 이 comes after a
// name that ends in a consonant (영숙이가; morphemes.md section 1) and
// stands where 들, 씩, 쯤 and 끼리 stand, before the particles. It is
// read only with a particle after it: alone it looks the same as the
// subject particle 이. The emphatic 야 after a vowel (에야, 이제야;
// particles.md section 3 gives 이야 / 야) is in the particle table only
// as 이야. The tables give neither a count. 은커녕 and 는커녕 are one
// particle each (the table has 커녕 alone); their count is 커녕's. 가
// after an adverbial particle (에서가 아니라) is an auxiliary particle
// of emphasis, read only there.
const emphName = "보조사:가(강조)"

// comparing, together, directional and shortTopics tell kinds of particle apart for
// chains (linkOf): the particles of comparison, those of company, and
// the shortened forms that end in 는.
var comparing = []string{"보다", "처럼", "같이", "만큼"}
var together = []string{"과", "와", "이랑", "랑"}
var directional = []string{"으로", "로", "으로부터", "로부터"}
var shortTopics = []string{"엔", "에선", "에겐", "한텐", "까진", "보단"}

var ownParticles = []entry{
	{name: "접미:이", form: "이", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkSuffix, inner: true},
	{name: "보조사:야", form: "야", class: classParticle, after: rule.AfterVowel, slot: rule.LinkTopic},
	{name: emphName, form: "가", class: classParticle, slot: rule.LinkEmph, count: 2},
	{name: "보조사:은커녕", form: "은커녕", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkTopic, count: 66},
	{name: "보조사:는커녕", form: "는커녕", class: classParticle, after: rule.AfterVowel, slot: rule.LinkTopic, count: 66},
}

func (ix *index) addParticles() {
	for _, m := range parse.ParticleMarkers() {
		for _, label := range caseNames(m.Kase, m.Slot) {
			ix.add(&entry{name: label + ":" + m.Form, form: m.Form, class: classParticle, after: m.After, slot: linkOf(label, m.Form, m.Slot), kase: m.Kase, count: particleCount(m)})
		}
	}
	for i := range ownParticles {
		e := ownParticles[i]
		ix.add(&e)
	}
}
