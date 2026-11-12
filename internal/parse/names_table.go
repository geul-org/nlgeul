package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The names the tables of the data package are written with (one tsv
// file a table), each with the value it stands for. They are closed
// sets: a name that is not here is an error when the table is read.
//
// AfterNames tell what must stand before a form of whole syllables;
// KindNames the kinds of a closing; CaseNames the cases a particle
// marks; FlagNames a closing's ways of attaching (the Marker's fields
// Eu, NoIda, NoYo, NeedIda and SelfIda).
var AfterNames = map[string]int{
	"-":         rule.AfterAny,
	"받침 뒤":      rule.AfterBatchim,
	"ㄹ 아닌 받침 뒤": rule.AfterHard,
	"모음·ㄹ 뒤":    rule.AfterOpen,
	"모음 뒤":      rule.AfterVowel,
	"ㄹ 받침 뒤":    rule.AfterLiquid,
}

var KindNames = map[string]int{"종결": kindFinal, "연결": kindLink, "꾸밈": kindAdn, "이름": kindNoun}

var CaseNames = map[string]int{"주격": caseSubj, "보격": caseComp, "목적격": caseObj, "관형격": caseAdn, "부사격": caseAdv, "호격": caseVoc}

var FlagNames = map[string]int{"으 붙음": fEu, "서술격 안 됨": fNoIda, "요 안 됨": fNoYo, "서술격만": fNeedIda, "서술격 품음": fSelfIda}

// leadNames are the letters a form begins with when it is not written
// in whole syllables: a final consonant put under the syllable before
// it, or 어/아 for a closing that begins with 어, 아 or 여.
var leadNames = map[string]int{"ㄴ": jongN, "ㄹ": jongL, "ㅁ": jongM, "ㅂ": jongB, "ㅆ": jongSS, vowelLead: leadA}

const vowelLead = "어/아"

// The files of the tables this package reads and the names of the
// closed sets of words (ClosedWords); joinedMark marks a dependent noun
// that is read when written onto a modifying closing too.
const (
	closingsFile  = "closings"
	particlesFile = "particles"
	closedFile    = "closed-words"
	piecesFile    = "closing-pieces"

	boundSet       = "의존명사"
	helperSet      = "보조 용언"
	determinerSet  = "관형사"
	conjunctionSet = "접속 부사"
	joinedMark     = "붙여 읽음"
	maxSlot        = 5
)
