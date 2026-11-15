package token

// negAdverbs and numDeterminers are the modifiers a new word is read
// apart at when the known word they modify is written onto them
// (closed-words-more.tsv; modifiesAt): the negative adverbs 안 and 못,
// which stand right before a predicate, and the numeral determiners 한
// and 두 of the closed set of determiners, which stand right before a
// nominal. Clause 2 of the spelling rules writes each word apart. Wider
// sets were measured on the answer sheet against these, and each read
// apart more new words that are a name or an old spelling than it put
// right: every adverb of the word table before a predicate, 1 more wrong
// and none more right; 세 and 네, 2 more wrong and 1 more right; the
// other determiners of the closed set (이 그 저 …), 2 more wrong and none
// more right; every determiner of the word table (총, 전, 딴 …), 18 more
// wrong and 3 more right. A closed table: only what is listed is read so.
var (
	negAdverbs     = wordSets[setNegAdverb]
	numDeterminers = wordSets[setNumDet]
)
