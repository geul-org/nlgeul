package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
)

// The closed sets of morphemes.md section 8, which the tables
// closed-words.tsv and bound-nouns.tsv in the data package hold. They
// are part of the grammar tables: no other word is listed anywhere. The
// tables hold what the token layer reads too; this package leaves those
// lines and columns out, and the token layer checks them.

// closedSets are the sets of closed-words.tsv, and boundSets the
// dependent nouns of bound-nouns.tsv that bear the set's name, with
// those of them read when written onto a modifying closing.
var (
	closedSets = data.Must(loadClosed(data.Must(data.Rows(closedFile, 2))))
	boundSets  = data.Must(loadBound(data.Must(data.Rows(boundFile, 5))))
)

// boundNouns are the dependent nouns: a word that begins with one makes
// the word before it a modifier.
var boundNouns = boundSets[boundSet]

// joinedNouns are the dependent nouns read when written onto a
// modifying closing (아는거, 할수). 대로, 뿐 and 만큼 are not here: they
// are read there as particles.
var joinedNouns = boundSets[joinedMark]

// shortened is 것 run together with a particle (게 = 것이, 건 = 것은,
// 걸 = 것을), with the case of that particle.
var shortened = map[string]int{"게": caseSubj | caseComp, "건": 0, "걸": caseObj}

// pointers are the demonstratives a shortened 것 is read after (이, 그,
// 저; the set 지시어 of closed-words.tsv, which the token layer reads too).
var pointers = closedSets[pointerSet]

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
