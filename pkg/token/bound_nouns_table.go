package token

import "github.com/geul-org/nlgeul/pkg/data"

// boundLists are the dependent nouns of the dependent-noun table
// (bound-nouns.tsv) by where they stand and by what else reads them
// (boundPlaces, boundUses). The table is the one list of dependent nouns:
// pkg/parse reads from it which bear the name of the closed set
// (ClosedWords) and which it reads written onto a modifying closing, and
// this layer the lists made from it (afterNounWords, dependentWords,
// boundNext, boundNominal, adnBound, adnOnto).
var boundLists = data.Must(loadBoundNouns(data.Must(data.Rows("bound-nouns", 5))))
