//ff:func feature=rule type=helper control=sequence
//ff:what JungOf — 도움 함수(규칙 아님). 한글 음절의 모음 번호(ㅏ 0 … ㅣ 20); 한글 음절이 아니면 -1
package rule

func JungOf(r rune) int {
	if r < 0xAC00 || r > 0xD7A3 {
		return -1
	}
	return int(r-0xAC00) / 28 % 21
}
