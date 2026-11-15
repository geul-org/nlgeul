package token

import (
	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// promiseClosings are the final closing 마 of a promise (주마, 가마;
// 으마 after a consonant other than ㄹ: 먹으마), which the standard
// dictionary gives after a verb's stem, read only with a word table,
// after a predicate's stem of the table, as haoClosings are: 마 begins
// and ends many words (마당, 마음, 하마), and as a table row it would
// stand after any stem that is not in the table, inside a word written
// with no spaces (고양이가+마 in 고양이가마당). The closing set of verbs
// only (closing-sets.tsv) keeps it off an adjective's stem.
var promiseClosings = []parse.Marker{
	{Form: "마", Eu: true, Kind: rule.KindFinal, NoIda: true},
}
