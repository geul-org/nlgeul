package parse

import (
	"github.com/geul-org/nlgeul/data"
)

// The closed sets of morphemes.md section 8, which the table
// closed-words.tsv in the data package holds. They are part of the
// grammar tables: no other word is listed anywhere.

// closedSets are the sets as their table gives them (closed-words.tsv).
var closedSets = data.Must(loadClosed(data.Must(data.Rows(closedFile, 3))))

// boundNouns are the dependent nouns: a word that begins with one makes
// the word before it a modifier.
var boundNouns = closedSets[boundSet]

// joinedNouns are the dependent nouns read when written onto a
// modifying closing (아는거, 할수). 대로, 뿐 and 만큼 are not here: they
// are read there as particles.
var joinedNouns = closedSets[joinedMark]

// shortened is 것 run together with a particle (게 = 것이, 건 = 것은,
// 걸 = 것을), with the case of that particle.
var shortened = map[string]int{"게": caseSubj | caseComp, "건": 0, "걸": caseObj}

// pointers are the demonstratives a shortened 것 is read after.
var pointers = []string{"이", "그", "저"}

// helperStems are the stems of the auxiliary verbs (주다, 보다, 버리다,
// 놓다, 두다, 있다, 지다), read after the 어/아 form of a main verb.
var helperStems = closedSets[helperSet]

// determiners (지시·수 관형사) and conjunctions (접속 부사) are words
// that stand whole.
var determiners = closedSets[determinerSet]
var conjunctions = closedSets[conjunctionSet]

// rangeParticles are the auxiliary particles that also mark an
// adverbial of range.
var rangeParticles = []string{"까지", "부터"}
