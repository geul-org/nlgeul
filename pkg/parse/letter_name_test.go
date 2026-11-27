//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestLetterName — a circled number (①, ⑳, ㉑, ㊿) is named 원문자:수, a parenthesized number (⑴, ⒇) 괄호문자:수, a Latin letter 특수:로마자; any other letter (a circled Latin letter, Hangul, a digit, a parenthesized Hangul letter) none
package parse

import "testing"

func TestLetterName(t *testing.T) {
	for r, want := range map[rune]string{
		'①': circledNumName, '⑳': circledNumName, '㉑': circledNumName, '㊿': circledNumName,
		'⑴': parenNumName, '⒇': parenNumName,
		'n': latinLetterName, 'X': latinLetterName,
		'ⓛ': "", '가': "", '7': "", '㈎': "", 'ⅳ': "",
	} {
		if got := letterName(r); got != want {
			t.Errorf("%c: %q, want %q", r, got, want)
		}
	}
}
