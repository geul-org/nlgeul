package token

// The copula 이다 (morphemes.md section 2). Its stem 이 after a noun is
// a reading of its own, and after a vowel the stem may be dropped
// before the closings listed here (학교다, 학교고, 학교며, 학교면,
// 학교지만). copulaCount is its observed count in the tables.
const (
	copulaName  = "서술격:이"
	droppedName = "서술격:이(빠짐)"
	copulaCount = 4520000
)

var droppable = []string{"맺음:다", "맺음:고", "맺음:며", "맺음:면", "맺음:지만", "꾸밈:ㄴ", "맺음:ㅂ니다"}

// notOnly is the one other stem the closings of the copula attach to.
const notOnly = "아니"
