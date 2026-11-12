//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNumeralLetter — the Sino-Korean digits and greater numbers, 수, and every letter of the native numerals of the tables are letters of a numeral; a unit word, another Hangul letter, a Latin letter, a digit and a sign are not
package token

import "testing"

func TestNumeralLetter(t *testing.T) {
	for _, r := range "일이삼사오육칠팔구십백천만억조수한두세네다섯여섯일곱여덟아홉열스무몇여러" {
		if !numeralLetter(r) {
			t.Errorf("%c is a letter of a numeral", r)
		}
	}
	for _, r := range "명개곰영공번째aN3 .-" {
		if numeralLetter(r) {
			t.Errorf("%c is no letter of a numeral", r)
		}
	}
}
