package token

// reasons says, for a trace, why an exception defeated a claim, by the
// exception's rule function.
var reasons = map[string]string{
	"PieceNoStem":   "앞에 줄기가 없다",
	"PieceAfter":    "앞 토큰의 끝이 맞지 않는다",
	"ChainLinks":    "조사 사슬에 맞지 않는다",
	"SlotOrder":     "선어말 칸의 순서가 어긋난다",
	"PieceBase":     "받침만 붙은 꼴을 읽는 음절이 아니다",
	"WordListed":    "통째로 닫힌 집합의 낱말이다",
	"PieceParticle": "바로 뒤에 조사가 붙는다",
	"PieceCut":      "떼어 낸 받침이 뒤 토큰과 표의 맺음을 이루지 않는다",
	"StemKind":      "앞말이 낱말 표에서 그 갈래의 줄기가 아니다",
	"PartsShape":    "기호가 잇는 숫자 마디들이 이 풀이의 꼴이 아니다",
	"PieceRunsOn":   "단위 뒤에 글자가 그대로 이어진다",
	"PieceNoNumber": "바로 앞이 수가 아니다",
	"PieceLone":     "낱말 표 없이 홀로 선 한자어 한 음절 수사다",
	"MarkRight":     "기호의 오른쪽이 수가 아니다",
	"MarkAfterWord": "기호의 왼쪽에 글자가 붙어 있다",
	"NextListed":    "뒤의 맺음은 서술격 이에 붙지 않는다",
	"PieceLast":     "뒤에 조사가 없다",
	"MarkSides":     "양쪽이 맞는 글자 열이 아니다",
	"MarkEnclosed":  "양쪽이 영숫자 열이다",
	"MarkAlone":     "점이 하나뿐이다",
	"MarkRepeated":  "점이 이어진다",
	"ExtUnlisted":   "확장자 목록에 없다",
	"LastUnlisted":  "최상위 도메인 목록에 없다",
	"VersionShape":  "버전의 꼴이 아니다",
	"DecimalShape":  "소수의 꼴이 아니다",
}
