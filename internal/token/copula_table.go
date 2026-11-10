package token

// The copula 이다 (morphemes.md section 2). Its stem 이 after a noun is
// a reading of its own, and after a vowel the stem may be dropped
// before the closings listed here (학교다, 학교고, 학교며, 학교면,
// 학교지만, 학교라서, 학교라면, 학교지, 누군지, 뭔데, 누군가, 뭘까),
// and before 잖 (학교잖아: 지 and 않). copulaCount is its observed count in the tables.
const (
	copulaName  = "서술격:이"
	droppedName = "서술격:이(빠짐)"
	copulaCount = 4520000
)

var droppable = []string{
	"맺음:다", "맺음:고", "맺음:며", "맺음:면", "맺음:지만", "꾸밈:ㄴ", "맺음:ㅂ니다",
	"맺음:라서", "맺음:라면", "맺음:지", "맺음:ㄴ지", "맺음:ㄴ데", "맺음:ㄴ가", "맺음:ㄹ까",
}

// dropAdn is the closing before which a dropped stem is rebutted: in
// front of a modifying form the stem mostly stays (의사인, not 의산).
// forceDropAdn is the rebuttal's force, a placeholder chosen among a
// few values tried on a thousand sentences.
var dropAdn = []string{"꾸밈:ㄴ"}

const forceDropAdn = 0.5

// notOnly is the one other stem the closings of the copula attach to.
const notOnly = "아니"
