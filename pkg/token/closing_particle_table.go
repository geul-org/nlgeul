package token

import "github.com/geul-org/nlgeul/pkg/parse"

// Some particles are written onto a few closings only (clause 41 of
// the orthography: a particle after an ending is written onto it too).
// pickedClosings gives, by the particle's form, the closings it is read
// right after, with every name its form has: 를 after the linking 지 of
// 지 않다 and 지 못하다 (끄지를 못하다, 오지를 않다; the dictionary's 를
// after the linking closing -지). Each form has a graph of its own
// (pickedAuxName and the form), whose exception is that the closing
// before it is not one of those. A closed table: only the pairs listed
// are read.
const pickedAuxName = "사슬:맺음 뒤 "

var pickedClosings = map[string][]string{
	"를": closingNames(parse.Marker{Form: "지"}),
}
