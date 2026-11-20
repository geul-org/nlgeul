package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
)

// haTakes are the names of the last piece of the word before 한 that
// make that word a complement of 하다, so that 한 there is 하다 with the
// modifying ㄴ (haOneRule): the object particle (일을 한), the adverbial
// 로/으로 of "…을 …으로 하다" (전제로 한), the dative 에게/한테/께 of the
// one a 하다 of saying or doing is done to (그에게 한 말), and the closings 하다 follows as
// an auxiliary verb or as a verb of saying: -게 하다, -도록 하다, -어야
// 하다, -려(고) 하다, -고자 하다, and the quoting -다고/-ㄴ다고/-는다고/
// -라고/-냐고/-자고 하다. haName is the name of the stem of 하다 as the
// word table reads it, and haAdn that of the modifying closing ㄴ after it.
var (
	haTakes = append([]string{"목적격:을", "목적격:를", "부사격:로", "부사격:으로", "부사격:에게", "부사격:한테", "부사격:께"}, closingNames(
		Marker{Form: "게"}, Marker{Form: "도록"}, Marker{Form: "야", Lead: rule.LeadVowel},
		Marker{Form: "려", Eu: true}, Marker{Form: "려고", Eu: true}, Marker{Form: "고자"},
		Marker{Form: "다고"}, Marker{Form: "다고", Lead: rule.JongN}, Marker{Form: "는다고"},
		Marker{Form: "라고"}, Marker{Form: "냐고"}, Marker{Form: "자고"},
	)...)
	haStem = "하"
	haName = keyKind + ":" + kindLabel[kindVerb]
	haAdn  = adnPrefix + "ㄴ"
)
