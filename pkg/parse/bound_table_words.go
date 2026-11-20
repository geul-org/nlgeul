package parse

import (
	"slices"
)

// tableBound are the dependent nouns of the dependent-noun table that
// stand after a modifying closing and bear the word table's name of a
// noun (이름표 -: 리, 바, 지, 터 …), not the closed set's, and that do not
// stand after a nominal as well (boundTableRule).
var tableBound = slices.DeleteFunc(slices.Clone(boundLists[adnPlace]), func(w string) bool {
	return slices.Contains(ClosedWords()[boundSet], w) || slices.Contains(boundNominal, w)
})
