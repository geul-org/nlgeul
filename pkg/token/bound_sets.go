package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/parse"
)

// boundNext are the dependent nouns a modifying form stands right
// before (boundBeforeRule): the closed set of the grammar tables and
// the spacing words' that follow a modifying closing. boundNominal are
// those of them that stand right after a nominal as well, which tell
// nothing of the word before them: the spacing words' that follow a
// nominal (때, 간, 외, 시 …) and 만 of a length of time (사흘 만에, the
// standard dictionary's 만 「1」: after a word of time).
var (
	boundNext    = slices.Concat(parse.ClosedWords()[boundSet], spacingSets[setBoundAdn])
	boundNominal = append(slices.Clone(afterNounWords), "만")
)
