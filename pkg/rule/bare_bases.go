//ff:func feature=rule type=helper control=sequence
//ff:what BareBases — 도움 함수(규칙 Listed의 기준(합쳐진 조각이 놓인 음절)과 읽기 코드가 함께 쓴다). 받침 하나뿐인 맺음을 읽는 음절: ㅁ은 하·되·이 밑, ㄴ·ㄹ·ㅂ은 하·되·시·키·이 밑
package rule

func BareBases(lead int) []string {
	if lead == JongM {
		return []string{"하", "되", "이"}
	}
	return []string{"하", "되", "시", "키", "이"}
}
