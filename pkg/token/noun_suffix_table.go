package token

import "github.com/geul-org/nlgeul/pkg/data"

// nounSuffixes are the Sino-Korean suffixes read right after a nominal
// of the word table (보장+률), one marker each, named 접미:<form>. They
// stand where 들 stands in a chain: particles and the copula follow
// them (보장+률+이+다). They are read only with a word table (addWide),
// the word before must be a nominal of it, and a suffix is not read
// where the stem and the suffix together are a word of the table
// (합격률): the graph suffixGateOf. 률 is read after a final consonant
// other than ㄴ (「한글 맞춤법」 제11항 붙임 1: after a vowel or ㄴ it is
// written 율). They have no count: each costs what its syllables would
// cost as stem syllables.
var nounSuffixes = data.Must(data.Each("noun-suffixes", 2, nounSuffixOf))
