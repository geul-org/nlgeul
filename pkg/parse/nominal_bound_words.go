package parse

import (
	"slices"
)

// nominalBound are the dependent nouns of the dependent-noun table that
// stand after a nominal and bear the word table's name of a noun (이름표
// -: 때문, 중, 내, 측, 분 …), not the closed set's (nominalBoundRule).
var nominalBound = slices.DeleteFunc(slices.Clone(afterNounWords), func(w string) bool {
	return slices.Contains(ClosedWords()[boundSet], w)
})
