//ff:func feature=token type=helper control=sequence
//ff:what TestIsDigitFirst — a text that begins with a digit 0 to 9; not an empty one, nor one that begins with a Hangul letter or a Latin one
package token

import "testing"

func TestIsDigitFirst(t *testing.T) {
	if !isDigitFirst("3개") || !isDigitFirst("0") || isDigitFirst("") || isDigitFirst("세") || isDigitFirst("a1") {
		t.Fatal("digit first")
	}
}
