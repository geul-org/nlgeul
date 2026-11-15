package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/parse"
)

// boundNext are the dependent nouns a modifying form stands right
// before (boundBeforeRule): the closed set of the grammar tables and
// the spacing words' that follow a modifying closing.
var boundNext = slices.Concat(parse.ClosedWords()[boundSet], spacingSets[setBoundAdn])
