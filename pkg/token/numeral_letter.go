//ff:func feature=token type=helper control=sequence
//ff:what numeralLetter — is r a letter a Hangul numeral is written with: one of the Sino-Korean digits and greater numbers, 수, or a letter of the native numerals of the tables; a run of numerals holds no other letter, so it is looked for only as far back as such letters go
package token

import "strings"

func numeralLetter(r rune) bool {
	native := strings.Join(nativeOnes, "") + strings.Join(nativeAlone, "")
	return r == sinoSome || strings.ContainsRune(sinoDigits+sinoBigs+native, r)
}
