package token

// reasons says, for a trace, why an exception defeated a claim, by the
// exception's rule function.
var reasons = map[string]string{
	"PieceNoStem": "앞에 줄기가 없다",
	"PieceAfter":  "앞 토큰의 끝이 맞지 않는다",
	"ChainOrder":  "조사 사슬의 순서가 어긋난다",
	"SlotOrder":   "선어말 칸의 순서가 어긋난다",
	"PieceBase":   "받침만 붙은 꼴을 읽는 음절이 아니다",
	"WordListed":  "통째로 닫힌 집합의 낱말이다",
}
