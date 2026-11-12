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
	"PieceRight":    "바로 뒤에 선 것이 이 조각 뒤에 오지 못한다",
	"StemKind":      "앞말이 낱말 표에서 그 갈래의 줄기가 아니다",
	"PartsShape":    "기호가 잇는 숫자 마디들이 이 풀이의 꼴이 아니다",
	"PieceNoNumber": "바로 앞이 수가 아니다",
	"PieceLone":     "낱말 표 없이 홀로 선 한자어 한 음절 수사다",
	"MarkRight":     "기호의 오른쪽이 수가 아니다",
	"MarkAfterWord": "기호의 왼쪽에 글자가 붙어 있다",
	"NextListed":    "뒤의 맺음은 서술격 이에 붙지 않는다",
	"NextUnlisted":  "뒤의 조각이 이 뒤에 오는 것이 아니다",
	"FormUnlisted":  "이 조각 뒤에 붙는 것이 아니다",
	"PieceMarked":   "경계 뒤의 통짜 새 낱말이 표지로 읽힌다",
	"PieceLast":     "뒤에 조사가 없다",
	"PieceNoAdn":    "앞 어절이 꾸밈 맺음으로 끝나지 않는다",
	"PieceAdn":      "앞 어절이 꾸밈 맺음으로 끝난다",
	"StemCoda":      "새 낱말의 줄기가 ㅆ 받침으로 끝난다",
	"MarkSides":     "양쪽이 맞는 글자 열이 아니다",
	"MarkBoth":      "양쪽이 모두 그 글자 열이다",
	"MarkAlone":     "점이 하나뿐이다",
	"MarkRepeated":  "점이 이어진다",
	"ExtUnlisted":   "확장자 목록에 없다",
	"LastUnlisted":  "최상위 도메인 목록에 없다",
	"VersionShape":  "버전의 꼴이 아니다",
}
