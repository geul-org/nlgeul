package parse

// shortenedParticles are the particles of the particle table that are a
// longer particle shortened: 서, short for 에서 (학교서, 서울서). Where
// the stem and such a particle together are a word of the word table
// (요구서, 확인서, 혼자서), the particle is not read there and the whole
// word stands first: the word before it must pass the graph joinedGate.
// The particles shortened with 는 (shortParticles: 선, 론, 으론) are kept
// off such a stem too, by their nominal gate (joinedGateOf).
var shortenedParticles = []string{"서"}

// joinedGate names the graph the word before a shortened particle of the
// particle table must pass.
const joinedGate = "앞말:줄인 조사 앞"
