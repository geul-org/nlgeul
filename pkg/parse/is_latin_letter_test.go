//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestIsLatinLetter — ASCII and fullwidth Latin letters are, a digit, a fullwidth digit, a Hangul letter and the replacement character are not
package parse

import "testing"

func TestIsLatinLetter(t *testing.T) {
	for _, r := range "aZｍＡｚＺ" {
		if !isLatinLetter(r) {
			t.Fatal(string(r))
		}
	}
	for _, r := range "5０가％�" {
		if isLatinLetter(r) {
			t.Fatal(string(r))
		}
	}
}
