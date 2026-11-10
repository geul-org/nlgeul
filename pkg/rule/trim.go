//ff:func feature=rule type=helper control=sequence
//ff:what Trim — 도움 함수(규칙 아님). 어절 앞뒤의 문장 부호와 기호를 뗀 토큰 열(문장 부호는 확정 경계다)
package rule

import (
	"strings"
	"unicode"
)

func Trim(field string) []rune {
	return []rune(strings.TrimFunc(field, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	}))
}
