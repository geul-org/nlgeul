//ff:func feature=rule type=helper control=sequence
//ff:what ChoOf — 도움 함수(규칙 아님). 한글 음절의 첫소리 번호(ㄱ 0 … ㅎ 18); 한글 음절이 아니면 -1
package rule

func ChoOf(r rune) int {
	if r < 0xAC00 || r > 0xD7A3 {
		return -1
	}
	return int(r-0xAC00) / 588
}
