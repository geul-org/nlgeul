//ff:func feature=rule type=helper control=selection
//ff:what Harmony — 조건식(규칙 PieceClash가 쓴다). 줄기 끝 음절 prev 뒤에 어/아/여로 시작하는 조각의 첫 글자 first가 모음에 맞는가(한글 맞춤법 제16항·제18항): 끝 모음 ㅏ·ㅑ·ㅗ 뒤는 아, 그 밖은 어, 하 뒤는 여; 받침 ㅆ 뒤(았·었·겠 뒤)는 늘 어; 한글 음절이 아닌 것 뒤는 무엇이든 맞다
package rule

func Harmony(prev, first rune) bool {
	v := JungOf(prev)
	bright := (v == 0 || v == 2 || v == 8) && JongOf(prev) != JongSS
	switch {
	case v < 0:
		return true
	case prev == '하':
		return first == '여' || first == '였'
	case first == '아' || first == '았':
		return bright
	case first == '어' || first == '었':
		return !bright
	case first == '여' || first == '였':
		return false
	}
	return true
}
