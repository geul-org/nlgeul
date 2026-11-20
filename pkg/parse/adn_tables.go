package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
)

// stopMarks are the names of the marks that end a sentence, before which
// no adnominal stands (unmodifiable): the full stop, the question mark
// and the exclamation mark. A comma, a middle dot and an ellipsis are
// not here: adnominals are listed with them (경제적, 사회적 요인).
var stopMarks = []string{"마침표", "물음표", "느낌표"}

// coordLinks are the names of the linking endings -고 and -며/-으며 that
// join predicates side by side: a row of adnominals may begin with such
// a predicate (그 맥없고 가련한 얼굴), so one of them may stand between an
// adnominal and its nominal (adnHeadRule).
var coordLinks = closingNames(Marker{Form: "고"}, Marker{Form: "며", Eu: true})

// adnAdverbs are the adverbs of the closed set 꾸밈 사이 부사 that may
// stand between an adnominal and its nominal: those school grammar
// gives as modifying a nominal (바로, 겨우, 이른바), the conjunctive ones
// that join nominals (즉, 또는, 및) and those of degree that modify the
// adnominal after them (가장 큰, 더 많은) (adnHeadRule, adnHeadAfterRule).
var adnAdverbs = wordSets[setAdnAdverb]

// linkClosings are the names of the closings the closing table gives as
// linking ones (연결, 종결|연결 too): a clause of adnominals may begin with
// a predicate's linking form before its own modifying form (다른 남아
// 있는 것), so a word read so right after an adnominal is not judged by
// adnHeadAfterRule.
var linkClosings = closingNamesOf(rule.KindLink)
