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

// helperStems are the stems of the auxiliary verbs (주다, 보다, 버리다,
// 놓다, 두다, 있다, 지다), read after the 어/아 form of a main verb.
var helperStems = closedSets[helperSet]

// determiners (지시·수 관형사) and conjunctions (접속 부사) are words
// that stand whole.
var determiners = closedSets[determinerSet]
var conjunctions = closedSets[conjunctionSet]
