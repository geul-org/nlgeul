package parse

import "github.com/geul-org/nlgeul/pkg/data"

// A particle is written onto the ending before it (clause 41 of the
// orthography). closingParticles gives, by the name of a particle
// (closing-particles.tsv, the first column), the names of the closings it
// is read right after (the second, joined with |), as the dictionary
// tells of each particle what it is put after. The particle is one more
// piece of the tail of that closing, under that name only, and more
// particles may follow as the chain's kinds allow. A closed table: only
// the pairs listed are read, and a pair that is itself a closing of the
// tables (다가, 는다만) is read as that closing only. Each particle has a
// graph of its own (namedAuxName and the name), whose exception is that
// the closing before it is not one of those listed.
const namedAuxName = "사슬:맺음 뒤 조사 "

var closingParticles = data.Must(data.Map("closing-particles"))
