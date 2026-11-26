//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what addParticles — one entry, and so one graph, for each case candidate of each particle of the table, with the observed count this layer keeps for its form, and for the suffixes only this layer reads; the vocative 야 is given the graph the word before it must pass (calledGate), the suffix 이 read where one is called the graph its name must pass (namedGate), and a shortened particle (shortenedParticles: 서) the graph that keeps it off a stem it makes a word of the word table with (joinedGateOf)
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// ownParticles are the particles and suffixes read here and not by
// ParticleMarkers. 이 comes after a
// name that ends in a consonant (영숙이가; morphemes.md section 1) and
// stands where 들, 씩, 쯤 and 끼리 stand, before the particles. It is
// read only with a particle after it: alone it looks the same as the
// subject particle 이. The emphatic 야 after a vowel (에야, 이제야;
// particles.md section 3 gives 이야 / 야) is in the particle table only
// as 이야. The tables give neither a count. 은커녕 and 는커녕 are one
// particle each (the table has 커녕 alone); their count is 커녕's. 가
// after an adverbial particle (에서가 아니라) is an auxiliary particle
// of emphasis, read only there; so is 를 after 에 and 으로 (집에를,
// emphObjName). The last 야 is the same auxiliary particle after an
// ending that ends in a consonant (간다면야: the dictionary gives 야
// after an ending, and 이야 only after a nominal or an adverbial that
// ends in a consonant); it is read only right after a closing (ended).
// So are the 가 of emphasis after the linking 지 (먹지가 않다: the
// dictionary's 가 after the linking ending -지) and 그려, which the
// dictionary gives after some final endings (갑니다그려): both are read
// only right after the closings of their rows (closingParticles).
const emphName = "보조사:가(강조)"

// comparing, together, directional and shortTopics tell kinds of particle apart for
// chains (linkOf): the particles of comparison, those of company, and
// the shortened forms that end in 는.
var comparing = []string{"보다", "처럼", "같이", "만큼"}
var together = []string{"과", "와", "이랑", "랑"}
var directional = []string{"으로", "로", "으로부터", "로부터"}
var shortTopics = []string{"엔", "에선", "에겐", "한텐", "까진", "보단"}

// leading are the auxiliary particles that set a word forth for what
// is then said of it (너야 물론 합격이지): 이야 after a consonant, 야
// after a vowel. A sentence does not as a rule end on one, so where a
// full stop, a question mark or an exclamation mark closes the unit
// right after such a particle that stands right on its stem it is not
// read (문제야. is 문제+(이)+야). After another particle or a closing it
// stays (에서야, 하고 나서야).
// Dots in a row trail off and close nothing (나야 뭐...). A closed
// table: only the candidates listed are held back.
var leading = []string{"보조사:이야", "보조사:야"}

// echoes are the particle candidates whose letters, right on a stem that
// ends in a syllable with no final consonant and before the auxiliary 요,
// are those of a closing written wrong: 고양이에요 has 에 and 요 as
// well as 예요 written 에요. The candidate is rebutted there (forceEcho)
// and stays. A closed table: only the candidates listed are rebutted.
var echoes = []string{"부사격:에"}

// vocName is the vocative 야, written onto the name of who is called
// (철수야). A dependent noun names no one that can be called, so the
// word before it must pass the graph callGate.
const (
	vocName  = "호격:야"
	callGate = "앞말:부르는 말"
)

// nameSuffix is the suffix 이 written after a person's name that ends
// in a consonant (the dictionary's suffix -이: 영숙이, 갑돌이). With a
// particle after it, it is read wherever it stands (inner); with none it
// looks the same as the subject particle 이, and is read only where one
// who is called is written: at the end of its unit's letters, right
// before a comma or an exclamation mark (called: 영숙이, 영숙이!). There
// the word before it must pass the graph nameGate: no dependent noun,
// and no word that stands as an adverb, a determiner, a predicate or an
// interjection of the word table (a name is a nominal or a new word).
// The subject and the complement candidates stay beside it.
const (
	nameSuffix = "접미:이"
	nameGate   = "앞말:부르는 이름"
)

var ownParticles = []entry{
	{name: nameSuffix, form: "이", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkSuffix, inner: true},
	{name: nameSuffix, form: "이", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkSuffix, called: true},
	{name: "보조사:야", form: "야", class: classParticle, after: rule.AfterVowel, slot: rule.LinkTopic},
	{name: emphName, form: "가", class: classParticle, slot: rule.LinkEmph, count: 2},
	{name: "보조사:은커녕", form: "은커녕", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkTopic, count: 66},
	{name: "보조사:는커녕", form: "는커녕", class: classParticle, after: rule.AfterVowel, slot: rule.LinkTopic, count: 66},
	{name: emphObjName, form: "를", class: classParticle, after: rule.AfterVowel, slot: rule.LinkEmph},
	{name: "보조사:야", form: "야", class: classParticle, after: rule.AfterBatchim, slot: rule.LinkTopic, ended: true},
	{name: "보조사:가", form: "가", class: classParticle, slot: rule.LinkTopic, ended: true},
	{name: "보조사:그려", form: "그려", class: classParticle, slot: rule.LinkTopic, ended: true},
}

func (ix *index) addParticles() {
	ix.call = calledGate()
	joined := joinedGateOf(joinedGate, 0)
	named := namedGate()
	ix.graphs, ix.edges = ix.graphs+3, ix.edges+4
	for _, m := range ParticleMarkers() {
		for _, label := range caseNames(m.Kase, m.Slot) {
			e := &entry{name: label + ":" + m.Form, form: m.Form, class: classParticle, after: m.After, slot: linkOf(label, m.Form, m.Slot), kase: m.Kase, count: particleCount(m)}
			if e.name == vocName {
				e.gate = ix.call
			}
			if slices.Contains(shortenedParticles, m.Form) {
				e.gate = joined
			}
			ix.add(e)
		}
	}
	for i := range ownParticles {
		e := ownParticles[i]
		if e.called {
			e.gate = named
		}
		ix.add(&e)
	}
}
