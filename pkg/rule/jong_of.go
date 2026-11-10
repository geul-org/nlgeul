//ff:func feature=rule type=helper control=sequence
//ff:what JongOf — 도움 함수(규칙 아님). 한글 음절의 받침 번호(없으면 0); 한글 음절이 아니면 -1
package rule

func JongOf(r rune) int {
	if r < 0xAC00 || r > 0xD7A3 {
		return -1
	}
	return int(r-0xAC00) % 28
}
