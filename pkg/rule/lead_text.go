//ff:func feature=rule type=helper control=selection
//ff:what LeadText — 도움 함수(규칙 아님). 조각의 머리를 글자로: 받침 ㄴ ㄹ ㅁ ㅂ ㅆ, 합쳐지는 어/아는 어, 온 음절 조각은 빈 글
package rule

func LeadText(lead int) string {
	switch lead {
	case JongN:
		return "ㄴ"
	case JongL:
		return "ㄹ"
	case JongM:
		return "ㅁ"
	case JongB:
		return "ㅂ"
	case JongSS:
		return "ㅆ"
	case LeadVowel:
		return "어"
	}
	return ""
}
