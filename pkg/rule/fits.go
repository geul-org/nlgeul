//ff:func feature=rule type=helper control=selection
//ff:what Fits — 조건식(규칙 PieceAfter와 읽기 코드가 함께 쓴다). 앞 음절의 끝이 조각이 붙는 조건(After…)과 맞는가; 한글 음절이 아닌 것은 모든 조건과 맞는다
package rule

func Fits(prev rune, after int) bool {
	j := JongOf(prev)
	switch after {
	case AfterBatchim:
		return j != 0
	case AfterHard:
		return j != 0 && j != JongL
	case AfterOpen:
		return j <= 0 || j == JongL
	case AfterVowel:
		return j <= 0
	}
	return true
}
