//ff:func feature=rule type=helper control=selection
//ff:what Fits — 조건식(규칙 PieceAfter와 읽기 코드가 함께 쓴다). 앞 음절의 끝이 조각이 붙는 조건(After…)과 맞는가; 한글 음절이 아닌 것은 받침을 묻는 조건 넷과 맞고, ㄹ 받침을 묻는 조건(AfterLiquid)과 한글 음절인지 묻는 조건(AfterHangul)과는 맞지 않는다
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
	case AfterLiquid:
		return j == JongL
	case AfterHangul:
		return j >= 0
	}
	return true
}
