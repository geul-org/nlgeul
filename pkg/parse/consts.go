package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// Cases a particle marks, as particles.md names its function, and the
// kinds of closing, as endings.md section 3 marks them. The rules read
// them, so they are the rule package's.
const (
	caseSubj = rule.CaseSubj
	caseComp = rule.CaseComp
	caseObj  = rule.CaseObj
	caseAdn  = rule.CaseAdn
	caseAdv  = rule.CaseAdv
	caseVoc  = rule.CaseVoc

	kindFinal = rule.KindFinal
	kindLink  = rule.KindLink
	kindAdn   = rule.KindAdn
	kindNoun  = rule.KindNoun
)

// Attachment conditions, final-consonant indexes and the lead of an
// 어/아 closing: the rule package's, which the tables here are written
// with.
const (
	afterAny     = rule.AfterAny
	afterBatchim = rule.AfterBatchim

	leadNone = 0
	leadA    = rule.LeadVowel
	jongN    = rule.JongN
	jongL    = rule.JongL
	jongM    = rule.JongM
	jongB    = rule.JongB
	jongSS   = rule.JongSS
)

// Ending flags: how a closing attaches. None of them names a constituent.
const (
	fEu      = 1 << iota // takes 으 after a consonant, attaches bare after a vowel or ㄹ
	fNoIda               // does not attach to the copula stem 이 (는, 은, 을, 습니다, and the closings of verbs)
	fNoYo                // 요 does not follow it
	fNeedIda             // attaches only to the copula stem 이 or to 아니 (에요, 야)
	fSelfIda             // holds the copula stem itself (예요 = 이 + 에요)
)
