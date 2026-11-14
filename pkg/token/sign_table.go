package token

import (
	"github.com/geul-org/nlgeul/data"
)

// The signs that are named where they stand at the edges of a unit or
// alone (the full stop has its own graphs: stopNames), each with its
// candidates in the order they are printed. A candidate is a graph of its
// own (signGraph). A sign with no candidate that stands, and a sign that
// is not in the table, is sent up as its letters with no name: no
// letter of a sentence is dropped.
var signNames = data.Must(loadSigns(data.Must(data.Rows("signs", 2))))
var signOrderNamed = data.Must(data.Runes("signs", 2))

// roleNames are the candidates of a sign by what it does for the words
// after it (a heading, a topic, the name of what is meant) and not as
// the sign itself. When only such candidates stand, the sign's letters
// with no name are sent up beside them. headingName is printed with
// the number of the signs in a row and no colon (## is 제목2).
var roleNames = []string{headingName, "태그", "명칭"}

const headingName = "제목"

// arrowSigns stand right next to < or > in a run that is no bracket
// (->, =>, <=, >=).
const arrowSigns = "-="

// stopSigns are the signs that close the word before them inside a
// unit, as a piece of that word (stopAt): the full stop and the comma.
const stopSigns = ".,"

// apartSigns are the signs that, between letters of a unit, are also read
// as a word of their own (a candidate added beside the one that keeps
// them in the word): the middle dots. cutSigns are the signs that cut a
// sentence into units where they stand between letters (cutter).
const (
	apartSigns = "·ㆍ・"
	cutSigns   = "+"
)
