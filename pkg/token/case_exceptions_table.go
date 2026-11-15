package token

import "github.com/geul-org/nlgeul/pkg/data"

// caseExceptions are the rows of the case exception table
// (case-exceptions.tsv in the data package), by the name of the graph
// each attacks. A row is a word that a marker graph must not read
// although the graph's rule holds for it, found case by case (a word of
// the word table that loses to the rule's reading) or listed by the
// norms, with the kind of its ground and the ground. Each marker graph
// whose name has rows gets one exception of its own (markerGraph), apart
// from the closed-set words that end with its form. Every graph a row
// names must be a marker graph of the index built with a word table
// (checkCases).
var caseExceptions = data.Must(loadCases(data.Must(data.Rows("case-exceptions", 6))))
