package token

import (
	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// oneStem is 한 standing alone, which is three words: the determiner
// (한 사람, 한 번), 하다 with the modifying closing ㄴ (일을 한 사람),
// and the noun 한(限) (-는 한). oneDeterminer is its name as the
// determiner of the closed set (관형사:한), the one the rules of a
// sentence prefer (determinerOneRule).
const oneStem = "한"

var oneDeterminer = determinerLabel + ":" + oneStem

// haTakes are the names of the last piece of the word before 한 that
// make that word a complement of 하다, so that 한 there is 하다 with
// the modifying ㄴ and not the determiner (determinerOneRule): the
// object particle (일을 한), the adverbial 로/으로 of "…을 …으로 하다"
// (전제로 한), and the closings 하다 follows as an auxiliary verb or as
// a verb of saying: -게 하다, -도록 하다, -어야 하다, -려(고) 하다,
// -고자 하다, and the quoting -다고/-ㄴ다고/-는다고/-라고/-냐고/-자고
// 하다.
var haTakes = append([]string{"목적격:을", "목적격:를", "부사격:로", "부사격:으로"}, closingNames(
	parse.Marker{Form: "게"}, parse.Marker{Form: "도록"}, parse.Marker{Form: "야", Lead: rule.LeadVowel},
	parse.Marker{Form: "려", Eu: true}, parse.Marker{Form: "려고", Eu: true}, parse.Marker{Form: "고자"},
	parse.Marker{Form: "다고"}, parse.Marker{Form: "다고", Lead: rule.JongN}, parse.Marker{Form: "는다고"},
	parse.Marker{Form: "라고"}, parse.Marker{Form: "냐고"}, parse.Marker{Form: "자고"},
)...)

// oneNoun is the name of 한 read as a nominal of the word table, the
// noun 한(限) of the construction -는 한 (limitOneRule), and neunAdn
// the name of the modifying closing 는 that stands before it.
var (
	oneNoun = keyKind + ":" + kindLabel[kindNominal]
	neunAdn = adnPrefix + "는"
)
