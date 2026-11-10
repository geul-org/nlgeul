//ff:func feature=rule type=helper control=sequence
//ff:what RunRune — 도움 함수(규칙 아님). 영문·숫자 열을 잇는 글자인가: 라틴 문자, 숫자, 밑줄
package rule

import "unicode"

func RunRune(r rune) bool {
	return r == '_' || unicode.IsDigit(r) || unicode.Is(unicode.Latin, r)
}
