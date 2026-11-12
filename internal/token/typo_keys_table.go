package token

import (
	"github.com/geul-org/nlgeul/data"
)

// typoKeys is the closed table of the wrong ways to write a word of
// a closed set that this layer reads (a written form to the word it
// stands for): 꺼 for the dependent noun 거. The table is typo-words.tsv
// in the data package; it is closed, and its lines are all of it, not
// taken from any other list or dictionary. It is read when there is a
// word table (addTypoKeys).
var typoKeys = data.Must(data.Map("typo-words"))
