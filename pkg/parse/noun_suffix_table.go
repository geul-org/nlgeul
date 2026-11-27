package parse

import "github.com/geul-org/nlgeul/pkg/data"

// nounSuffixes are the Sino-Korean suffixes read right after a nominal
// of the word table (보장+률, 팀+당), one marker each, named 접미:<form>.
// 률 and 율 stand where 들 stands in a chain: particles and the copula
// follow them (보장+률+이+다). 당 (當, for each: 팀당 두 명, 가구당 소득; the
// standard dictionary's suffix -당 after a number or a unit, here after
// a nominal that is counted by) ends its word: no particle is read after
// it (rule.LinkLast), as the dictionary's examples have it before the
// amount it tells of (마리당 삼천 원, 시간당 얼마), so that a word of the
// noun 당 (堂, 黨) with a particle is not cut so (경로당에). 사 is the
// suffix written onto a name or a nominal for a company or a body, a
// person of a trade, a temple, a history and so on (the dictionary's
// -사 社, 師, 士, 寺, 史 …: 구글사, 전담사); it has one name, 접미:사,
// whatever it means, and particles and the copula follow it. 용 (用, for
// a use or a user: 사업용, 파괴용) is read the same way, and after a run of
// Latin letters and digits the table lists no word of (OTS용: suffixedRun)
// every suffix is read as after a nominal. 별 (別, each: 사업자별로, 필지별) is
// read the same way after any nominal of the table, one syllable too
// (시별); the noun 별 (a star) stays a candidate. It is not
// read either where the stem and the suffix are two words of the table
// written together (twoWords: 기동+전사, not 기동전+사). They are read only with a word table (addWide),
// the word before must be a nominal of it, and a suffix is not read
// where the stem and the suffix together are a word of the table
// (합격률): the graph suffixGateOf. 률 is read after a final consonant
// other than ㄴ and 율 after a vowel or ㄴ (「한글 맞춤법」 제11항 붙임 1:
// after a vowel or ㄴ 률 is written 율: 증가율, 출산율); 당 and 사 after anything.
// They have no count: each costs what its syllables would cost as stem syllables.
var nounSuffixes = data.Must(data.Each("noun-suffixes", 3, nounSuffixOf))

// numberSuffixes are the suffixes of nounSuffixes read right after a
// number too (number-suffixes.tsv): 사 (three companies named and then
// 3사, 제2사) and 용 (4용). A number is a nominal (a numeral), so the gate of a
// suffix passes it as one (suffixedNumber); after a number and its unit the
// suffix is read as any particle chain is (6개사). The noun 사 read
// apart from the number stays a candidate.
var numberSuffixes = data.Must(data.Column("number-suffixes", 1))

// longSuffixes are the suffixes of nounSuffixes read only after a stem
// of two syllables or more (long-stem-suffixes.tsv): 용 (not 별). A word of one
// syllable and such a suffix is a word of its own that the word table
// holds (군용), and where it holds none the letters are another word's:
// shortStem.
var longSuffixes = data.Must(data.Column("long-stem-suffixes", 1))
