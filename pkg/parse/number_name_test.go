//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNumberName — the kind 수 with any form and the names a letter's shape gives are names of a number; the names of an order, a unit and a parenthesized Hangul letter are not
package parse

import "testing"

func TestNumberName(t *testing.T) {
	for n, want := range map[Name]bool{
		{"수", ""}: true, {"수", "십진"}: true, {"수", "2진법"}: true, {"원문자", "수"}: true, {"괄호문자", "수"}: true, {"특수", "로마자"}: true,
		{"괄호문자", "한글"}: false, {"차례", ""}: false, {"단위", "개"}: false, {"특수", "기호"}: false, {}: false,
	} {
		if got := numberName(n); got != want {
			t.Errorf("%+v: %v, want %v", n, got, want)
		}
	}
}
